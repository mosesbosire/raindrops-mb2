package raindrops_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRaindrops(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Raindrops Suite")
}
