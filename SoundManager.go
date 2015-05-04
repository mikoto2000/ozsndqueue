package ozsndqueue

import (
	"sort"
)

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

// Play, 一番優先順位の高いファイルを再生する。
func (this *SoundManager) Play() {
	if len(this.priorityQueue) == 0 {
		return
	}

	this.prioritiedDequeue()
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
