package SoundManager_test_test

import (
	"../../ozsndqueue"
	"fmt"
	"testing"
)

func TestCreateSoundManager(t *testing.T) {
	soundManager := ozsndqueue.CreateSoundManager()

	value := fmt.Sprint(soundManager)
	expected := "&{{ } map[]}"
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
	expected := "&{{ } map[0:[PATH_TO_FILE1 PATH_TO_FILE1_1] 1:[PATH_TO_FILE2] 2:[PATH_TO_FILE3] -1:[PATH_TO_FILE4]]}"
	if value != expected {
		t.Fatalf("Expected %v, but %v:", expected, value)
	}
}

func TestPlay(t *testing.T) {
	soundManager := ozsndqueue.CreateSoundManager()

	soundManager.Put("PATH_TO_FILE1", 0)
	soundManager.Put("PATH_TO_FILE2", 0)
	soundManager.Put("PATH_TO_FILE3", 1)
	soundManager.Put("PATH_TO_FILE4", 2)
	soundManager.Put("PATH_TO_FILE5", 2)

	value1 := fmt.Sprint(soundManager)
	expected1 := "&{{ } map[0:[PATH_TO_FILE1 PATH_TO_FILE2] 1:[PATH_TO_FILE3] 2:[PATH_TO_FILE4 PATH_TO_FILE5]]}"
	if value1 != expected1 {
		t.Fatalf("Expected %v, but %v:", expected1, value1)
	}

	soundManager.Play()
	value2 := fmt.Sprint(soundManager)
	expected2 := "&{{ } map[0:[PATH_TO_FILE2] 1:[PATH_TO_FILE3] 2:[PATH_TO_FILE4 PATH_TO_FILE5]]}"
	if value2 != expected2 {
		t.Fatalf("Expected %v, but %v:", expected2, value2)
	}

	soundManager.Play()
	value3 := fmt.Sprint(soundManager)
	expected3 := "&{{ } map[1:[PATH_TO_FILE3] 2:[PATH_TO_FILE4 PATH_TO_FILE5]]}"
	if value3 != expected3 {
		t.Fatalf("Expected %v, but %v:", expected3, value3)
	}

	soundManager.Play()
	value4 := fmt.Sprint(soundManager)
	expected4 := "&{{ } map[2:[PATH_TO_FILE4 PATH_TO_FILE5]]}"
	if value4 != expected4 {
		t.Fatalf("Expected %v, but %v:", expected4, value4)
	}

	soundManager.Play()
	value5 := fmt.Sprint(soundManager)
	expected5 := "&{{ } map[2:[PATH_TO_FILE5]]}"
	if value5 != expected5 {
		t.Fatalf("Expected %v, but %v:", expected5, value5)
	}

	soundManager.Play()
	value6 := fmt.Sprint(soundManager)
	expected6 := "&{{ } map[]}"
	if value6 != expected6 {
		t.Fatalf("Expected %v, but %v:", expected6, value6)
	}

}
