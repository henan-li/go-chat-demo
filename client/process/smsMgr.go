package process

import (
	"../../common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMse(mes *message.Message){

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.DATA),&smsMes)
	if err != nil{
		fmt.Println("unmarshal err = ",err)
		return
	}

	info:= fmt.Sprintf("userid:\t%d say to everyone:\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
}
