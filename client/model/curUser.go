package model

import (
	"../../common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}