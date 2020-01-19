package process2

import (
	"../../common/message"
	"../model"
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
	//if loginMse.USERID == 123 && loginMse.USERPWD == "test" {
	//	loginResMes.CODE = 200
	//	loginResMes.ERROR = "ok"
	//} else {
	//	loginResMes.CODE = 500
	//	loginResMes.ERROR = "fail"
	//}
	user, err := model.MyUserDao.Login(loginMse.USERID, loginMse.USERPWD)
	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.CODE = 500
			loginResMes.ERROR = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.CODE = 403
			loginResMes.ERROR = err.Error()
		} else {
			loginResMes.CODE = 505
			loginResMes.ERROR = "internal error"
		}

	} else {
		loginResMes.CODE = 200
		loginResMes.ERROR = "ok"
		fmt.Println("user is ", user)
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
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}

func (this *UserProcess) ServerProcessRegister(mse *message.Message) (err error) {
	var registerMse message.RegisterMes
	err = json.Unmarshal([]byte(mse.DATA), &registerMse)
	if err != nil {
		fmt.Println("serverProcessRegister Unmarshal fail, err = ", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMse.User)
	if err != nil {

		if err == model.ERROR_USER_EXISTS {
			registerResMes.CODE = 505
			registerResMes.ERROR = model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMes.CODE = 506
			registerResMes.ERROR = "internal server error"
		}

	} else {
		registerResMes.CODE = 200
		registerResMes.ERROR = "ok"
		//fmt.Println("user is ", user)
	}

	// prepare response (client and server both use message.Message struct to keep consistence)
	var resMse message.Message
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("line 110, marshal err, err = ", err)
		return
	}
	resMse.TYPE = message.RegisterResMesType
	resMse.DATA = string(data)

	// return to client
	data, err = json.Marshal(resMse)
	if err != nil {
		fmt.Println("line 119, marshal err, err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}
