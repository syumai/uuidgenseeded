package uuidseeded

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	seedReader := strings.NewReader("example")
	want := "475a87df-68e7-4c04-a4f0-e11504aec553"
	got, err := New(seedReader)
	if err != nil {
		t.Errorf("New() error = %v, wantErr %v", err, false)
		return
	}
	if got != want {
		t.Errorf("New() got = %v, want %v", got, want)
	}
}
