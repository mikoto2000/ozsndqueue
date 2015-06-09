package main

import (
	"../../ozsndqueue"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	soundManager := ozsndqueue.CreateSoundManager(5)
	dbusServiceListener := &ozsndqueue.DBusServiceListenerForSoundManager{soundManager}

	// 終了通知用 channel
	stopChan := make(chan int)

	dbusService, err := ozsndqueue.CreateDBusService(stopChan)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// SoundManager に通知する Listener を登録
	dbusService.DBusServiceListener = dbusServiceListener

	// 停止シグナル待ち受け goroutine 作成
	go captureSigint(stopChan)

	// SoundManager のメインループ開始
	go soundManager.StartMainLoop()

	// DBusService 開始
	dbusService.Run()
}

func captureSigint(stop chan int) {
	sigint := make(chan os.Signal)
	signal.Notify(sigint, syscall.SIGINT)

	// stop シグナル待ち受け。
	// stop シグナルが来たら停止用 channel に送信
	<-sigint
	stop <- 0
}
