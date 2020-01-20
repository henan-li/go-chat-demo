package process2

import "fmt"

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

var (
	userMgr *UserMgr
)

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

func (this *UserMgr) DelOnlineUser(userid int) {
	delete(this.onlineUsers, userid)
}

func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

func (this *UserMgr) GetOnlineUserById(userid int) (up *UserProcess, err error) {
	up,ok := this.onlineUsers[userid]
	if !ok{
		err = fmt.Errorf("user %d is not here",userid)
		return
	}
	return
}
