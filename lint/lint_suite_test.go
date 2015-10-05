package lint_test

import (
	"testing"

	"github.build.ge.com/212472270/gocomment/lint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lint Suite")
}

var _ = BeforeSuite(func() {
	lint.CommentMaxLength = 80

})
