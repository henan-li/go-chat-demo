package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// for general info
type Message struct {
	TYPE string `json:"type"`
	DATA string `json:"data"`
}

// for login info, client->server
type LoginMes struct {
	USERID   int    `json:"userid"`
	USERPWD  string `json:"userpwd"`
	USERNAME string `json:"username"`
}

// for login info response, server->client
type LoginResMes struct {
	CODE    int    `json:"code"`
	ERROR   string `json:"error"`
	Usersid []int
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	CODE  int    `json:"code"`
	ERROR string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`
	User
}
