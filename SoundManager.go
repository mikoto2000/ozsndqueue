package ozsndqueue

import (
	"fmt"
	"sort"
)

type SoundManager struct {
	SoundService SoundService
	priorityQueue map[int][]string
	isListen bool
	isPlay bool
	// Put された事を通知する channel
	putChan chan int
	// 終了依頼が出されたことを通知する channel
	endChan chan int
}

func CreateSoundManager(queueSize int) *SoundManager {

	soundManager := &SoundManager{}

	soundManager.SoundService = NaiveSoundService{}
	soundManager.priorityQueue = make(map[int][]string)
	soundManager.isListen = true
	soundManager.isPlay = true

	soundManager.putChan = make(chan int, queueSize)
	soundManager.endChan = make(chan int)

	return soundManager
}

func (this SoundManager) StartMainLoop() {
	this.mainLoop()
}

func (this SoundManager) mainLoop() {
	for {
		select {
		case <-this.putChan:
			this.PlayNext()
		case <-this.endChan:
			return
		}
	}
}

func (this SoundManager) Stop() {
	this.endChan <- 0
}

func (this SoundManager) StartListen() {
	this.isListen = true
}

func (this SoundManager) PauseListen() {
	this.isListen = false
}

func (this SoundManager) StartPlay() {
	this.isPlay = true
}

func (this SoundManager) PausePlay() {
	this.isPlay = false
}

// TODO: Put を無視した場合の戻り値を考える。
func (this *SoundManager) Put(fileUri string, queueNumber int) {
	// listen 中でなければ Put されたものを無視する
	if !this.isListen {
		return
	}

	_, exist := this.priorityQueue[queueNumber]

	// key の存在確認。
	// なければ key に対する空スライスを追加する。
	if !exist {
		this.priorityQueue[queueNumber] = []string{}
	}

	this.priorityQueue[queueNumber] = append(this.priorityQueue[queueNumber], fileUri)

	// Put 通知を channel に投げる。
	this.putChan <- 0
}

// PlayNext, 一番優先順位の高いファイルを再生する。
func (this *SoundManager) PlayNext() error {
	if len(this.priorityQueue) == 0 {
		return fmt.Errorf("priority queue is empty")
	}

	fileUri := this.prioritiedDequeue()

	return this.SoundService.Play(fileUri)
}

// prioritiedDequeue, キュー郡から一番優先順位の高いファイル名を取得する。
func (this *SoundManager) prioritiedDequeue() string {
	// そもそも map に要素がないのであれば空文字を返却する。
	if len(this.priorityQueue) == 0 {
		return ""
	}

	// キーを抜き出してスライスに詰める
	keys := []int{}
	for key := range this.priorityQueue {
		keys = append(keys, key)
	}

	// 優先度順にソート
	sort.Ints(keys)

	// 優先度順に探索して、最初に見つかったものを返却する。
	// 返却した要素は削除する。
	for _, value := range keys {
		fileUri, fileUris := dequeue(this.priorityQueue[value])

		// キューの更新
		if (len(fileUris) == 0) {
			// キューが空になったら map から削除する
			delete(this.priorityQueue, value)
		} else {
			// 空でなければ Dequeue 済みのスライスと入れ替える
			this.priorityQueue[value] = fileUris
		}

		// Dequeue した fuleUri を返却。
		return fileUri
	}

	// 見つからないことは無いはずだけど、
	// 見つからなければから文字を返す。
	return ""
}

// ファイル名リスト(スライス)から要素を Dequeue する。
func dequeue(fileUris []string) (string, []string) {
	// そもそも要素がないのであれば空文字を返却する。
	if len(fileUris) == 0 {
		return "", fileUris
	}

	// 先頭 1 要素とそれ以降のスライスを返却。
	// この前に長さチェックしてるからそのまま返せるはず。
	return fileUris[0], fileUris[1:len(fileUris)]
}
