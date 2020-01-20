package process

import (
	"../../common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId:     notifyUserStatusMes.UserId,
			UserPwd:    "",
			UserName:   "",
			UserStatus: 0,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}

func outputOnlineUser()  {
	for id,_ := range onlineUsers{
		fmt.Println("user id:\t",id)
	}
}
