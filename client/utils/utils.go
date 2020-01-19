package utils

import (
	"../../common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// send length to client
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("err at 63, ", err)
		return
	}

	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("err at 69, ", err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)
	fmt.Println("read data from client")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//e = errors.New("reading pkg header error")
		return
	}

	// 将bytes转成长度
	// for example, the msg is: abc啊啊, it is 7bytes but the length is 5
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		err = errors.New("reading pkg body error")
		return
	}

	// new(message.Message), pointer as unmarshal second param
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("unmarshal fail, err = ", err)
		return
	}
	return

}
