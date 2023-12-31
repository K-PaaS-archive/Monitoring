package runrunc_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"code.cloudfoundry.org/guardian/rundmc/runrunc"
	fakes "code.cloudfoundry.org/guardian/rundmc/runrunc/runruncfakes"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontainers/runtime-spec/specs-go"
)

var _ = Describe("Execer", func() {
	var (
		logger *lagertest.TestLogger

		execPreparer *fakes.FakeExecPreparer
		execRunner   *fakes.FakeExecRunner
		execer       *runrunc.Execer
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
		execRunner = new(fakes.FakeExecRunner)
		execPreparer = new(fakes.FakeExecPreparer)

		execer = runrunc.NewExecer(
			execPreparer,
			execRunner,
		)
	})

	It("runs the execRunner with the passed process ID and the prepared process spec", func() {
		execPreparer.PrepareStub = func(log lager.Logger, bundlePath string, spec garden.ProcessSpec) (*runrunc.PreparedSpec, error) {
			return &runrunc.PreparedSpec{
				Process: specs.Process{
					Args: []string{spec.Path, bundlePath},
				},
				HostUID: 10,
				HostGID: 10,
			}, nil
		}

		_, err := execer.Exec(logger, "some-bundle-path", "some-id", garden.ProcessSpec{
			ID:   "some-process-id",
			Path: "potato",
		}, garden.ProcessIO{})
		Expect(err).NotTo(HaveOccurred())

		Expect(execRunner.RunCallCount()).To(Equal(1))
		_, passedID, spec, bundlePath, processesPath, id, _, _ := execRunner.RunArgsForCall(0)
		Expect(passedID).To(Equal("some-process-id"))
		Expect(spec.Args).To(ConsistOf("potato", "some-bundle-path"))
		Expect(bundlePath).To(Equal("some-bundle-path"))
		Expect(processesPath).To(Equal("some-bundle-path/processes"))
		Expect(id).To(Equal("some-id"))
	})
})

var _ = Describe("ExecPreparer", func() {
	var (
		bndl         goci.Bndl
		spec         *runrunc.PreparedSpec
		bundleLoader *fakes.FakeBundleLoader
		users        *fakes.FakeUserLookupper
		enver        *fakes.FakeEnvDeterminer
		mkdirer      *fakes.FakeMkdirer
		bundlePath   string
		logger       lager.Logger
		rootless     bool

		preparer runrunc.ExecPreparer
	)

	BeforeEach(func() {
		rootless = false

		bndl = goci.Bundle().WithHostname("some-hostname")
		logger = lagertest.NewTestLogger("test")
		bundleLoader = new(fakes.FakeBundleLoader)
		users = new(fakes.FakeUserLookupper)
		enver = new(fakes.FakeEnvDeterminer)
		mkdirer = new(fakes.FakeMkdirer)

		var err error
		bundlePath, err = ioutil.TempDir("", "bundle")
		Expect(err).NotTo(HaveOccurred())

		bundleLoader.LoadStub = func(path string) (goci.Bndl, error) {
			return bndl, nil
		}

		users.LookupReturns(&runrunc.ExecUser{}, nil)
		enver.EnvForReturns([]string{"FOO=bar"})

		Expect(ioutil.WriteFile(filepath.Join(bundlePath, "pidfile"), []byte("999"), 0644)).To(Succeed())
	})

	JustBeforeEach(func() {
		runningAsRoot := func() bool { return !rootless }
		preparer = runrunc.NewExecPreparer(bundleLoader, users, enver, mkdirer, []string{"foo", "bar", "brains"}, runningAsRoot)
	})

	It("passes a process.json with the correct path and args", func() {
		spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{Path: "to enlightenment", Args: []string{"infinity", "and beyond"}})
		Expect(err).NotTo(HaveOccurred())

		Expect(spec.Args).To(Equal([]string{"to enlightenment", "infinity", "and beyond"}))
	})

	It("returns the HostUID and HostGID in the returned spec", func() {
		users.LookupReturns(&runrunc.ExecUser{Uid: 234, Gid: 567}, nil)

		spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{Path: "to enlightenment", Args: []string{}})
		Expect(err).NotTo(HaveOccurred())

		Expect(spec.HostUID).To(BeEquivalentTo(234))
		Expect(spec.HostGID).To(BeEquivalentTo(567))
	})

	It("passes the correct env", func() {
		users.LookupReturns(&runrunc.ExecUser{Uid: 234, Gid: 567}, nil)

		processSpec := garden.ProcessSpec{ID: "some-id"}
		preparedSpec, err := preparer.Prepare(logger, bundlePath, processSpec)
		Expect(err).NotTo(HaveOccurred())

		Expect(enver.EnvForCallCount()).To(Equal(1))
		actualUID, actualBndl, actualProcessSpec := enver.EnvForArgsForCall(0)
		Expect(actualUID).To(Equal(234))
		Expect(actualBndl).To(Equal(bndl))
		Expect(actualProcessSpec).To(Equal(processSpec))

		Expect(preparedSpec.Env).To(Equal([]string{"FOO=bar"}))
	})

	It("sets the rlimits correctly", func() {
		ptr := func(n uint64) *uint64 { return &n }
		spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{
			Limits: garden.ResourceLimits{
				As:         ptr(12),
				Core:       ptr(24),
				Cpu:        ptr(36),
				Data:       ptr(99),
				Fsize:      ptr(101),
				Locks:      ptr(111),
				Memlock:    ptr(987),
				Msgqueue:   ptr(777),
				Nice:       ptr(111),
				Nofile:     ptr(222),
				Nproc:      ptr(1234),
				Rss:        ptr(888),
				Rtprio:     ptr(254),
				Sigpending: ptr(101),
				Stack:      ptr(44),
			},
		})
		Expect(err).ToNot(HaveOccurred())

		Expect(spec.Process.Rlimits).To(ConsistOf(
			specs.POSIXRlimit{Type: "RLIMIT_AS", Hard: 12, Soft: 12},
			specs.POSIXRlimit{Type: "RLIMIT_CORE", Hard: 24, Soft: 24},
			specs.POSIXRlimit{Type: "RLIMIT_CPU", Hard: 36, Soft: 36},
			specs.POSIXRlimit{Type: "RLIMIT_DATA", Hard: 99, Soft: 99},
			specs.POSIXRlimit{Type: "RLIMIT_FSIZE", Hard: 101, Soft: 101},
			specs.POSIXRlimit{Type: "RLIMIT_LOCKS", Hard: 111, Soft: 111},
			specs.POSIXRlimit{Type: "RLIMIT_MEMLOCK", Hard: 987, Soft: 987},
			specs.POSIXRlimit{Type: "RLIMIT_MSGQUEUE", Hard: 777, Soft: 777},
			specs.POSIXRlimit{Type: "RLIMIT_NICE", Hard: 111, Soft: 111},
			specs.POSIXRlimit{Type: "RLIMIT_NOFILE", Hard: 222, Soft: 222},
			specs.POSIXRlimit{Type: "RLIMIT_NPROC", Hard: 1234, Soft: 1234},
			specs.POSIXRlimit{Type: "RLIMIT_RSS", Hard: 888, Soft: 888},
			specs.POSIXRlimit{Type: "RLIMIT_RTPRIO", Hard: 254, Soft: 254},
			specs.POSIXRlimit{Type: "RLIMIT_SIGPENDING", Hard: 101, Soft: 101},
			specs.POSIXRlimit{Type: "RLIMIT_STACK", Hard: 44, Soft: 44},
		))
	})

	It("sets a default console size", func() {
		spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{
			TTY: &garden.TTYSpec{},
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(spec.Process.ConsoleSize.Height).To(BeEquivalentTo(24))
		Expect(spec.Process.ConsoleSize.Width).To(BeEquivalentTo(80))
	})

	It("sets console size if a TTY is configured with a WindowSize", func() {
		spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{
			TTY: &garden.TTYSpec{
				WindowSize: &garden.WindowSize{
					Columns: 25,
					Rows:    81,
				},
			},
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(spec.Process.ConsoleSize.Width).To(BeEquivalentTo(25))
		Expect(spec.Process.ConsoleSize.Height).To(BeEquivalentTo(81))
	})

	It("sets Terminal to true iff a TTY is configured", func() {
		spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{
			TTY: &garden.TTYSpec{},
		})
		Expect(err).ToNot(HaveOccurred())

		Expect(spec.Process.Terminal).To(BeTrue())

		spec, err = preparer.Prepare(logger, bundlePath, garden.ProcessSpec{
			TTY: nil,
		})
		Expect(err).ToNot(HaveOccurred())

		Expect(spec.Process.Terminal).To(BeFalse())
	})

	Describe("passing the correct uid and gid", func() {
		Context("when the bundle can be loaded", func() {
			JustBeforeEach(func() {
				users.LookupReturns(&runrunc.ExecUser{Uid: 9, Gid: 7}, nil)

				var err error
				spec, err = preparer.Prepare(logger, bundlePath, garden.ProcessSpec{User: "spiderman"})
				Expect(err).ToNot(HaveOccurred())
			})

			It("looks up the user and group IDs of the user in the right rootfs", func() {
				Expect(users.LookupCallCount()).To(Equal(1))
				actualRootfsPath, actualUserName := users.LookupArgsForCall(0)
				Expect(actualRootfsPath).To(Equal(filepath.Join("/proc", "999", "root")))
				Expect(actualUserName).To(Equal("spiderman"))
			})

			It("passes a process.json with the correct user name, user id, and group id", func() {
				Expect(spec.Process.User).To(Equal(specs.User{Username: "spiderman", UID: 9, GID: 7, AdditionalGids: []uint32{}}))
			})
		})

		Context("when the bundle can't be loaded", func() {
			JustBeforeEach(func() {
				bundleLoader.LoadReturns(goci.Bundle(), errors.New("whoa! Hold them horses!"))
			})

			It("fails", func() {
				_, err := preparer.Prepare(logger, bundlePath,
					garden.ProcessSpec{User: "spiderman"})
				Expect(err).To(MatchError(ContainSubstring("Hold them horses")))
			})
		})

		Context("when User Lookup returns an error", func() {
			It("passes a process.json with the correct user and group ids", func() {
				users.LookupReturns(&runrunc.ExecUser{Uid: 0, Gid: 0}, errors.New("bang"))

				_, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{User: "spiderman"})
				Expect(err).To(MatchError(ContainSubstring("bang")))
			})
		})

		Context("when the pidfile can't be read", func() {
			It("returns an appropriate error", func() {
				Expect(os.Remove(filepath.Join(bundlePath, "pidfile"))).To(Succeed())

				_, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{User: "spiderman"})
				Expect(err).To(MatchError(ContainSubstring("pidfile")))
			})
		})
	})

	Context("when the container has capabilities", func() {
		BeforeEach(func() {
			bndl := goci.Bundle()
			bndl = bndl.WithCapabilities("foo", "bar", "baz")
			bundleLoader.LoadReturns(bndl, nil)
		})

		Context("and the user is root", func() {
			It("passes them on to the process", func() {
				spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{})
				Expect(err).NotTo(HaveOccurred())
				Expect(spec.Process.Capabilities.Effective).To(BeEmpty())
				Expect(spec.Process.Capabilities.Bounding).To(Equal([]string{"foo", "bar", "baz"}))
				Expect(spec.Process.Capabilities.Inheritable).To(Equal([]string{"foo", "bar", "baz"}))
				Expect(spec.Process.Capabilities.Permitted).To(Equal([]string{"foo", "bar", "baz"}))
				Expect(spec.Process.Capabilities.Ambient).To(BeEmpty())
			})
		})

		Context("and the user is not root", func() {
			It("removes any caps not in nonRootMaxCaps list", func() {
				users.LookupReturns(&runrunc.ExecUser{Uid: 1234, Gid: 0}, nil)
				spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{})

				Expect(err).NotTo(HaveOccurred())
				Expect(spec.Process.Capabilities.Effective).To(BeEmpty())
				Expect(spec.Process.Capabilities.Bounding).To(Equal([]string{"foo", "bar"}))
				Expect(spec.Process.Capabilities.Inheritable).To(Equal([]string{"foo", "bar"}))
				Expect(spec.Process.Capabilities.Permitted).To(Equal([]string{"foo", "bar"}))
				Expect(spec.Process.Capabilities.Ambient).To(BeEmpty())
			})
		})
	})

	Context("when the container has no capabilities", func() {
		BeforeEach(func() {
			bndl := goci.Bundle()
			bundleLoader.LoadReturns(bndl, nil)
		})

		It("does not set the capabilities object on the process", func() {
			spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{})
			Expect(err).NotTo(HaveOccurred())
			Expect(spec.Process.Capabilities).To(BeNil())
		})
	})

	Describe("working directory", func() {
		Context("when the working directory is specified", func() {
			It("passes the correct cwd to the spec", func() {
				spec, err := preparer.Prepare(
					logger, bundlePath,
					garden.ProcessSpec{Dir: "/home/dir"},
				)
				Expect(err).NotTo(HaveOccurred())

				Expect(spec.Process.Cwd).To(Equal("/home/dir"))
			})

			Describe("Creating the working directory", func() {
				JustBeforeEach(func() {
					users.LookupReturns(&runrunc.ExecUser{Uid: 1012, Gid: 1013}, nil)

					_, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{
						Dir: "/path/to/banana/dir",
					})
					Expect(err).NotTo(HaveOccurred())
				})

				Context("when the container is privileged", func() {
					It("creates the working directory", func() {
						Expect(mkdirer.MkdirAsCallCount()).To(Equal(1))
						rootfs, uid, gid, mode, recreate, dirs := mkdirer.MkdirAsArgsForCall(0)
						Expect(rootfs).To(Equal(filepath.Join("/proc", "999", "root")))
						Expect(dirs).To(ConsistOf("/path/to/banana/dir"))
						Expect(mode).To(BeNumerically("==", 0755))
						Expect(recreate).To(BeFalse())
						Expect(uid).To(BeEquivalentTo(1012))
						Expect(gid).To(BeEquivalentTo(1013))
					})
				})

				Context("when the container is unprivileged", func() {
					BeforeEach(func() {
						bundleLoader.LoadStub = func(path string) (goci.Bndl, error) {
							bndl := goci.Bundle()
							bndl.Spec.Linux.UIDMappings = []specs.LinuxIDMapping{{
								HostID:      1712,
								ContainerID: 1012,
								Size:        1,
							}}
							bndl.Spec.Linux.GIDMappings = []specs.LinuxIDMapping{{
								HostID:      1713,
								ContainerID: 1013,
								Size:        1,
							}}
							return bndl, nil
						}
					})

					It("creates the working directory as the mapped user", func() {
						Expect(mkdirer.MkdirAsCallCount()).To(Equal(1))
						rootfs, uid, gid, mode, recreate, dirs := mkdirer.MkdirAsArgsForCall(0)
						Expect(rootfs).To(Equal(filepath.Join("/proc", "999", "root")))
						Expect(dirs).To(ConsistOf("/path/to/banana/dir"))
						Expect(mode).To(BeEquivalentTo(0755))
						Expect(recreate).To(BeFalse())
						Expect(uid).To(BeEquivalentTo(1712))
						Expect(gid).To(BeEquivalentTo(1713))
					})
				})
			})
		})

		Context("when the working directory is not specified", func() {
			It("defaults to the user's HOME directory", func() {
				users.LookupReturns(&runrunc.ExecUser{Home: "/the/home/dir"}, nil)

				spec, err := preparer.Prepare(
					logger, bundlePath,
					garden.ProcessSpec{Dir: ""},
				)
				Expect(err).NotTo(HaveOccurred())

				Expect(spec.Process.Cwd).To(Equal("/the/home/dir"))
			})

			It("creates the directory", func() {
				users.LookupReturns(&runrunc.ExecUser{Uid: 1012, Gid: 1013, Home: "/some/dir"}, nil)

				_, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{})
				Expect(err).NotTo(HaveOccurred())

				Expect(mkdirer.MkdirAsCallCount()).To(Equal(1))
				_, _, _, _, _, dirs := mkdirer.MkdirAsArgsForCall(0)
				Expect(dirs).To(ConsistOf("/some/dir"))
			})

			Context("when running rootless", func() {
				BeforeEach(func() {
					rootless = true
				})

				It("does not create the directory", func() {
					users.LookupReturns(&runrunc.ExecUser{Uid: 1012, Gid: 1013, Home: "/some/dir"}, nil)

					_, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{})
					Expect(err).NotTo(HaveOccurred())

					Expect(mkdirer.MkdirAsCallCount()).To(Equal(0))
				})
			})
		})

		Context("when the working directory creation fails", func() {
			It("returns an error", func() {
				mkdirer.MkdirAsReturns(errors.New("BOOOOOM"))
				_, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{})
				Expect(err).To(MatchError(ContainSubstring("create working directory: BOOOOOM")))
			})
		})
	})

	Context("when an ApparmorProfile is defined in the base process", func() {
		BeforeEach(func() {
			bundleLoader.LoadStub = func(path string) (goci.Bndl, error) {
				bndl := goci.Bundle()
				bndl = bndl.WithProcess(specs.Process{
					ApparmorProfile: "default-profile",
				})
				return bndl, nil
			}
		})

		It("should pass it to the process spec", func() {
			spec, err := preparer.Prepare(logger, bundlePath, garden.ProcessSpec{})
			Expect(err).ToNot(HaveOccurred())

			Expect(spec.Process.ApparmorProfile).To(Equal("default-profile"))
		})
	})
})

var _ = Describe("WaitWatcher", func() {
	It("calls Wait only once process.Wait returns", func() {
		waiter := new(fakes.FakeWaiter)
		waitReturns := make(chan struct{})
		waiter.WaitStub = func() (int, error) {
			<-waitReturns
			return 0, nil
		}

		runner := new(fakes.FakeRunner)

		watcher := runrunc.Watcher{}
		go watcher.OnExit(lagertest.NewTestLogger("test"), waiter, runner)

		Consistently(runner.RunCallCount).ShouldNot(Equal(1))
		close(waitReturns)
		Eventually(runner.RunCallCount).Should(Equal(1))
	})
})

var _ = Describe("RemoveFiles", func() {
	It("removes all the paths", func() {
		a := tmpFile("testremovefiles")
		b := tmpFile("testremovefiles")
		runrunc.RemoveFiles([]string{a, b}).Run(lagertest.NewTestLogger("test"))
		Expect(a).NotTo(BeAnExistingFile())
		Expect(b).NotTo(BeAnExistingFile())
	})
})

func tmpFile(name string) string {
	tmp, err := ioutil.TempFile("", name)
	Expect(err).NotTo(HaveOccurred())
	Expect(tmp.Close()).To(Succeed())
	return tmp.Name()
}
