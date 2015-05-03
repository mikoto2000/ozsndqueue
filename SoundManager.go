package ozsndqueue

type SoundManager struct {
	priorityQueue map[int][]string
}

func CreateSoundManager() *SoundManager {
	soundManager := &SoundManager{}
	soundManager.priorityQueue = make(map[int][]string)
	return soundManager
}

func (this *SoundManager) Put(fileUri string, queueNumber int) {
	_, exist := this.priorityQueue[queueNumber]

	// key の存在確認。
	// なければ key に対する空スライスを追加する。
	if !exist {
		this.priorityQueue[queueNumber] = []string{}
	}

	this.priorityQueue[queueNumber] = append(this.priorityQueue[queueNumber], fileUri)
}
