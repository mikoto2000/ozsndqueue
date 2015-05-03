package SoundManager_test_test

import (
	"../../ozsndqueue"
	"fmt"
	"testing"
)

func TestCreateSoundManager(t *testing.T) {
	soundManager := ozsndqueue.CreateSoundManager()

	value := fmt.Sprint(soundManager)
	expected := "&{map[]}"
	if value != expected {
		t.Fatalf("Expected %v, but %v:", expected, value)
	}
}

func TestPut(t *testing.T) {
	soundManager := ozsndqueue.CreateSoundManager()

	soundManager.Put("PATH_TO_FILE1", 0)
	soundManager.Put("PATH_TO_FILE1_1", 0)
	soundManager.Put("PATH_TO_FILE2", 1)
	soundManager.Put("PATH_TO_FILE3", 2)
	soundManager.Put("PATH_TO_FILE4", -1)

	value := fmt.Sprint(soundManager)
	expected := "&{map[0:[PATH_TO_FILE1 PATH_TO_FILE1_1] 1:[PATH_TO_FILE2] 2:[PATH_TO_FILE3] -1:[PATH_TO_FILE4]]}"
	if value != expected {
		t.Fatalf("Expected %v, but %v:", expected, value)
	}
}
