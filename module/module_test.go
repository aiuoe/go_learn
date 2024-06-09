package gomodule

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "rub3n"

	want := regexp.MustCompile(`\b` + name + `\b`)
	got, err := Hello(name)

	if ok := want.MatchString(got); !ok || err != nil {
		t.Fatalf(`Hello(%s) failed. Expected string, got %T`, name, name)
	}
}
