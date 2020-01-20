package process

import (
	"../../common/message"
	"../utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("ok, xxx just logged into chat room")
	fmt.Println("1. view all users")
	fmt.Println("2. send msg")
	fmt.Println("3. msg list")
	fmt.Println("4. log out")
	fmt.Println("pls enter 1,2,3 or 4 to continue")

	var key int
	fmt.Scanf("%d\n", &key)

	var content string
	smsProcess := &SmsProcess{}


	switch key {
	case 1:
		//fmt.Println("1. view all users")
		outputOnlineUser()
	case 2:
		fmt.Println("enter the content")
		fmt.Scanf("%s\n",&content)
		smsProcess.sendGroupMes(content)
	case 3:
		fmt.Println("3. msg list")
	case 4:
		fmt.Println("4. log out")
		os.Exit(0)
	default:
		fmt.Println("select one")

	}
}


func serverProcessMes(conn net.Conn){
	tf:=&utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}

	for{
		mes,err := tf.ReadPkg()
		if err != nil{
			fmt.Println("communication error, err = ",err)
			return
		}

		//fmt.Println("reading msg is ",msg)
		switch mes.TYPE {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.DATA),&notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMse(&mes)
		default:
			fmt.Println("server return unknown data type")
		}
	}
}
