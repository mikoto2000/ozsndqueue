package main

import (
	"../../ozsndqueue"
	"fmt"
)

func main() {
	soundManager := ozsndqueue.CreateSoundManager()

	soundManager.Put("PATH_TO_FILE1", 0)
	soundManager.Put("PATH_TO_FILE1_1", 0)
	soundManager.Put("PATH_TO_FILE2", 1)
	soundManager.Put("PATH_TO_FILE3", 2)
	soundManager.Put("PATH_TO_FILE4", -1)

	fmt.Println(soundManager)
}
