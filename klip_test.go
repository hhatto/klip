package klip

import "testing"

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
