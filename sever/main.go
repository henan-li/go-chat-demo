package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()

	for{
		buf := make([]byte,8096)
		fmt.Println("read data from client")
		n,e := conn.Read(buf[:4])
		if e != nil || n != 4{
			fmt.Println("reading err=",e)
			return
		}

		fmt.Println("buf read=",buf)
	}
}

func main() {
	fmt.Println("server listen port: 8889")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen error=", err)
		return
	}

	for {
		fmt.Println("wait for client to connect...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.accept err=",err)
		}

		go process(conn)
	}
}
