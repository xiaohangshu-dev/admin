package user

// UserDto 用户信息数据传输对象
type UserDto struct {
	UserName    string  `json:"username"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	Nikename    string  `json:"nikename"`
	AccessToken string  `json:"access_token"`
}
