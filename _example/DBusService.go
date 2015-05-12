/**

とりあえずの動作確認用。
下記コマンドで、 dbus にメソッドが公開されているかを確認できる。

dbus-send --dest=jp.dip.oyasirazu.ozsndqueue --type=method_call --print-reply /jp/dip/oyasirazu/ozsndqueue jp.dip.oyasirazu.ozsndqueue.StartListen
dbus-send --dest=jp.dip.oyasirazu.ozsndqueue --type=method_call --print-reply /jp/dip/oyasirazu/ozsndqueue jp.dip.oyasirazu.ozsndqueue.PauseListen
dbus-send --dest=jp.dip.oyasirazu.ozsndqueue --type=method_call --print-reply /jp/dip/oyasirazu/ozsndqueue jp.dip.oyasirazu.ozsndqueue.StartPlay
dbus-send --dest=jp.dip.oyasirazu.ozsndqueue --type=method_call --print-reply /jp/dip/oyasirazu/ozsndqueue jp.dip.oyasirazu.ozsndqueue.PausePlay
dbus-send --dest=jp.dip.oyasirazu.ozsndqueue --type=method_call --print-reply /jp/dip/oyasirazu/ozsndqueue jp.dip.oyasirazu.ozsndqueue.Put string:"PATH_TO_FILE" int32:0
dbus-send --dest=jp.dip.oyasirazu.ozsndqueue --type=method_call --print-reply /jp/dip/oyasirazu/ozsndqueue jp.dip.oyasirazu.ozsndqueue.PlayNow string:"PATH_TO_FILE" */

package main

import (
	"../../ozsndqueue"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type ExampleListenerService struct {}

func (this ExampleListenerService) StartListen() {
	fmt.Println("StartListen")
}

func (this ExampleListenerService) PauseListen() {
	fmt.Println("PauseListen")
}

func (this ExampleListenerService) StartPlay() {
	fmt.Println("StartPlay")
}

func (this ExampleListenerService) PausePlay() {
	fmt.Println("PausePlay")
}

func (this ExampleListenerService) Put(fileUri string, queueNumber int32) {
	fmt.Println("Put:", fileUri, queueNumber)
}

func (this ExampleListenerService) PlayNow(fileUri string) {
	fmt.Println("PlayNow:", fileUri)
}

func main() {
	// 終了通知用 channel
	stopChan := make(chan int)

	dbusService, err := ozsndqueue.CreateDBusService(stopChan)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dbusService.DBusServiceListener = ExampleListenerService{}

	// 停止シグナル待ち受け goroutine 作成
	go captureSigint(stopChan)

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
