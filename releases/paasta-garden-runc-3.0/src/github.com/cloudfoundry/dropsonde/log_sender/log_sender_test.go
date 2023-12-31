package log_sender_test

import (
	"bytes"
	"errors"

	"github.com/cloudfoundry/dropsonde/emitter/fake"
	"github.com/cloudfoundry/dropsonde/log_sender"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/gogo/protobuf/proto"

	"io"
	"strings"
	"time"

	"github.com/cloudfoundry/dropsonde/metric_sender"
	"github.com/cloudfoundry/dropsonde/metricbatcher"
	"github.com/cloudfoundry/dropsonde/metrics"
	"github.com/cloudfoundry/loggregatorlib/loggertesthelper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LogSender", func() {
	var (
		emitter *fake.FakeEventEmitter
		sender  log_sender.LogSender
	)

	BeforeEach(func() {
		emitter = fake.NewFakeEventEmitter("origin")
		metricSender := metric_sender.NewMetricSender(emitter)
		batcher := metricbatcher.New(metricSender, time.Millisecond)
		metrics.Initialize(metricSender, batcher)
		sender = log_sender.NewLogSender(emitter, loggertesthelper.Logger())
	})

	AfterEach(func() {
		emitter.Close()
		for !emitter.IsClosed() {
			time.Sleep(10 * time.Millisecond)
		}
	})

	Describe("SendAppLog", func() {
		It("sends a log message event to its emitter", func() {
			err := sender.SendAppLog("app-id", "custom-log-message", "App", "0")
			Expect(err).NotTo(HaveOccurred())

			Expect(emitter.GetMessages()).To(HaveLen(1))
			log := emitter.GetMessages()[0].Event.(*events.LogMessage)
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_OUT))
			Expect(log.GetMessage()).To(BeEquivalentTo("custom-log-message"))
			Expect(log.GetAppId()).To(Equal("app-id"))
			Expect(log.GetSourceType()).To(Equal("App"))
			Expect(log.GetSourceInstance()).To(Equal("0"))
			Expect(log.GetTimestamp()).ToNot(BeNil())
		})

		It("totals number of log messages sent to emitter", func() {
			sender.SendAppLog("app-id", "custom-log-message", "App", "0")
			sender.SendAppLog("app-id", "custom-log-message", "App", "0")

			Eventually(emitter.GetEvents).Should(ContainElement(&events.CounterEvent{Name: proto.String("logSenderTotalMessagesRead"), Delta: proto.Uint64(2)}))
		})
	})

	Describe("SendAppErrorLog", func() {
		It("sends a log error message event to its emitter", func() {
			err := sender.SendAppErrorLog("app-id", "custom-log-error-message", "App", "0")
			Expect(err).NotTo(HaveOccurred())

			Expect(emitter.GetMessages()).To(HaveLen(1))
			log := emitter.GetMessages()[0].Event.(*events.LogMessage)
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_ERR))
			Expect(log.GetMessage()).To(BeEquivalentTo("custom-log-error-message"))
			Expect(log.GetAppId()).To(Equal("app-id"))
			Expect(log.GetSourceType()).To(Equal("App"))
			Expect(log.GetSourceInstance()).To(Equal("0"))
			Expect(log.GetTimestamp()).ToNot(BeNil())
		})

		It("totals number of log messages sent to emitter", func() {
			sender.SendAppErrorLog("app-id", "custom-log-message", "App", "0")
			sender.SendAppErrorLog("app-id", "custom-log-message", "App", "0")

			Eventually(emitter.GetEvents).Should(ContainElement(&events.CounterEvent{Name: proto.String("logSenderTotalMessagesRead"), Delta: proto.Uint64(2)}))
		})
	})

	Describe("counter emission", func() {
		It("emits on a timer", func() {
			sender.SendAppLog("app-id", "custom-log-message", "App", "0")
			Eventually(emitter.GetEvents).Should(ContainElement(&events.CounterEvent{Name: proto.String("logSenderTotalMessagesRead"), Delta: proto.Uint64(1)}))

			sender.SendAppLog("app-id", "custom-log-message", "App", "0")
			Eventually(emitter.GetEvents).Should(ContainElement(&events.CounterEvent{Name: proto.String("logSenderTotalMessagesRead"), Delta: proto.Uint64(1)}))
		})

		It("does not emit when no logs are written", func() {
			Consistently(emitter.GetEvents, 1).Should(HaveLen(0))
		})
	})

	Context("when messages cannot be emitted", func() {
		BeforeEach(func() {
			emitter.ReturnError = errors.New("expected error")
		})

		Describe("SendAppLog", func() {
			It("sends an error when log messages cannot be emitted", func() {
				err := sender.SendAppLog("app-id", "custom-log-message", "App", "0")
				Expect(err).To(HaveOccurred())
			})

		})

		Describe("SendAppErrorLog", func() {
			It("sends an error when log error messages cannot be emitted", func() {
				err := sender.SendAppErrorLog("app-id", "custom-log-error-message", "App", "0")
				Expect(err).To(HaveOccurred())
			})

		})
	})

	Describe("ScanLogStream", func() {

		It("sends lines from stream to emitter", func() {
			buf := bytes.NewBufferString("line 1\nline 2\n")

			sender.ScanLogStream("someId", "app", "0", buf)

			messages := emitter.GetMessages()
			Expect(messages).To(HaveLen(2))

			log := emitter.GetMessages()[0].Event.(*events.LogMessage)
			Expect(log.GetMessage()).To(BeEquivalentTo("line 1"))
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_OUT))
			Expect(log.GetAppId()).To(Equal("someId"))
			Expect(log.GetSourceType()).To(Equal("app"))
			Expect(log.GetSourceInstance()).To(Equal("0"))

			log = emitter.GetMessages()[1].Event.(*events.LogMessage)
			Expect(log.GetMessage()).To(BeEquivalentTo("line 2"))
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_OUT))
		})

		It("logs a message and returns on read errors", func() {
			var errReader fakeReader
			sender.ScanLogStream("someId", "app", "0", &errReader)

			messages := emitter.GetMessages()
			Expect(messages).To(HaveLen(1))

			log := emitter.GetMessages()[0].Event.(*events.LogMessage)
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_OUT))
			Expect(log.GetMessage()).To(BeEquivalentTo("one"))

			loggerMessage := loggertesthelper.TestLoggerSink.LogContents()
			Expect(loggerMessage).To(ContainSubstring("Read Error"))
		})

		It("stops when reader returns EOF", func() {
			var reader infiniteReader
			reader.stopChan = make(chan struct{})
			doneChan := make(chan struct{})

			go func() {
				sender.ScanLogStream("someId", "app", "0", reader)
				close(doneChan)
			}()

			Eventually(func() int { return len(emitter.GetMessages()) }).Should(BeNumerically(">", 1))
			close(reader.stopChan)
			Eventually(doneChan).Should(BeClosed())
		})

		It("drops over-length messages and resumes scanning", func() {
			// Scanner can't handle tokens over 64K
			bigReader := strings.NewReader(strings.Repeat("x", 64*1024+1) + "\nsmall message\n")

			doneChan := make(chan struct{})
			go func() {
				sender.ScanLogStream("someId", "app", "0", bigReader)
				close(doneChan)
			}()

			Eventually(emitter.GetMessages).Should(HaveLen(4))

			Eventually(doneChan).Should(BeClosed())

			messages := getLogMessages(emitter.GetMessages())

			Expect(messages[0]).To(ContainSubstring("Dropped log message: message too long (>64K without a newline)"))
			Expect(messages[1]).To(Equal("x"))
			Expect(messages[2]).To(Equal("small message"))
		})

		It("ignores empty lines", func() {
			reader := strings.NewReader("one\n\ntwo\n")

			sender.ScanLogStream("someId", "app", "0", reader)

			Expect(emitter.GetMessages()).To(HaveLen(2))
			messages := getLogMessages(emitter.GetMessages())

			Expect(messages[0]).To(Equal("one"))
			Expect(messages[1]).To(Equal("two"))
		})
	})

	Describe("ScanErrorLogStream", func() {

		It("sends lines from stream to emitter", func() {
			buf := bytes.NewBufferString("line 1\nline 2\n")

			sender.ScanErrorLogStream("someId", "app", "0", buf)

			messages := emitter.GetMessages()
			Expect(messages).To(HaveLen(2))

			log := emitter.GetMessages()[0].Event.(*events.LogMessage)
			Expect(log.GetMessage()).To(BeEquivalentTo("line 1"))
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_ERR))
			Expect(log.GetAppId()).To(Equal("someId"))
			Expect(log.GetSourceType()).To(Equal("app"))
			Expect(log.GetSourceInstance()).To(Equal("0"))

			log = emitter.GetMessages()[1].Event.(*events.LogMessage)
			Expect(log.GetMessage()).To(BeEquivalentTo("line 2"))
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_ERR))
		})

		It("logs a message and stops on read errors", func() {
			var errReader fakeReader
			sender.ScanErrorLogStream("someId", "app", "0", &errReader)

			messages := emitter.GetMessages()
			Expect(messages).To(HaveLen(1))

			log := emitter.GetMessages()[0].Event.(*events.LogMessage)
			Expect(log.GetMessageType()).To(Equal(events.LogMessage_ERR))
			Expect(log.GetMessage()).To(BeEquivalentTo("one"))

			loggerMessage := loggertesthelper.TestLoggerSink.LogContents()
			Expect(loggerMessage).To(ContainSubstring("Read Error"))
		})

		It("stops when reader returns EOF", func() {
			var reader infiniteReader
			reader.stopChan = make(chan struct{})
			doneChan := make(chan struct{})

			go func() {
				sender.ScanErrorLogStream("someId", "app", "0", reader)
				close(doneChan)
			}()

			Eventually(func() int { return len(emitter.GetMessages()) }).Should(BeNumerically(">", 1))

			close(reader.stopChan)
			Eventually(doneChan).Should(BeClosed())

		})

		It("drops over-length messages and resumes scanning", func() {
			// Scanner can't handle tokens over 64K
			bigReader := strings.NewReader(strings.Repeat("x", 64*1024+1) + "\nsmall message\n")
			sender.ScanErrorLogStream("someId", "app", "0", bigReader)
			var lenMessages = func() int { return len(emitter.GetMessages()) }
			Eventually(lenMessages).Should(BeNumerically(">=", 3))

			messages := getLogMessages(emitter.GetMessages())

			Expect(messages[0]).To(ContainSubstring("Dropped log message: message too long (>64K without a newline)"))
			Expect(messages[1]).To(Equal("x"))
			Expect(messages[2]).To(Equal("small message"))
		})

		It("ignores empty lines", func() {
			reader := strings.NewReader("one\n \ntwo\n")

			sender.ScanErrorLogStream("someId", "app", "0", reader)

			Expect(emitter.GetMessages()).To(HaveLen(2))
			messages := getLogMessages(emitter.GetMessages())

			Expect(messages[0]).To(Equal("one"))
			Expect(messages[1]).To(Equal("two"))
		})
	})
})

type fakeReader struct {
	counter int
}

func (f *fakeReader) Read(p []byte) (int, error) {
	f.counter++

	switch f.counter {
	case 1: // message
		return copy(p, "one\n"), nil
	case 2: // read error
		return 0, errors.New("Read Error")
	case 3: // message
		return copy(p, "two\n"), nil
	default: // eof
		return 0, io.EOF
	}
}

type infiniteReader struct {
	stopChan chan struct{}
}

func (i infiniteReader) Read(p []byte) (int, error) {
	select {
	case <-i.stopChan:
		return 0, io.EOF
	default:
	}

	return copy(p, "hello\n"), nil
}

func getLogMessages(messages []fake.Message) []string {
	var logMessages []string
	for _, msg := range messages {
		log, ok := msg.Event.(*events.LogMessage)
		if ok {
			logMessages = append(logMessages, string(log.GetMessage()))
		}
	}
	return logMessages
}
