package ozsndqueue

import (
	"fmt"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
)

//// 基本情報
const SERVICE_NAME = "jp.dip.oyasirazu.ozsndqueue"
const SERVICE_PATH dbus.ObjectPath = "/jp/dip/oyasirazu/ozsndqueue"

// 公開メソッド
const SERVICE_METHOD_StartListen = "StartListen"
const SERVICE_METHOD_PauseListen = "PauseListen"
const SERVICE_METHOD_StartPlay = "StartPlay"
const SERVICE_METHOD_PausePlay = "PausePlay"
const SERVICE_METHOD_Put = "Put"
const SERVICE_METHOD_PlayNow = "PlayNow"

// intro 情報
const SERVICE_INTRO = `<node>
		<interface name="` + SERVICE_NAME + `">
			<method name="` + SERVICE_METHOD_StartListen + `" />
			<method name="` + SERVICE_METHOD_PauseListen + `" />
			<method name="` + SERVICE_METHOD_StartPlay + `" />
			<method name="` + SERVICE_METHOD_PausePlay + `" />
			<method name="` + SERVICE_METHOD_Put + `">
				<arg direction="in" type="s"/>
				<arg direction="in" type="i"/>
			</method>
			<method name="` + SERVICE_METHOD_PlayNow + `">
				<arg direction="in" type="s"/>
			</method>
		</interface>` + introspect.IntrospectDataString + `</node> `

type DBusService struct {
	conn     *dbus.Conn
	stopChan chan int
}

// CreateDBusService, DBus サービスを作成する。
func CreateDBusService(stopChan chan int) (*DBusService, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	reply, err := conn.RequestName(SERVICE_NAME, dbus.NameFlagDoNotQueue)
	if err != nil {
		return nil, err
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		return nil, err
	}

	// DBusService 作成
	service := &DBusService{conn, stopChan}

	// 各種サービスのエクスポート
	conn.Export(service, SERVICE_PATH, SERVICE_NAME)
	conn.Export(introspect.Introspectable(SERVICE_INTRO),
		SERVICE_PATH,
		"org.freedesktop.DBus.Introspectable")

	return service, nil
}

func (this *DBusService) StartListen() *dbus.Error {
	fmt.Println("StartListen")
	return nil
}

func (this *DBusService) PauseListen() *dbus.Error {
	fmt.Println("PauseListen")
	return nil
}

func (this *DBusService) StartPlay() *dbus.Error {
	fmt.Println("StartPlay")
	return nil
}

func (this *DBusService) PausePlay() *dbus.Error {
	fmt.Println("PausePlay")
	return nil
}

func (this *DBusService) Put(fileUri string, queueNumber int32) *dbus.Error {
	fmt.Println("Put:", fileUri, queueNumber)
	return nil
}

func (this *DBusService) PlayNow(fileUri string) *dbus.Error {
	fmt.Println("PlayNow:", fileUri)
	return nil
}

func (this *DBusService) stop() {
	fmt.Println("Close dbus connection.")
	// conn のクローズ
	if this.conn != nil {
		this.conn.Close()
	}
	fmt.Println("Stop DBusService.")
}

// 終了通知が来るまでメインループ実行し続ける。
// 終了通知が来たら、 Conn をクローズして終了。
func (this *DBusService) Run() {
	fmt.Println("Start DBusService.")

	defer this.stop()

	select {
	case <-this.stopChan:
		fmt.Println("Catch stop cannel.")
		return
	}
}
