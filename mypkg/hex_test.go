package mypkg_test

import (
	"testing"
	"mypkg"
)

func TestHex_String(t *testing.T) {
	want := "a"
	got := mypkg.Hex(10).String()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}