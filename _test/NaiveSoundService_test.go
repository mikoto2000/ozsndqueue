package ozsndqueue

import (
	"../../ozsndqueue"
	"testing"
)

func TestSimple(t *testing.T) {
	ss := &ozsndqueue.NaiveSoundService{}

	error := ss.Play("data/wavtest.wav")
	if error != nil {
		t.Fatalf(error.Error())
	}

	error = ss.Play("data/mp3test.mp3")
	if error != nil {
		t.Fatalf(error.Error())
	}
}
