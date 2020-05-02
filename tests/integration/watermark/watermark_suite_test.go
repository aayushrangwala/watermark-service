package watermark_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWatermark(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Watermark Suite")
}
