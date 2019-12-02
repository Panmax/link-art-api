package param_bind

// https://godoc.org/github.com/go-playground/validator

type Login struct {
	Phone    string `form:"phone" binding:"required,len=11"`
	Password string `form:"password" binding:"required,min=8,max=32"`
}

type Register struct {
	Phone      string `form:"phone" binding:"required,len=11"`
	Password   string `form:"password" binding:"required,min=8,max=32"`
	Sms        string `form:"sms" binding:"required,len=6"`
	InviteCode string `form:"invite_code"`
}
