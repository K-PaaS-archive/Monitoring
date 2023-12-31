package goci_test

import (
	"os/exec"

	"code.cloudfoundry.org/guardian/rundmc/goci"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
	BeforeEach(func() {
		goci.DefaultRuncBinary = goci.RuncBinary{Path: "funC", Root: ""}
	})

	Describe("StartCommand", func() {
		It("creates an *exec.Cmd to start a bundle", func() {
			cmd := goci.StartCommand("my-bundle-path", "my-bundle-id", false, "mylog.file")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "mylog.file", "--log-format", "json", "start", "my-bundle-id"}))
			Expect(cmd.Dir).To(Equal("my-bundle-path"))
		})

		It("turns on debug logging", func() {
			cmd := goci.StartCommand("my-bundle-path", "my-bundle-id", true, "mylog.file")
			Expect(cmd.Args).To(ContainElement("--debug"))
		})

		It("passes the detach flag if requested", func() {
			cmd := goci.StartCommand("my-bundle-path", "my-bundle-id", true, "mylog.file")
			Expect(cmd.Args).To(ContainElement("-d"))
		})
	})

	Describe("ExecCommand", func() {
		It("creates an *exec.Cmd to exec a process in a bundle", func() {
			cmd := goci.ExecCommand("my-bundle-id", "my-process-json.json", "some-pid-file")
			Expect(cmd.Args).To(Equal([]string{"funC", "exec", "my-bundle-id", "--pid-file", "some-pid-file", "-p", "my-process-json.json"}))
		})
	})

	Describe("EventsCommand", func() {
		It("creates an *exec.Cmd to get events for a bundle", func() {
			cmd := goci.EventsCommand("my-bundle-id")
			Expect(cmd.Args).To(Equal([]string{"funC", "events", "my-bundle-id"}))
		})
	})

	Describe("KillCommand", func() {
		It("creates an *exec.Cmd to signal the bundle", func() {
			cmd := goci.KillCommand("my-bundle-id", "TERM", "log.file")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "log.file", "--log-format", "json", "kill", "my-bundle-id", "TERM"}))
		})
	})

	Describe("StateCommand", func() {
		It("creates an *exec.Cmd to get the state of the bundle", func() {
			cmd := goci.StateCommand("my-bundle-id", "log.file")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "log.file", "--log-format", "json", "state", "my-bundle-id"}))
		})
	})

	Describe("StatsCommand", func() {
		It("creates an *exec.Cmd to get the state of the bundle", func() {
			cmd := goci.StatsCommand("my-bundle-id", "log.file")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "log.file", "--log-format", "json", "events", "--stats", "my-bundle-id"}))
		})
	})

	Describe("DeleteCommand", func() {
		It("creates an *exec.Cmd to delete the bundle", func() {
			cmd := goci.DeleteCommand("my-bundle-id", "log.file")
			Expect(cmd.Args).To(Equal([]string{"funC", "--debug", "--log", "log.file", "--log-format", "json", "delete", "my-bundle-id"}))
		})
	})

	Context("when a runc root is passed", func() {
		BeforeEach(func() {
			goci.DefaultRuncBinary = goci.RuncBinary{Path: "funC", Root: "/run/funC"}
		})

		DescribeTable(
			"passing --root to runc commands",
			func(commander func() *exec.Cmd) {
				cmd := commander()
				Expect(cmd.Args[0:3]).To(Equal([]string{"funC", "--root", "/run/funC"}))
			},
			Entry("StartCommand", func() *exec.Cmd {
				return goci.StartCommand("", "", false, "")
			}),
			Entry("ExecCommand", func() *exec.Cmd {
				return goci.ExecCommand("", "", "")
			}),
			Entry("EventsCommand", func() *exec.Cmd {
				return goci.EventsCommand("")
			}),
			Entry("KillCommand", func() *exec.Cmd {
				return goci.KillCommand("", "", "")
			}),
			Entry("StateCommand", func() *exec.Cmd {
				return goci.StateCommand("", "")
			}),
			Entry("StatsCommand", func() *exec.Cmd {
				return goci.StatsCommand("", "")
			}),
			Entry("DeleteCommand", func() *exec.Cmd {
				return goci.DeleteCommand("", "")
			}),
		)
	})
})
