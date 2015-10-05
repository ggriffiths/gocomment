package lint

import (
	"fmt"
	"strings"
)

var (
	// CommentMaxLength is a configurable max comment length that
	// can be set with the -l flag.
	CommentMaxLength int
)

//Line represents a line of code in a source file.
type Line struct {
	File   string
	Number int
	Text   string
}

//CommentShortEnough checks if a line is too long based on a configurable length.
func CommentShortEnough(comment Line) bool {
	if len(comment.Text) > CommentMaxLength {
		if strings.Contains(comment.Text[CommentMaxLength:], " ") {
			return false
		}
	}
	return true
}

//CommentHasPunctuation checks if a line ends in a period or not.
func CommentHasPunctuation(i int, comment Line, comments []Line) bool {
	if comment.Text[len(comment.Text)-1:] == "." {
		return true
	}

	if (i < len(comments)-1) &&
		(comments[i+1].Number == comment.Number+1) &&
		(comments[i+1].File == comment.File) {
		return CommentHasPunctuation(i+1, comments[i+1], comments)
	}

	return false
}

// GetComments returns takes a list of lines and returns those that are comments.
func GetComments(lines []Line) []Line {
	var comments []Line
	for _, line := range lines {
		if len(line.Text) > 1 && line.Text[0:2] == "//" {
			comments = append(comments, line)
		}
	}
	return comments
}

// FindBadComments performs all of our linting checks.
func FindBadComments(comments []Line) []string {
	var badComments []string
	var badComment string

	for i, comment := range comments {
		if !CommentShortEnough(comment) {
			badComment = fmt.Sprint(comment.File, ":", comment.Number, " line comment is too long.", "\n")
			badComments = append(badComments, badComment)

			fmt.Printf(badComment)
		}

		if !CommentHasPunctuation(i, comment, comments) {
			badComment = fmt.Sprint(comment.File, ":", comment.Number, " line comment must end in a period.", "\n")
			badComments = append(badComments, badComment)

			fmt.Printf(badComment)
		}
	}
	return badComments
}

// Run performs the linting by getting commments, finding bad ones, and prints
// those that do not meet gocomment standards. Returns 0 if all comments meet
// gocomment standards and 1 otherwise,
func Run(lines []Line) int {
	comments := GetComments(lines)
	problemLines := FindBadComments(comments)
	if len(problemLines) > 0 {
		return 1
	}

	for _, line := range problemLines {
		fmt.Println(line)
	}

	return 0
}
