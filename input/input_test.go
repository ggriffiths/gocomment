package input_test

import (
	"fmt"

	"github.build.ge.com/212472270/gocomment/input"
	"github.build.ge.com/212472270/gocomment/lint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	fileLines []lint.Line
	fileName  string
)

var _ = Describe("Input", func() {
	Describe("Given a user has ran gocomment", func() {
		Context("when a user passes in a golang source file", func() {
			It("should get the file's source code", func() {
				var err error
				osArgs := []string{fileName}
				fileName, err = input.ParseFlags(osArgs)
				Expect(err).NotTo(HaveOccurred())

				l1 := lint.Line{File: fileName, Number: 1, Text: "package comment"}
				l2 := lint.Line{File: fileName, Number: 2, Text: ""}
				l3 := lint.Line{File: fileName, Number: 3, Text: "// almostOkayFunction is a function with a comment that is just slightly too long."}
				l4 := lint.Line{File: fileName, Number: 4, Text: "func almostOkayFunction() {"}
				l5 := lint.Line{File: fileName, Number: 5, Text: "\treturn nil"}
				l6 := lint.Line{File: fileName, Number: 6, Text: "}"}

				expectedLines := []lint.Line{l1, l2, l3, l4, l5, l6}

				Expect(fileLines).To(Equal(expectedLines))
			})
		})

		Context("when a user passes in a golang source file with length flag", func() {
			It("should change the default max length to the argument provided", func() {
				osArgs := []string{"-l=5", fileName}
				fmt.Println("osArgs", osArgs)

				_, err := input.ParseFlags(osArgs)
				Expect(err).NotTo(HaveOccurred())
				Expect(lint.CommentMaxLength).To(Equal(5))
			})
		})
	})
})
