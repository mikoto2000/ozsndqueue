package SoundManager_test_test

import (
	"../../ozsndqueue"
	"fmt"
	"strings"
	"testing"
)

func TestCreateSoundManager(t *testing.T) {
	soundManager := ozsndqueue.CreateSoundManager(0)

	value := fmt.Sprint(soundManager)
	expected := "&{{ } map[] true true "
	if !strings.HasPrefix(value, expected) {
		t.Fatalf("Expected %v, but %v:", expected, value)
	}
}

func TestPut(t *testing.T) {
	soundManager := ozsndqueue.CreateSoundManager(5)

	soundManager.Put("PATH_TO_FILE1", 0)
	soundManager.Put("PATH_TO_FILE1_1", 0)
	soundManager.Put("PATH_TO_FILE2", 1)
	soundManager.Put("PATH_TO_FILE3", 2)
	soundManager.Put("PATH_TO_FILE4", -1)

	value := fmt.Sprint(soundManager)
	expected := "&{{ } map[0:[PATH_TO_FILE1 PATH_TO_FILE1_1] 1:[PATH_TO_FILE2] 2:[PATH_TO_FILE3] -1:[PATH_TO_FILE4]] true true "
	if !strings.HasPrefix(value, expected) {
		t.Fatalf("Expected %v, but %v:", expected, value)
	}
}

func TestPlay(t *testing.T) {
	soundManager := ozsndqueue.CreateSoundManager(5)

	soundManager.Put("PATH_TO_FILE1", 0)
	soundManager.Put("PATH_TO_FILE2", 0)
	soundManager.Put("PATH_TO_FILE3", 1)
	soundManager.Put("PATH_TO_FILE4", 2)
	soundManager.Put("PATH_TO_FILE5", 2)

	value1 := fmt.Sprint(soundManager)
	expected1 := "&{{ } map[0:[PATH_TO_FILE1 PATH_TO_FILE2] 1:[PATH_TO_FILE3] 2:[PATH_TO_FILE4 PATH_TO_FILE5]] true true "
	if !strings.HasPrefix(value1, expected1) {
		t.Fatalf("Expected %v, but %v:", expected1, value1)
	}

	soundManager.PlayNext()
	value2 := fmt.Sprint(soundManager)
	expected2 := "&{{ } map[0:[PATH_TO_FILE2] 1:[PATH_TO_FILE3] 2:[PATH_TO_FILE4 PATH_TO_FILE5]] true true "
	if !strings.HasPrefix(value2, expected2) {
		t.Fatalf("Expected %v, but %v:", expected2, value2)
	}

	soundManager.PlayNext()
	value3 := fmt.Sprint(soundManager)
	expected3 := "&{{ } map[1:[PATH_TO_FILE3] 2:[PATH_TO_FILE4 PATH_TO_FILE5]] true true "
	if !strings.HasPrefix(value3, expected3) {
		t.Fatalf("Expected %v, but %v:", expected3, value3)
	}

	soundManager.PlayNext()
	value4 := fmt.Sprint(soundManager)
	expected4 := "&{{ } map[2:[PATH_TO_FILE4 PATH_TO_FILE5]] true true "
	if !strings.HasPrefix(value4, expected4) {
		t.Fatalf("Expected %v, but %v:", expected4, value4)
	}

	soundManager.PlayNext()
	value5 := fmt.Sprint(soundManager)
	expected5 := "&{{ } map[2:[PATH_TO_FILE5]] true true "
	if !strings.HasPrefix(value5, expected5) {
		t.Fatalf("Expected %v, but %v:", expected5, value5)
	}

	soundManager.PlayNext()
	value6 := fmt.Sprint(soundManager)
	expected6 := "&{{ } map[] true true "
	if !strings.HasPrefix(value6, expected6) {
		t.Fatalf("Expected %v, but %v:", expected6, value6)
	}

}
