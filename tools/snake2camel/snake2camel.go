package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func exitWithError(msg string) {
	err := fmt.Sprintf("error: %s\n", msg)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func replaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) (string, int) {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			if v[i] == -1 || v[i+1] == -1 {
				groups = append(groups, "")
			} else {
				groups = append(groups, str[v[i]:v[i+1]])
			}
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:], lastIndex
}

func snake2camel(content []byte) []byte {
	re := regexp.MustCompile(`(.*)([a-z]+[0-9]*)_([a-zA-Z0-9])(.*)`)
	res := string(content)
	last := -1

	for last != 0 {
		res, last = replaceAllStringSubmatchFunc(re, res, func(groups []string) string {
			return groups[1] + groups[2] + strings.ToUpper(groups[3]) + groups[4]
		})
	}

	return []byte(res)
}

func snake2camelFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	updated := snake2camel(content)

	return os.WriteFile(path, updated, 0)
}

func main() {
	if len(os.Args) != 2 {
		exitWithError("must be called with the file path as the only argument")
	}
	path := os.Args[1]
	if err := snake2camelFile(path); err != nil {
		exitWithError(err.Error())
	}
}
