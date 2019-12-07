package command

type CreateProductCommand struct {
	Name        string
	Type        uint
	Self        bool
	Price       uint
	Stock       int
	Size        string
	Year        string
	Material    string
	MainPic     string
	DetailsPics []string
	Description string
}
