package process

import (
	"../../common/message"
	"../utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//{
	//	type:xxx,
	//	data:{
	//		id:xxx,
	//		pwd:xxx,
	//		name:xxx,
	//	},
	//}

	// save into struct
	var loginMes message.LoginMes
	loginMes.USERID = userId
	loginMes.USERPWD = userPwd
	loginMes.USERNAME = ""

	// prepare, server need {type:xxxx,data:{xxx}}, which is message.Message struct
	var mes message.Message
	data, e := json.Marshal(loginMes)
	if e != nil {
		fmt.Println("json.marshal err=", e)
	}

	mes.TYPE = message.LoginMesType
	mes.DATA = string(data)

	data, e = json.Marshal(mes)
	if e != nil {
		fmt.Println("json.marshel msg err=", e)
		return
	}

	// send
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// connect tcp
	conn, e := net.Dial("tcp", "localhost:8889")
	if e != nil {
		fmt.Println("net.dial connect error=", e)
		return
	}
	defer conn.Close()

	// send data length first
	n, e := conn.Write(buf[:4])
	if n != 4 || e != nil {
		fmt.Println("conn.write(bytes) fail, err=", e)
		return
	}

	//fmt.Printf("length has sent to server,length=%d, data = %s",len(data),string(data))
	// send login data
	_, e = conn.Write(data)
	if e != nil {
		fmt.Println("conn.write(data) fail, err=", e)
		return
	}

	// receive res from server (struct is message.Message)
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	mse, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("login.go line 71, err=", e)
		return
	}

	// write response msg into login response struct
	var loginMesRes message.LoginResMes
	err = json.Unmarshal([]byte(mse.DATA), &loginMesRes)
	if err != nil {
		fmt.Println("login.go line 79, err=", e)
		return
	}

	// notify result
	if loginMesRes.CODE == 200 {
		go serverProcessMes(conn)
		//fmt.Println(loginMesRes.ERROR)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginMesRes.ERROR)
	}
	return
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	conn, e := net.Dial("tcp", "localhost:8889")
	if e != nil {
		fmt.Println("net.dial connect error=", e)
		return
	}
	defer conn.Close()

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	var mes message.Message
	data, e := json.Marshal(registerMes)
	if e != nil {
		fmt.Println("json.marshal err=", e)
	}

	mes.TYPE = message.RegisterMesType
	mes.DATA = string(data)

	data, e = json.Marshal(mes)
	if e != nil {
		fmt.Println("json.marshel msg err=", e)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}

	e = tf.WritePkg(data)
	if e != nil {
		fmt.Println("conn.write(data) fail, err=", e)
		return
	}

	mse, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("login.go line 71, err=", e)
		return
	}

	var registerMesRes message.RegisterResMes
	err = json.Unmarshal([]byte(mse.DATA), &registerMesRes)
	if err != nil {
		fmt.Println("login.go line 79, err=", e)
		return
	}

	if registerMesRes.CODE == 200 {
		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(registerMesRes.ERROR)
	}
	return
}
