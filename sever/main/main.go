package main

import (
	"fmt"
	"net"
)


func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("server main.go 121 err = ", err)
		return
	}

}

func main() {
	fmt.Println("server listen port: 8889")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.listen error=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("wait for client to connect...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.accept err=", err)
		}

		go process(conn)
	}
}
