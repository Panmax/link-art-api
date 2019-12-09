package command

type CreateProductCommand struct {
	Name        string `binding:"required"`
	Type        uint   `binding:"required"`
	Self        bool   `binding:"required"`
	Price       uint   `binding:"required"`
	Stock       int    `binding:"required"`
	Length      *uint
	Width       *uint
	Year        *string
	Material    string   `binding:"required"`
	MainPic     string   `json:"main_pic" binding:"required"`
	DetailPics  []string `json:"detail_pics" binding:"required"`
	Description string
}
