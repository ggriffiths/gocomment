package input_test

import (
	"github.build.ge.com/212472270/gocomment/input"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Input Suite")
}

var _ = BeforeSuite(func() {
	var err error
	fileName = "../test_comments/comment.go"
	fileLines, err = input.GetFileLines(fileName)
	Expect(err).NotTo(HaveOccurred())

})
