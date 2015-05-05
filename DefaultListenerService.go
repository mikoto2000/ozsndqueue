package ozsndqueue

import (
	"fmt"
)

type DefaultListenerService struct {
	SoundManager *SoundManager
}

func CreateDefaultListenerService(queueSize int) *DefaultListenerService {
	defaultListenerService := &DefaultListenerService{}

	defaultListenerService.SoundManager = CreateSoundManager(queueSize)

	return defaultListenerService
}

func (this DefaultListenerService) StartListen() {
	fmt.Println("StartListen")
}

func (this DefaultListenerService) PauseListen() {
	fmt.Println("PauseListen")
}

func (this DefaultListenerService) StartPlay() {
	fmt.Println("StartPlay")
}

func (this DefaultListenerService) PausePlay() {
	fmt.Println("PausePlay")
}

func (this DefaultListenerService) Put(fileUri string, queueNumber int32) {
	fmt.Println("Put:", fileUri, queueNumber)
}

func (this DefaultListenerService) PlayNow(fileUri string) {
	fmt.Println("PlayNow:", fileUri)
}

