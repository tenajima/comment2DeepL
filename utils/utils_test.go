package utils

import (
	"log"
	"testing"
)

func TestRemoveHeadCommentSymbol(t *testing.T){
	result, _:= RemoveHeadCommentSymbol("//This is test.")
	if result != "This is test." {
		log.Fatal("error")
	}
}

