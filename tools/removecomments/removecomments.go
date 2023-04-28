package main

import (
	"fmt"
	"os"
	"regexp"
)

// commentLineRegex matches a comment line and its newline characters. A
// comment line, for the purposes of this function, begins with optional
// whitespace followed by "//line" then any other text.
//
// optional newline character ---------------------------------+++
// optional carriage return --------------------------------+++|||
// end of line --------------------------------------------+||||||
// wildcard ---------------------------------------------++|||||||
// the literal word "line" --------------------------++++|||||||||
// comment delimiters -----------------------------++|||||||||||||
// one more more whitespace characters ---------+++|||||||||||||||
// beginning of line --------------------------+||||||||||||||||||
// allow multiline matching ---------------++++|||||||||||||||||||
//                                         |||||||||||||||||||||||
var commentLineRegex = regexp.MustCompile(`(?m)^\s*//line.*$\r?\n?`)

func removeComments(src []byte) []byte {
	return commentLineRegex.ReplaceAllLiteral(src, []byte(""))
}

func removeCommentsFromFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	updated := removeComments(content)
	return os.WriteFile(path, updated, 0)
}

func exitWithError(msg string) {
	err := fmt.Sprintf("error: %s\n", msg)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		exitWithError("must be called with the file path as the only argument")
	}
	path := os.Args[1]
	if err := removeCommentsFromFile(path); err != nil {
		exitWithError(err.Error())
	}
}
