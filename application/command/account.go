package command

// https://godoc.org/github.com/go-playground/validator
type LoginCommand struct {
	Phone    string `form:"phone" binding:"required,len=11"`
	Password string `form:"password" binding:"required,min=8,max=32"`
}

type RegisterCommand struct {
	Phone      string `form:"phone" binding:"required,len=11"`
	Password   string `form:"password" binding:"required,min=8,max=32"`
	Sms        string `form:"sms" binding:"required,len=6"`
	InviteCode string `form:"invite_code"`
}

type UpdateProfileCommand struct {
	Name      string `from:"string" binding:"required,max=16"`
	Gender    uint8  `from:"gender" binding:"gte=0,lte=2"`
	Introduce string `form:"invite_code,max=512"`
	Birth     *int64 `from:"birth"`
}
