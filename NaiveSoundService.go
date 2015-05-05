package ozsndqueue

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

const playCmdWav = "aplay"
const playCmdMp3 = "mpg321"

type NaiveSoundService struct{
	PlayCmdWav string
	PlayCmdMp3 string
}

func (this NaiveSoundService) Play(fileUri string) error {
	filePath := filepath.Ext(fileUri)

	var command string
	if strings.EqualFold(filePath, ".wav") {
		if this.PlayCmdWav == "" {
			command = playCmdWav
		} else {
			command = this.PlayCmdWav
		}
	} else if strings.EqualFold(filePath, ".mp3") {
		if this.PlayCmdMp3 == "" {
			command = playCmdMp3
		} else {
			command = this.PlayCmdMp3
		}
	} else {
		return fmt.Errorf("file type not support")
	}
	return exec.Command(command, fileUri).Run()
}
