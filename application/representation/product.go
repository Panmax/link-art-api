package representation

type ProductRepresentation struct {
	Name        string
	Type        uint
	Self        bool
	Price       uint
	Stock       int
	Size        string
	Year        string
	Material    string
	MainPic     string
	DetailPics  []string
	Description string
}
