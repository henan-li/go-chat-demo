package main

import (
	"../../common/message"
	"../process"
	"../utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// call different function based on msg type
func (this *Processor) serverProcessMse(mse *message.Message) (err error) {

	switch mse.TYPE {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mse)
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mse)
	default:
		fmt.Println("msg type is wrong, can not handle this request")
	}
	return
}

// read transferred data and decode it
func (this *Processor) process2() (err error) {
	for {

		tf := &utils.Transfer{
			Conn: this.Conn,
			Buf:  [8096]byte{},
		}
		mes, err := tf.ReadPkg() // this mes structure is message.Message
		if err != nil {
			if err == io.EOF {
				fmt.Println("client terminate ")
				return err
			} else {
				fmt.Println("reading data err = ", err)
				return err
			}

		}

		fmt.Println("msg is ", mes)
		err = this.serverProcessMse(&mes)
		if err != nil {
			return err
		}
	}

}
