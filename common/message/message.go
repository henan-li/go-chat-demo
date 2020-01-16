package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
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
	CODE  int    `json:"code"`
	ERROR string `json:"error"`
}
