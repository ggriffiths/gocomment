package lint_test

import (
	"github.build.ge.com/212472270/gocomment/input"
	"github.build.ge.com/212472270/gocomment/lint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Input", func() {

	Describe("Given lint has been ran with valid flags on a golang file", func() {
		Context("when said file has any line comments with incorrect format", func() {
			It("should return 1", func() {
				lines, err := input.GetFileLines("../test_comments/badcomments.go")
				Expect(err).NotTo(HaveOccurred())

				Expect(lint.Run(lines)).To(Equal(1))
			})
		})

		Context("when said file has no line comments with incorrect format", func() {
			It("should return 0", func() {
				lines, err := input.GetFileLines("../test_comments/goodcomments.go")
				Expect(err).NotTo(HaveOccurred())

				Expect(lint.Run(lines)).To(Equal(0))
			})
		})
	})

	Describe("Given gocomment has valid flags sent to it and comments have been properly extracted", func() {
		Context("when a line comment is longer than the default value of 80", func() {
			It("should identify that the line comment is too long", func() {
				line := lint.Line{
					File:   "source.go",
					Number: 5,
					Text:   "// ThisComment represents a bad comment due to it's length being pretty pretty pretty long.",
				}
				Expect(lint.CommentShortEnough(line)).To(Equal(false))
			})
		})
		Context("when a line comment is shorter than the default value of 80", func() {
			It("should identify that the line comment is not too long", func() {
				line := lint.Line{
					File:   "source.go",
					Number: 3,
					Text:   "// Knock knock. Who's there? Go Fmt Yourself.",
				}

				Expect(lint.CommentShortEnough(line)).To(Equal(true))
			})
		})

		Context("when a line comment has no period at the end", func() {
			It("should identify that the line comment does not have punctuation", func() {
				line := lint.Line{
					File:   "source.go",
					Number: 3,
					Text:   "// LineComment",
				}

				Expect(lint.CommentHasPunctuation(1, line, []lint.Line{line})).To(Equal(false))
			})
		})

		Context("when a line comment has a period at the end", func() {
			It("should identify that the line comment has punctuation", func() {
				line := lint.Line{
					File:   "source.go",
					Number: 3,
					Text:   "// LineComment.",
				}

				Expect(lint.CommentHasPunctuation(1, line, []lint.Line{line})).To(Equal(true))
			})
		})
	})
})
