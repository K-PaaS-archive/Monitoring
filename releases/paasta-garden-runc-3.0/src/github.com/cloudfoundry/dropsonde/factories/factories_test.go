package factories_test

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/url"

	"github.com/cloudfoundry/dropsonde/factories"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/gogo/protobuf/proto"
)

var _ = Describe("HTTP event creation", func() {
	var (
		applicationId *uuid.UUID
		requestId     *uuid.UUID
		req           *http.Request
		reqMethod     string
	)

	BeforeEach(func() {
		applicationId, _ = uuid.NewV4()
		requestId, _ = uuid.NewV4()
		reqMethod = "GET"
	})

	JustBeforeEach(func() {
		req, _ = http.NewRequest(reqMethod, "http://foo.example.com/", nil)

		// According to the godoc for http.Request, server requests only contain
		// Path and RawQuery fields.
		req.URL = &url.URL{
			Path: "/",
		}

		req.RemoteAddr = "127.0.0.1"
		req.Header.Set("User-Agent", "our-testing-client")
	})

	Describe("NewHttpStartStop", func() {

		It("should extract X-Forwarded-For from header", func() {
			forwardedAddress := "123.123.123.123"
			req.Header.Set("X-Forwarded-For", forwardedAddress)

			startStopEvent := factories.NewHttpStartStop(req, http.StatusOK, 3, events.PeerType_Server, requestId)
			Expect(startStopEvent.GetUri()).To(Equal(fmt.Sprintf("http://%s/", forwardedAddress)))
		})

		It("should extract ApplicationId from request header", func() {
			applicationId, _ := uuid.NewV4()
			req.Header.Set("X-CF-ApplicationID", applicationId.String())

			startStopEvent := factories.NewHttpStartStop(req, http.StatusOK, 3, events.PeerType_Server, requestId)
			Expect(startStopEvent.GetApplicationId()).To(Equal(factories.NewUUID(applicationId)))
		})

		It("should extract InstanceIndex from request header", func() {
			instanceIndex := "1"
			req.Header.Set("X-CF-InstanceIndex", instanceIndex)

			startStopEvent := factories.NewHttpStartStop(req, http.StatusOK, 3, events.PeerType_Server, requestId)
			Expect(startStopEvent.GetInstanceIndex()).To(BeNumerically("==", 1))
		})

		It("should extract InstanceID from request header", func() {
			instanceId := "fake-id"
			req.Header.Set("X-CF-InstanceID", instanceId)

			startStopEvent := factories.NewHttpStartStop(req, http.StatusOK, 3, events.PeerType_Server, requestId)
			Expect(startStopEvent.GetInstanceId()).To(Equal(instanceId))
		})

		It("should set appropriate fields", func() {
			expectedEvent := &events.HttpStartStop{
				RequestId:     factories.NewUUID(requestId),
				PeerType:      events.PeerType_Server.Enum(),
				Method:        events.Method_GET.Enum(),
				Uri:           proto.String("http://foo.example.com/"),
				RemoteAddress: proto.String("127.0.0.1"),
				UserAgent:     proto.String("our-testing-client"),
				StatusCode:    proto.Int32(http.StatusOK),
				ContentLength: proto.Int64(1234),
			}

			event := factories.NewHttpStartStop(req, http.StatusOK, 1234, events.PeerType_Server, requestId)

			Expect(event.GetStartTimestamp()).ToNot(BeZero())
			event.StartTimestamp = nil
			Expect(event.GetStopTimestamp()).ToNot(BeZero())
			event.StopTimestamp = nil

			Expect(event).To(Equal(expectedEvent))
		})
	})

	Describe("NewHttpStart", func() {

		Context("without an application ID or instanceIndex", func() {

			It("should extract X-Forwarded-For from header", func() {
				forwardedAddress := "123.123.123.123"
				req.Header.Set("X-Forwarded-For", forwardedAddress)

				startEvent := factories.NewHttpStart(req, events.PeerType_Server, requestId)
				Expect(startEvent.GetUri()).To(Equal(fmt.Sprintf("http://%s/", forwardedAddress)))
			})

			It("should set appropriate fields", func() {
				expectedStartEvent := &events.HttpStart{
					RequestId:     factories.NewUUID(requestId),
					PeerType:      events.PeerType_Server.Enum(),
					Method:        events.Method_GET.Enum(),
					Uri:           proto.String("http://foo.example.com/"),
					RemoteAddress: proto.String("127.0.0.1"),
					UserAgent:     proto.String("our-testing-client"),
				}

				startEvent := factories.NewHttpStart(req, events.PeerType_Server, requestId)

				Expect(startEvent.GetTimestamp()).ToNot(BeZero())
				startEvent.Timestamp = nil

				Expect(startEvent).To(Equal(expectedStartEvent))
			})
		})

		Context("with an application ID", func() {
			It("should include it in the start event", func() {
				applicationId, _ := uuid.NewV4()
				req.Header.Set("X-CF-ApplicationID", applicationId.String())

				startEvent := factories.NewHttpStart(req, events.PeerType_Server, requestId)

				Expect(startEvent.GetApplicationId()).To(Equal(factories.NewUUID(applicationId)))
			})
		})

		Context("with an application instance index", func() {
			It("should include it in the start event", func() {
				req.Header.Set("X-CF-InstanceIndex", "1")

				startEvent := factories.NewHttpStart(req, events.PeerType_Server, requestId)

				Expect(startEvent.GetInstanceIndex()).To(BeNumerically("==", 1))
			})
		})

		Context("with an application instance ID", func() {
			It("should include it in the start event", func() {
				req.Header.Set("X-CF-InstanceID", "fake-id")

				startEvent := factories.NewHttpStart(req, events.PeerType_Server, requestId)

				Expect(startEvent.GetInstanceId()).To(Equal("fake-id"))
			})
		})

		Context("with other HTTP methods", func() {
			BeforeEach(func() {
				reqMethod = "PATCH"
			})

			It("sends the other method through", func() {
				startEvent := factories.NewHttpStart(req, events.PeerType_Server, requestId)
				Expect(startEvent.GetMethod()).To(Equal(events.Method_PATCH))
			})
		})
	})

	Describe("NewHttpStop", func() {

		It("should extract X-Forwarded-For from header", func() {
			forwardedAddress := "123.123.123.123"
			req.Header.Set("X-Forwarded-For", forwardedAddress)

			stopEvent := factories.NewHttpStop(req, http.StatusOK, 3, events.PeerType_Server, requestId)
			Expect(stopEvent.GetUri()).To(Equal(fmt.Sprintf("http://%s/", forwardedAddress)))
		})

		It("should set appropriate fields", func() {
			req.Header.Set("X-CF-ApplicationID", applicationId.String())
			expectedStopEvent := &events.HttpStop{
				ApplicationId: factories.NewUUID(applicationId),
				RequestId:     factories.NewUUID(requestId),
				Uri:           proto.String("http://foo.example.com/"),
				PeerType:      events.PeerType_Server.Enum(),
				StatusCode:    proto.Int32(200),
				ContentLength: proto.Int64(3),
			}

			stopEvent := factories.NewHttpStop(req, 200, 3, events.PeerType_Server, requestId)

			Expect(stopEvent.GetTimestamp()).ToNot(BeZero())
			stopEvent.Timestamp = nil

			Expect(stopEvent).To(Equal(expectedStopEvent))
		})
	})

	Describe("NewLogMessage", func() {
		It("should set appropriate fields", func() {
			expectedLogEvent := &events.LogMessage{
				Message:     []byte("hello"),
				AppId:       proto.String("app-id"),
				MessageType: events.LogMessage_OUT.Enum(),
				SourceType:  proto.String("App"),
			}

			logEvent := factories.NewLogMessage(events.LogMessage_OUT, "hello", "app-id", "App")

			Expect(logEvent.GetTimestamp()).ToNot(BeZero())
			logEvent.Timestamp = nil

			Expect(logEvent).To(Equal(expectedLogEvent))
		})
	})

	Describe("NewContainerMetric", func() {
		It("should set the appropriate fields", func() {
			expectedContainerMetric := &events.ContainerMetric{
				ApplicationId: proto.String("some_app_id"),
				InstanceIndex: proto.Int32(7),
				CpuPercentage: proto.Float64(42.24),
				MemoryBytes:   proto.Uint64(1234),
				DiskBytes:     proto.Uint64(13231231),
			}

			containerMetric := factories.NewContainerMetric("some_app_id", 7, 42.24, 1234, 13231231)

			Expect(containerMetric).To(Equal(expectedContainerMetric))
		})
	})
})
