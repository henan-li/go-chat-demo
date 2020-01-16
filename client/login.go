package main

import (
	"../common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error) {
	//{
	//	type:xxx,
	//	data:{
	//		id:xxx,
	//		pwd:xxx,
	//		name:xxx,
	//	},
	//}

	// save use inputs
	var loginMes message.LoginMes
	loginMes.USERID = userId
	loginMes.USERPWD = userPwd
	loginMes.USERNAME = ""
	data, e := json.Marshal(loginMes)
	if e != nil{
		fmt.Println("json.marshal err=",e)
	}

	// prepare user login data for server
	var mes message.Message
	mes.TYPE = message.LoginMesType
	mes.DATA = string(data)
	data,e=json.Marshal(mes)
	if e != nil{
		fmt.Println("json.marshel msg err=",e)
		return
	}

	// now, need to send data to server
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	// send data length first
	conn, e := net.Dial("tcp", "localhost:8889")
	if e != nil {
		fmt.Println("net.dial connect error=", e)
		return
	}
	defer conn.Close()

	n,e:=conn.Write(buf[:4])
	if n!=4 || e != nil{
		fmt.Println("conn.write(bytes) fail, err=",e)
		return
	}

	fmt.Printf("length has sent to server,length=%d, data = %s",len(data),string(data))
	return
}
