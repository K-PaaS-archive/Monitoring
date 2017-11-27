package groupedsinks_test

import (
	"net"
	"net/url"
	"sync"
	"time"

	"code.cloudfoundry.org/loggregator/doppler/internal/groupedsinks"
	"code.cloudfoundry.org/loggregator/doppler/internal/sinks"
	"code.cloudfoundry.org/loggregator/doppler/internal/syslog"
	"code.cloudfoundry.org/loggregator/metricemitter/testhelper"
	"github.com/cloudfoundry/dropsonde/emitter"
	"github.com/cloudfoundry/dropsonde/factories"
	"github.com/cloudfoundry/sonde-go/events"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupedSink", func() {
	var groupedSinks *groupedsinks.GroupedSinks
	var inputChan chan *events.Envelope

	BeforeEach(func() {
		groupedSinks = groupedsinks.NewGroupedSinks(
			&spyMetricBatcher{},
			testhelper.NewMetricClient(),
		)
		inputChan = make(chan *events.Envelope, 10)
	})

	Describe("Broadcast", func() {
		Context("when all pre-existing firehose connections have been deleted", func() {
			It("sends message to all registered app sinks", func() {
				appSink := syslog.NewSyslogSink("123", &url.URL{Host: "url"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
				appSinkInputChan := make(chan *events.Envelope, 10)
				groupedSinks.RegisterAppSink(appSinkInputChan, appSink)

				msg, _ := emitter.Wrap(factories.NewLogMessage(events.LogMessage_OUT, "test message", "123", "App"), "origin")
				groupedSinks.Broadcast("123", msg)

				Eventually(appSinkInputChan).Should(Receive(Equal(msg)))
			})
		})

		It("sends message to all registered sinks that match the appId", func() {
			appId := "123"
			appSink := syslog.NewSyslogSink("123", &url.URL{Host: "url"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			otherInputChan := make(chan *events.Envelope)
			groupedSinks.RegisterAppSink(otherInputChan, appSink)

			appId = "789"
			appSink = syslog.NewSyslogSink(appId, &url.URL{Host: "url"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, appSink)

			msg, _ := emitter.Wrap(factories.NewLogMessage(events.LogMessage_OUT, "test message", appId, "App"), "origin")
			groupedSinks.Broadcast(appId, msg)

			Eventually(inputChan).Should(Receive(Equal(msg)))
			Expect(otherInputChan).To(HaveLen(0))
		})

		It("does not block when sending to an appId that has no sinks", func(done Done) {
			appId := "NonExistantApp"
			msg, _ := emitter.Wrap(factories.NewLogMessage(events.LogMessage_OUT, "test message", appId, "App"), "origin")
			groupedSinks.Broadcast(appId, msg)
			close(done)
		})

		It("does not block when sending to slow sink", func() {
			appId := "syslog-a"
			fakeSink1A := &fakeSink{sinkId: "sink1", appId: appId}
			inputChan1A := make(chan *events.Envelope)
			groupedSinks.RegisterAppSink(inputChan1A, fakeSink1A)

			c := make(chan struct{})
			go func() {
				defer close(c)
				msg, _ := emitter.Wrap(factories.NewLogMessage(events.LogMessage_OUT, "test message", appId, "App"), "origin")
				groupedSinks.Broadcast(appId, msg)
			}()

			Eventually(c).Should(BeClosed())
		})
	})

	Describe("BroadcastError", func() {
		It("sends message to all registered sinks that match the appId", func() {
			appId := "123"
			health := newSpyHealthRegistrar()
			appSink := sinks.NewDumpSink(appId, 10, time.Second, health)
			otherInputChan := make(chan *events.Envelope, 1)
			groupedSinks.RegisterAppSink(otherInputChan, appSink)

			appId = "789"
			appSink = sinks.NewDumpSink(appId, 10, time.Second, health)

			groupedSinks.RegisterAppSink(inputChan, appSink)
			msg, _ := emitter.Wrap(factories.NewLogMessage(events.LogMessage_OUT, "error message", appId, "App"), "origin")
			groupedSinks.BroadcastError(appId, msg)

			Eventually(inputChan).Should(Receive(Equal(msg)))
			Expect(otherInputChan).To(HaveLen(0))
		})

		It("does not send to sinks that don't want errors", func() {
			appId := "789"

			health := newSpyHealthRegistrar()
			sink1 := sinks.NewDumpSink(appId, 10, time.Second, health)
			sink2 := syslog.NewSyslogSink(appId, &url.URL{Host: "url"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, sink1)
			groupedSinks.RegisterAppSink(inputChan, sink2)
			msg, _ := emitter.Wrap(factories.NewLogMessage(events.LogMessage_OUT, "error message", appId, "App"), "origin")
			groupedSinks.BroadcastError(appId, msg)
			Expect(<-inputChan).To(Equal(msg))
			Expect(inputChan).To(HaveLen(0))
		})

		It("does not block when sending to slow sink", func() {
			appId := "syslog-a"
			fakeSink1A := &fakeSink{sinkId: "sink1", appId: appId, shouldRxErrors: true}
			inputChan1A := make(chan *events.Envelope)
			groupedSinks.RegisterAppSink(inputChan1A, fakeSink1A)

			c := make(chan struct{})
			go func() {
				defer close(c)
				msg, _ := emitter.Wrap(factories.NewLogMessage(events.LogMessage_OUT, "test message", appId, "App"), "origin")
				groupedSinks.BroadcastError(appId, msg)
			}()

			Eventually(c).Should(BeClosed())
		})
	})

	Describe("Register", func() {
		It("returns false for empty app ids", func() {
			appId := ""
			appSink := syslog.NewSyslogSink(appId, &url.URL{Host: "url"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			result := groupedSinks.RegisterAppSink(inputChan, appSink)
			Expect(result).To(BeFalse())
		})

		It("returns false for empty identifiers", func() {
			appId := "appId"
			appSink := syslog.NewSyslogSink(appId, &url.URL{Host: ""}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			result := groupedSinks.RegisterAppSink(inputChan, appSink)
			Expect(result).To(BeFalse())
		})

		It("returns false when registering a duplicate", func() {
			appId := "789"
			appSink := syslog.NewSyslogSink(appId, &url.URL{Host: "url"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			groupedSinks.RegisterAppSink(inputChan, appSink)
			result := groupedSinks.RegisterAppSink(inputChan, appSink)
			Expect(result).To(BeFalse())
		})
	})

	Describe("CloseAndDelete", func() {
		It("only deletes a specific sink", func() {
			target := "789"

			sink1 := syslog.NewSyslogSink(target, &url.URL{Host: "url1"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			sink2 := syslog.NewSyslogSink(target, &url.URL{Host: "url2"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, sink1)
			groupedSinks.RegisterAppSink(inputChan, sink2)

			ok := groupedSinks.CloseAndDelete(sink1)
			Expect(ok).To(BeTrue())
		})

		It("handle deletes for non-existing appIds", func() {
			target := "789"
			sink1 := syslog.NewSyslogSink(target, &url.URL{Host: "url1"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			ok := groupedSinks.CloseAndDelete(sink1)
			Expect(ok).To(BeFalse())
		})

		It("handle deletes for existing appIds but unregistered drain URLs", func() {
			target := "789"

			sink1 := syslog.NewSyslogSink(target, &url.URL{Host: "url1"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			sink2 := syslog.NewSyslogSink(target, &url.URL{Host: "url2"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, sink1)

			ok := groupedSinks.CloseAndDelete(sink2)
			Expect(ok).To(BeFalse())
		})

		It("closes the inputChan", func() {
			target := "789"
			sink := syslog.NewSyslogSink(target, &url.URL{Host: "url1"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, sink)
			groupedSinks.CloseAndDelete(sink)
			Expect(inputChan).To(BeClosed())
		})

	})

	Describe("DeleteAll", func() {
		It("removes all the sinks", func() {
			sink := &fakeSink{sinkId: "sink1", appId: "app1"}
			groupedSinks.RegisterAppSink(make(chan *events.Envelope), sink)

			groupedSinks.DeleteAll()

			Expect(groupedSinks.DrainFor("app1", "sink1")).To(BeNil())
		})

		It("closes all the sinks input chans", func() {
			sink1 := &fakeSink{sinkId: "sink1", appId: "app1"}

			groupedSinks.RegisterAppSink(inputChan, sink1)

			groupedSinks.DeleteAll()

			Expect(inputChan).To(BeClosed())
		})
	})

	Describe("DrainFor", func() {
		It("returns only sinks that match the appid and drain URL", func() {
			target := "789"

			sink1 := syslog.NewSyslogSink(target, &url.URL{Scheme: "syslog", Host: "other sink"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			sink2 := syslog.NewSyslogSink(target, &url.URL{Scheme: "syslog", Host: "sink we are searching for"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, sink1)
			groupedSinks.RegisterAppSink(inputChan, sink2)

			sinkDrain := groupedSinks.DrainFor(target, "syslog://sink we are searching for")
			Expect(sinkDrain).To(Equal(sink2))
		})

		It("returns nil if no drains are registered", func() {
			target := "789"
			sink := syslog.NewSyslogSink(target, &url.URL{Host: "url2"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, sink)

			Expect(groupedSinks.DrainFor(target, "url1")).To(BeNil())
		})

		It("return nils if no drains exist", func() {
			Expect(groupedSinks.DrainFor("empty", "empty")).To(BeNil())
		})
	})

	Describe("DumpFor", func() {
		It("returns only dumps", func() {
			appId := "789"
			sink1 := syslog.NewSyslogSink(appId, &url.URL{Host: "url1"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			sink2 := syslog.NewSyslogSink(appId, &url.URL{Host: "url2"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")
			health := newSpyHealthRegistrar()
			sink3 := sinks.NewDumpSink(appId, 5, time.Second, health)

			groupedSinks.RegisterAppSink(inputChan, sink1)
			groupedSinks.RegisterAppSink(inputChan, sink2)
			groupedSinks.RegisterAppSink(inputChan, sink3)

			Expect(groupedSinks.DumpFor(appId)).To(Equal(sink3))
		})

		It("returns only dumps that match the appId", func() {
			appId := "789"
			otherAppId := "790"

			health := newSpyHealthRegistrar()
			sink1 := sinks.NewDumpSink(appId, 5, time.Second, health)
			sink2 := sinks.NewDumpSink(otherAppId, 5, time.Second, health)

			groupedSinks.RegisterAppSink(inputChan, sink1)
			groupedSinks.RegisterAppSink(inputChan, sink2)

			Expect(groupedSinks.DumpFor(appId)).To(Equal(sink1))
		})

		It("returns nil if no dumps are registered", func() {
			target := "789"

			sink1 := syslog.NewSyslogSink(target, &url.URL{Host: "url1"}, 100, DummySyslogWriter{}, dummyErrorHandler, "dropsonde-origin")

			groupedSinks.RegisterAppSink(inputChan, sink1)

			Expect(groupedSinks.DumpFor(target)).To(BeNil())
		})

		It("returns nil if no sinks exist", func() {
			Expect(groupedSinks.DumpFor("empty")).To(BeNil())
		})
	})

	Describe("ContainerMetricsFor", func() {
		It("returns only container metric sinks", func() {
			appId := "456"

			health := newSpyHealthRegistrar()
			sink1 := sinks.NewContainerMetricSink(appId, 1*time.Second, time.Second, health)
			sink2 := sinks.NewDumpSink(appId, 5, time.Second, health)

			groupedSinks.RegisterAppSink(inputChan, sink1)
			groupedSinks.RegisterAppSink(inputChan, sink2)

			Expect(groupedSinks.ContainerMetricsFor(appId)).To(Equal(sink1))
		})

		It("returns only container metrics for appId", func() {
			appId1 := "123"
			appId2 := "456"

			health := newSpyHealthRegistrar()
			sink1 := sinks.NewContainerMetricSink(appId1, 1*time.Second, time.Second, health)
			sink2 := sinks.NewContainerMetricSink(appId2, 1*time.Second, time.Second, health)

			groupedSinks.RegisterAppSink(inputChan, sink1)
			groupedSinks.RegisterAppSink(inputChan, sink2)

			Expect(groupedSinks.ContainerMetricsFor(appId1)).To(Equal(sink1))
		})

		It("returns nil if no container metrics sinks are registered", func() {
			appId := "1234"
			health := newSpyHealthRegistrar()
			sink2 := sinks.NewDumpSink(appId, 5, time.Second, health)
			groupedSinks.RegisterAppSink(inputChan, sink2)

			Expect(groupedSinks.ContainerMetricsFor(appId)).To(BeNil())
		})

		It("returns nil if no sinks exist", func() {
			Expect(groupedSinks.ContainerMetricsFor("1234")).To(BeNil())
		})

	})
})

func dummyErrorHandler(_, _ string) {}

type DummySyslogWriter struct{}

func (d DummySyslogWriter) Connect() error { return nil }
func (d DummySyslogWriter) Write(p int, b []byte, source, sourceId string, timestamp int64) (int, error) {
	return 0, nil
}
func (d DummySyslogWriter) Close() error { return nil }

type fakeMessageWriter struct {
	RemoteAddress string
}

func (fake *fakeMessageWriter) RemoteAddr() net.Addr {
	return fakeAddr{remoteAddress: fake.RemoteAddress}
}

func (fake *fakeMessageWriter) WriteMessage(messageType int, data []byte) error {
	return nil
}

func (fake *fakeMessageWriter) SetWriteDeadline(t time.Time) error {
	return nil
}

type fakeAddr struct {
	remoteAddress string
}

func (fake fakeAddr) Network() string {
	return "RemoteAddressNetwork"
}

func (fake fakeAddr) String() string {
	return fake.remoteAddress
}

type SpyHealthRegistrar struct {
	mu     sync.Mutex
	values map[string]float64
}

func newSpyHealthRegistrar() *SpyHealthRegistrar {
	return &SpyHealthRegistrar{
		values: make(map[string]float64),
	}
}

func (s *SpyHealthRegistrar) Inc(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[name]++
}

func (s *SpyHealthRegistrar) Dec(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[name]--
}

func (s *SpyHealthRegistrar) Get(name string) float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.values[name]
}

type fakeSink struct {
	sinkId         string
	appId          string
	shouldRxErrors bool
}

func (f *fakeSink) AppID() string {
	return f.appId
}

func (f *fakeSink) Run(<-chan *events.Envelope) {
}

func (f *fakeSink) ShouldReceiveErrors() bool {
	return f.shouldRxErrors
}

func (f *fakeSink) Identifier() string {
	return f.sinkId
}

func (f *fakeSink) GetInstrumentationMetric() sinks.Metric {
	return sinks.Metric{}
}

type spyMetricBatcher struct{}

func (s *spyMetricBatcher) BatchIncrementCounter(name string) {}
