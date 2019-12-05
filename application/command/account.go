package command

// https://godoc.org/github.com/go-playground/validator
type LoginCommand struct {
	Phone    string `binding:"required,len=11"`
	Password string `binding:"required,min=8,max=32"`
}

type RegisterCommand struct {
	Phone      string `binding:"required,len=11"`
	Password   string `binding:"required,min=8,max=32"`
	Sms        string `binding:"required,len=6"`
	InviteCode string
}

type UpdateProfileCommand struct {
	Name      string `binding:"required,max=16"`
	Gender    uint8  `binding:"gte=0,lte=2"`
	Introduce string `binding:"max=512"`
	Birth     *int64
}

type UpdateAvatarCommand struct {
	Url string `json:"url" binding:"required,max=512"`
}

type SubmitApprovalCommand struct {
	Type        uint8   `json:"type" binding:"min=0,max=2"`
	CompanyName *string `json:"company_name"`
	Photo       string  `json:"photo" binding:"required,max=512"`
}
