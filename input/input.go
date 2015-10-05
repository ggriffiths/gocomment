package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.build.ge.com/212472270/gocomment/lint"
)

// GetFileLines returns the lines of a given source file or directory. It accepts
// regular expressions such as ./*.go or ./*/*.go
func GetFileLines(fileName string) ([]lint.Line, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var line lint.Line
	var lines []lint.Line
	lineNum := 1
	for scanner.Scan() {
		line = lint.Line{
			File:   fileName,
			Text:   scanner.Text(),
			Number: lineNum,
		}
		lines = append(lines, line)
		lineNum++
	}

	return lines, err
}

// ParseFlags checks which flags are set and handles them accordingly.
func ParseFlags(args []string) (string, error) {
	fileName := args[len(args)-1]
	maxLength := 0
	var err error
	for _, arg := range args {
		if strings.Contains(arg, "-l") {
			lengthStr := strings.TrimLeft(arg, "-l=")
			maxLength, err = strconv.Atoi(lengthStr)
		}
	}

	if maxLength < 1 {
		maxLength = 80
	}

	lint.CommentMaxLength = maxLength

	return fileName, err
}
