package main

import (
	"../model"
	"fmt"
	"net"
	"time"
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

// listen, accept and assign work
func main() {
	// redis conn
	initPool("localhost:6379",16,0,300*time.Second)
	initUserDao()

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

func initUserDao(){
	model.MyUserDao = model.NewUserDao(pool)
}