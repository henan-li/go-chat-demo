package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	TYPE string `json:"type"`
	DATA string `json:"data"`
}

type LoginMes struct {
	USERID   int    `json:"userid"`
	USERPWD  string `json:"userpwd"`
	USERNAME string `json:"username"`
}

type LoginResMes struct {
	CODE  int    `json:"code"`
	ERROR string `json:"error"`
}
