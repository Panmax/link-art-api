package representation

import "link-art-api/domain/model"

type AccountProfileRepresentation struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Avatar    *string `json:"avatar"`
	Phone     string  `json:"phone"`
	Introduce string  `json:"introduce"`
	Gender    uint8   `json:"gender"`
	Birth     *int64  `json:"birth"`
	Follow    int     `json:"follow_count"`
	Fans      int     `json:"fans_count"`
	Points    uint    `json:"points"`
	IsArtist  bool    `json:"is_artist"`
}

func NewAccountProfileRepresentation(account *model.Account, follow int, fans int, points uint) *AccountProfileRepresentation {
	var birth *int64
	if account.Birth != nil {
		birthUnix := account.Birth.Unix()
		birth = &birthUnix
	}

	return &AccountProfileRepresentation{
		ID:        account.ID,
		Name:      account.Name,
		Avatar:    account.Avatar,
		Phone:     account.Phone,
		Introduce: account.Introduce,
		Gender:    account.Gender,
		Birth:     birth,
		Follow:    follow,
		Fans:      fans,
		Points:    points,
		IsArtist:  account.Artist,
	}
}

type ArtistRepresentation struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Avatar *string `json:"avatar"`
}
