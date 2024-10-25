package logger

import (
	"regexp"
	"strings"
)

func Match(input string, pattern string) bool {
	var result strings.Builder

	for i, literal := range strings.Split(pattern, "*") {

		// Replace * with .*
		if i > 0 {
			result.WriteString(".*")
		}

		// Quote any regular expression meta characters in the
		// literal text.
		result.WriteString(regexp.QuoteMeta(literal))
	}

	exp, err := regexp.Compile(result.String())

	if err != nil {
		return false
	}

	return exp.MatchString(input)
}
