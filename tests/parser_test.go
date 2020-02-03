package tests

import (
	"testing"
	"github.com/noyoshi/weatherparser/parser"
)

// Run this using `go test` inside the test directory, or specify that directory as a cli arg
func TestRunme(t *testing.T) {
	if parser.Runme() < 0 {
		t.Error("error from Runme")
	}
}
