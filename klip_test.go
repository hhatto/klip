package klip

import (
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	var clips []KindleClipping
	clips, err := Load("sample_clippings.txt")
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	if len(clips) <= 0 {
		t.Errorf("invalid data")
	}
}

func TestTitle(t *testing.T) {
	var clips []KindleClipping
	clips, err := Load("sample_clippings.txt")
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	for i := range clips {
		if -1 != strings.Index(clips[i].Title, "(") {
			t.Errorf("invalid data")
		}
	}
}
