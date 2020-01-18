package process2

import (
	"../../common/message"
	"../utils"
	"encoding/json"
	"fmt"
	"net"
)


type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mse *message.Message) (err error) {

	// get user input from mse
	var loginMse message.LoginMes
	err = json.Unmarshal([]byte(mse.DATA), &loginMse)
	if err != nil {
		fmt.Println("serverProcessLogin Unmarshal fail, err = ", err)
		return
	}

	// verify and put login attempt result into struct
	var loginResMes message.LoginResMes
	if loginMse.USERID == 123 && loginMse.USERPWD == "test" {
		loginResMes.CODE = 200
		loginResMes.ERROR = "ok"
	} else {
		loginResMes.CODE = 500
		loginResMes.ERROR = "fail"
	}

	// prepare response (client and server both use message.Message struct to keep consistence)
	var resMse message.Message
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("line 35, marshal err, err = ", err)
		return
	}
	resMse.TYPE = message.LoginResMesType
	resMse.DATA = string(data)

	// return to client
	data, err = json.Marshal(resMse)
	if err != nil {
		fmt.Println("line 48, marshal err, err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)

	return
}