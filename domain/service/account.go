package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"link-art-api/application/command"
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"link-art-api/domain/repository"
	"math/rand"
	"time"
)

func SendSms(phone string) int32 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rnd.Int31n(1000000)
	return code
}

func AccountRegister(phone, password string) (*model.Account, error) {
	_, err := repository.FindAccountByPhone(phone)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("手机号码已注册，可直接登录")
	}

	account := model.NewAccount(phone, password)
	if err := model.CreateOne(account); err != nil {
		return nil, err
	}

	return account, nil
}

func GetProfile(id uint) (*representation.AccountProfileRepresentation, error) {
	account, err := repository.FindAccount(id)
	if err != nil {
		return nil, err
	}

	profile := representation.NewAccountProfileRepresentation(account,
		len(ListAccountFollow(account.ID)),
		len(ListAccountFans(account.ID)),
		GetAccountPoints(id))

	return profile, nil
}

func UpdateProfile(id uint, updateCommand *command.UpdateProfileCommand) (bool, error) {
	account, err := repository.FindAccount(id)
	if err != nil {
		return false, err
	}
	account.Name = updateCommand.Name
	account.Gender = updateCommand.Gender
	account.Introduce = updateCommand.Introduce
	if updateCommand.Birth != nil {
		birth := time.Unix(*updateCommand.Birth, 0)
		account.Birth = &birth
	}
	err = model.SaveOne(account)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ListAccountFollow(id uint) []map[string]string {
	return make([]map[string]string, 0) // TODO
}

func ListAccountFans(id uint) []map[string]string {
	return make([]map[string]string, 0) // TODO
}

func GetAccountPoints(id uint) uint {
	return 0 // TODO
}

func SubmitApproval(accountId uint, submitCommand *command.SubmitApprovalCommand) error {
	_, err := repository.FindApprovalByAccount(accountId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("认证审核中，请勿重复提交")
	}

	approval := model.NewApproval(accountId, submitCommand.Type, submitCommand.CompanyName, submitCommand.Photo)
	return model.CreateOne(approval)
}

func ApprovalPass(id uint) error {
	approval, err := repository.FindApproval(id)
	if err != nil {
		return err
	}

	account, err := repository.FindAccount(approval.AccountId)
	if err != nil {
		return err
	}

	approval.Pass()
	account.BeArtist()

	tx := model.DB.Begin()
	tx.Save(approval)
	tx.Save(account)
	return tx.Commit().Error
}

func ApprovalReject(id uint) error {
	approval, err := repository.FindApproval(id)
	if err != nil {
		return err
	}

	account, err := repository.FindAccount(approval.AccountId)
	if err != nil {
		return err
	}

	approval.Reject()
	account.CancelArtist()

	tx := model.DB.Begin()
	tx.Save(approval)
	tx.Save(account)
	return tx.Commit().Error
}

func GetUser(accountId uint) (*representation.UserRepresentation, error) {
	account, err := repository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	artist := &representation.UserRepresentation{
		ID:       account.ID,
		Name:     account.Name,
		Avatar:   account.Avatar,
		IsArtist: account.Artist,
	}
	return artist, nil
}

func Follow(accountId, followerId uint) error {
	account, err := repository.FindAccount(followerId)
	if err != nil {
		return err
	}
	if !account.Artist {
		return errors.New("非艺术家，不可关注")
	}

	flows, err := repository.FindAllFollowFlow("account_id = ? AND follower_id = ?", accountId, followerId)
	if err != nil {
		return err
	}

	if len(flows) > 0 {
		return nil
	}

	flow := &model.FollowFlow{
		AccountId:  accountId,
		FollowerId: followerId,
	}
	return model.CreateOne(flow)
}

func UnFollow(accountId, followerId uint) error {
	return repository.DeleteFollowFlow(accountId, followerId)
}

func CheckFollow(accountId, followerId uint) bool {
	flows, _ := repository.FindAllFollowFlow("account_id = ? AND follower_id = ?", accountId, followerId)
	return len(flows) > 0
}

func ListFollow(accountId uint) ([]*representation.UserRepresentation, error) {
	followerIds, err := listFollowerAccountId(accountId)
	if err != nil {
		return nil, err
	}

	results := make([]*representation.UserRepresentation, 0)
	for _, followerId := range followerIds {
		user, err := GetUser(followerId)
		if err != nil {
			return nil, err
		}

		results = append(results, user)
	}

	return results, nil
}

func listFollowerAccountId(accountId uint) ([]uint, error) {
	flows, err := repository.FindAllFollowFlow("account_id = ?", accountId)
	if err != nil {
		return nil, err
	}

	ids := make([]uint, 0)
	for _, flow := range flows {
		ids = append(ids, flow.FollowerId)
	}

	return ids, nil
}

func ListFans(accountId uint) ([]*representation.UserRepresentation, error) {
	flows, err := repository.FindAllFollowFlow("follower_id = ?", accountId)
	if err != nil {
		return nil, err
	}

	results := make([]*representation.UserRepresentation, 0)
	for _, flow := range flows {
		user, err := GetUser(flow.FollowerId)
		if err != nil {
			return nil, err
		}

		results = append(results, user)
	}

	return results, nil
}

func SearchArtist(keyword string) ([]*representation.UserRepresentation, error) {
	accounts, err := repository.FindAllAccount("artist = ? AND name LIKE ?", true, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}

	results := make([]*representation.UserRepresentation, 0)
	for _, account := range accounts {
		if !account.Artist {
			continue // 过滤非艺术家
		}
		user, err := GetUser(account.ID)
		if err != nil {
			return nil, err
		}

		results = append(results, user)
	}

	return results, nil
}
