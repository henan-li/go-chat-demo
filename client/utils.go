package main

import (
	"../common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func writePkg(conn net.Conn, data []byte) (err error) {
	// send length to client
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("err at 63, ", err)
		return
	}

	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("err at 69, ", err)
		return
	}
	return
}

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("read data from client")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//e = errors.New("reading pkg header error")
		return
	}

	// 将bytes转成长度
	// for example, the msg is: abc啊啊, it is 7bytes but the length is 5
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		err = errors.New("reading pkg body error")
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("unmarshal fail, err = ", err)
		return
	}
	return

}
