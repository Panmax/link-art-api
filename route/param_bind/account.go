package param_bind

// https://godoc.org/github.com/go-playground/validator

type Login struct {
	Phone    string `form:"phone" binding:"required,len=11"`
	Password string `form:"password" binding:"required,min=8,max=32"`
}
