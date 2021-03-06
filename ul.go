package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func splitter(s string, _count ...int) (string, string) {
	// default count to length of string
	var count int = len(s)

	// update count if necessary
	// should be positive and less than string length
	if len(_count) > 0 && _count[0] > 0 && _count[0] < count {
		count = _count[0]
	}

	// split into part to capitalize and remaining
	first := s[:count]
	remaining := ""
	if len(remaining) < len(s) {
		remaining = s[count:]
	}

	return first, remaining
}

// ToUpper return uppercase of the provided string using the
// count to determine number of characters to affect
// 0 and negative numbers affect the entire string.
func ToUpper(s string, _count ...int) string {

	capPart, remaining := splitter(s, _count...)
	capPart = strings.ToUpper(capPart)
	return strings.Join([]string{capPart, remaining}, "")
}

// ToUpper return lowercase of the provided string using the
// count to determine number of characters to affect
// 0 and negative numbers affect the entire string
func ToLower(s string, _count ...int) string {
	lowerPart, remaining := splitter(s, _count...)
	lowerPart = strings.ToLower(lowerPart)
	return strings.Join([]string{lowerPart, remaining}, "")
}

func run() {

	var usageText = func() {
		usageText := `
an exciting tool.

usage: cmdline [command] [options]

Commands:
	upper: uppercase a text
	lower: lowercase a text

Options:
	-text	text to affect
	-count	number of characters to affect

example:
	cmdline upper -text simple -count 4

--help, -h for more information about command
`
		_, _ = fmt.Fprintf(os.Stdout, "%s \n\n", usageText)
	}
	flag.Usage = usageText
	upperCmd := flag.NewFlagSet("upper", flag.ExitOnError)
	lowerCmd := flag.NewFlagSet("lower", flag.ExitOnError)

	// return if no arg is provided
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "upper":
		text := upperCmd.String("text", "", "text to uppercase")
		count := upperCmd.Int("count", 0, "no of characters to affect, -1, 0 for whole length")
		_ = upperCmd.Parse(os.Args[2:])

		newText := ToUpper(*text, *count)
		fmt.Println(newText)

	case "lower":
		text := lowerCmd.String("text", "", "text to lowercase")
		count := lowerCmd.Int("count", 0, "no of characters to affect, -1, 0 for whole length")

		_ = lowerCmd.Parse(os.Args[2:])

		newText := ToLower(*text, *count)
		fmt.Println(newText)
	default:
		flag.Usage()
	}
}

func main() {
	// .ul upper -text "something inside so strong" -count 22
	run()
}
