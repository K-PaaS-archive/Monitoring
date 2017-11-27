package store_test

import (
	"code.cloudfoundry.org/loggregator/doppler/internal/store"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceInfo", func() {
	var appService store.ServiceInfo

	BeforeEach(func() {
		appService = store.NewServiceInfo("app-id", "syslog://example.com:12345", "org.space.app.1")
	})

	It("should have an Id", func() {
		Expect(appService.Id()).NotTo(BeZero())
	})

	It("should have a different ID for different Urls", func() {
		otherAppService := store.NewServiceInfo("app-id", "syslog://example.com:12346", "org.space.app.1")
		Expect(appService.Id()).NotTo(Equal(otherAppService.Id()))
	})

	It("should return an ID that has no / or : characters", func() {
		Expect(appService.Id()).NotTo(MatchRegexp(`/`))
		Expect(appService.Id()).NotTo(MatchRegexp(`:`))
	})
})
