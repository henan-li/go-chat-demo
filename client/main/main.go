package main

import "fmt"
import "../process"

var userId int
var userPwd string
var userName string

func main() {

	var key int
	var loop = true

	for loop {
		fmt.Println("you are in the chat room main menu now")
		fmt.Println("register")
		fmt.Println("login")
		fmt.Println("logout")
		fmt.Println("pls enter 1,2 or 3 to continue")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("login")
			loop = false
		case 2:
			fmt.Println("register")
			loop = false
		case 3:
			fmt.Println("logout")
			loop = false
		default:
			fmt.Println("something is wrong. enter number again")
		}
	}

	if key == 1 {
		fmt.Println("enter user id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("enter pwd")
		fmt.Scanf("%s\n", &userPwd)

		up := &process.UserProcess{}
		up.Login(userId, userPwd)

	} else if key == 2 {
		fmt.Println("register")
		fmt.Println("enter user id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("enter pwd")
		fmt.Scanf("%s\n", &userPwd)
		fmt.Println("enter username")
		fmt.Scanf("%s\n", &userName)

		up := &process.UserProcess{}
		up.Register(userId, userPwd,userName)
	}
}

// when you run, you need to create a go run config. in the config page. change run kind from file to dir
// and enter the dir of current main.go file. it will compile all go files under this dir

// or you use go build at current dir. it will generate exe file, then you run that exe file in terminal
// this is a easy way and later on, can upload this exe file to server and run it
