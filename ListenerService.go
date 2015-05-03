package ozsndqueue

type ListenerService interface {
	StartListen()
	PauseListen()
	StartPlay()
	PausePlay()
	Put(fileUri string, queueNumber int32)
	PlayNow(fileUri string)
}
