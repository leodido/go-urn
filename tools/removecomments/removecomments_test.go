package main

import (
	"fmt"
	"testing"
)

func TestRemoveComments(t *testing.T) {
	txt := `//line this should be deleted

This is a test file

it would contain go code

// and comments
    // and comments after spaces
		// and comments after tabs
//line but not this comment
    //line or this one after spaces
		//line or this one after tabs

and that is all
//line this should be deleted`

	expected := `
This is a test file

it would contain go code

// and comments
    // and comments after spaces
		// and comments after tabs

and that is all
`

	s := removeComments([]byte(txt))
	if string(s) != expected {
		fmt.Printf("expected length: %d, got %d\n", len(expected), len(s))
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, s)
	}
}
