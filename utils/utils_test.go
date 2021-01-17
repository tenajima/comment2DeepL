package utils

import (
	"log"
	"testing"
)

func TestRemoveHeadCommentSymbol(t *testing.T) {
	result := removeHeadSymbol("// This is test.")
	if result != "This is test." {
		log.Fatal("error")
	}

	result = removeHeadSymbol("//This is test.")
	if result != "This is test." {
		log.Fatal("error")
	}
}

func TestRemoveNewLineCharacter(t *testing.T) {
	s := `This is first line and
// this is second line.`
	result := removeNewLineCharacter(s)
	if result != "This is first line and this is second line." {
		log.Fatal(result)
	}
}

func TestFormatComment(t *testing.T) {
	s := `// This is first line and
// this is second line.`
	result := FormatComment(s)
	if result != "This is first line and this is second line." {
		log.Fatal(result)
	}
}
