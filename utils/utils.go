package utils

import (
	"regexp"
)

func FormatComment(s string) string {
	str := removeNewLineCharacter(s)
	str = removeHeadSymbol(str)
	return str
}

func removeHeadSymbol(s string) string {
	// Remove comment symbols.
	rep := regexp.MustCompile(`^//( |)`)
	str := rep.ReplaceAllString(s, "")

	return str
}

func removeNewLineCharacter(s string) string {
	// Remove `\n//`.
	rep := regexp.MustCompile(`\n//( |)`)
	str := rep.ReplaceAllString(s, " ")
	return str
}
