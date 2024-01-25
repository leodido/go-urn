package main

import (
	"fmt"
	"testing"
)

func TestSnake2CamelNoReplace(t *testing.T) {
	txt := `const start int = 1`

	exp := `const start int = 1`

	got := snake2camel([]byte(txt))
	if string(got) != exp {
		fmt.Printf("expected length: %d, got %d\n", len(exp), len(got))
		t.Fatalf("expected:\n%s\ngot:\n%s", exp, got)
	}
}

func TestSnake2CamelOneUnderscore(t *testing.T) {
	txt := `const en_urn int = 5
	const en_scim int = 44
	const en_fail int = 81
	const en_main int = 1`

	exp := `const enUrn int = 5
	const enScim int = 44
	const enFail int = 81
	const enMain int = 1`

	got := snake2camel([]byte(txt))
	if string(got) != exp {
		fmt.Printf("expected length: %d, got %d\n", len(exp), len(got))
		t.Fatalf("expected:\n%s\ngot:\n%s", exp, got)
	}
}

func TestSnake2CamelMoreUnderscores(t *testing.T) {
	txt := `if (m.p) == (m.pe) {
		goto _test_eof
	}
	switch m.cs {
	case 1:
		goto st_case_1
	}`

	exp := `if (m.p) == (m.pe) {
		goto _testEof
	}
	switch m.cs {
	case 1:
		goto stCase1
	}`

	got := snake2camel([]byte(txt))
	if string(got) != exp {
		fmt.Printf("expected length: %d, got %d\n", len(exp), len(got))
		t.Fatalf("expected:\n%s\ngot:\n%s", exp, got)
	}
}
