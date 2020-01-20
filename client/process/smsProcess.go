package process

import (
	"../../common/message"
	"../utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) sendGroupMes(content string) (err error) {

	var mes message.Message
	mes.TYPE = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return
	}

	mes.DATA = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)

	if err != nil {
		fmt.Println("write package err = ", err)
		return
	}
	return
}
