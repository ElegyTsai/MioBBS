package system

type Register struct {
	Username string `json:"username"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
