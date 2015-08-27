package webster_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWebster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Webster Suite")
}
