package process2

import (
	"../../common/message"
	"../utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	
}

func (this *SmsProcess) SendGroupMes(mes *message.Message)  {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.DATA),&smsMes)
	if err != nil{
		fmt.Println("server unmarshal err = ",err)
		return
	}

	data,err:= json.Marshal(mes)
	if err != nil{
		fmt.Println("marshal err = ",err)
		return
	}


	for id,up := range userMgr.onlineUsers{

		if id==smsMes.UserId{
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}

func (this *SmsProcess)SendMesToEachOnlineUser(data []byte, conn net.Conn)  {
	
	tf:=&utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}

	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("server response err = ",err)
	}
}