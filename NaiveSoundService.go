package ozsndqueue

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

const playCmdWav = "aplay"
const playCmdMp3 = "mpg321"

type NaiveSoundService struct{}

func (*NaiveSoundService) Play(fileUri string) error {
	filePath := filepath.Ext(fileUri)

	if strings.EqualFold(filePath, ".wav") {
		return exec.Command(playCmdWav, fileUri).Run()
	} else if strings.EqualFold(filePath, ".mp3") {
		return exec.Command(playCmdMp3, fileUri).Run()
	} else {
		return fmt.Errorf("file type not support")
	}
}
