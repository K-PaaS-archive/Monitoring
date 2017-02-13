package helpers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDockerLifecycleHealthcheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docker-App-Lifecycle-Helpers Suite")
}
