package urn

import "testing"

func TestNilURNString(t *testing.T) {
	var u *URN
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panicked: %v", r)
		}
	}()
	if u.String() != "" {
		t.Fatal("want empty")
	}
}
