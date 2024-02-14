package chatClient

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"` // 다양한 타입을 포함시킬 수 있음

}
