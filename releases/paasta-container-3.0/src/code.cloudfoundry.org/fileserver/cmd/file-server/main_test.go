package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"code.cloudfoundry.org/fileserver/cmd/file-server/config"

	"github.com/hashicorp/consul/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

type ByteEmitter struct {
	written int
	length  int
}

func NewEmitter(length int) *ByteEmitter {
	return &ByteEmitter{
		length:  length,
		written: 0,
	}
}

func (emitter *ByteEmitter) Read(p []byte) (n int, err error) {
	if emitter.written >= emitter.length {
		return 0, io.EOF
	}
	time.Sleep(time.Millisecond)
	p[0] = 0xF1
	emitter.written++
	return 1, nil
}

var _ = Describe("File server", func() {
	var (
		port            int
		servedDirectory string
		session         *gexec.Session
		err             error
		configPath      string
	)

	start := func(extras ...string) *gexec.Session {
		args := []string{"-config", configPath}
		session, err = gexec.Start(exec.Command(fileServerBinary, args...), GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gbytes.Say("file-server.ready"))

		return session
	}

	BeforeEach(func() {
		servedDirectory, err = ioutil.TempDir("", "file_server-test")
		Expect(err).NotTo(HaveOccurred())

		configFile, err := ioutil.TempFile("", "file_server-test-config")
		Expect(err).NotTo(HaveOccurred())
		configPath = configFile.Name()

		port = 8182 + GinkgoParallelNode()
		cfg := config.FileServerConfig{
			StaticDirectory: servedDirectory,
			ConsulCluster:   consulRunner.URL(),
			ServerAddress:   fmt.Sprintf("localhost:%d", port),
		}

		encoder := json.NewEncoder(configFile)
		err = encoder.Encode(&cfg)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		session.Kill().Wait()
		os.RemoveAll(servedDirectory)
		os.RemoveAll(configPath)
	})

	Context("when started without any arguments", func() {
		It("should fail", func() {
			session, err = gexec.Start(exec.Command(fileServerBinary), GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(2))
			Eventually(session.Out).Should(gbytes.Say("failed-to-parse-config"))
		})
	})

	Context("when started correctly", func() {
		BeforeEach(func() {
			session = start()
			ioutil.WriteFile(filepath.Join(servedDirectory, "test"), []byte("hello"), os.ModePerm)
		})

		It("should return that file on GET request", func() {
			resp, err := http.Get(fmt.Sprintf("http://localhost:%d/v1/static/test", port))
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(body)).To(Equal("hello"))
		})

		It("registers itself with consul", func() {
			services, err := consulRunner.NewClient().Agent().Services()
			Expect(err).NotTo(HaveOccurred())
			Expect(services).Should(HaveKeyWithValue("file-server",
				&api.AgentService{
					Service: "file-server",
					ID:      "file-server",
					Port:    port,
				}))
		})

		It("registers a TTL healthcheck", func() {
			checks, err := consulRunner.NewClient().Agent().Checks()
			Expect(err).NotTo(HaveOccurred())
			Expect(checks).Should(HaveKeyWithValue("service:file-server",
				&api.AgentCheck{
					Node:        "0",
					CheckID:     "service:file-server",
					Name:        "Service 'file-server' check",
					Status:      "passing",
					ServiceID:   "file-server",
					ServiceName: "file-server",
				}))
		})
	})
})
