package biz

import (
	"account/errorCode"
	"account/internal"
	"account/model"
	pb "account/proto"
	"context"
	"errors"
	"gorm.io/gorm"
)

type AccountServer struct {
}

func PageReq(pageNumber, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNumber == 0 {
			pageNumber = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNumber - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}

}

func (a *AccountServer) GetAccountList(ctx context.Context, req *pb.PagingRequest) (*pb.AccountListRes, error) {
	var accountList []model.Account
	// 指定位置开始找
	result := internal.DB.Scopes(PageReq(int(req.PageNo), int(req.PageSize))).Find(&accountList)
	if result.Error != nil {
		return nil, result.Error
	}
	accountListRes := &pb.AccountListRes{}
	accountListRes.Total = int32(result.RowsAffected)
	for _, account := range accountList {
		accountRes := ModelToPb(account)
		accountListRes.AccountList = append(accountListRes.AccountList, accountRes)
	}
	return accountListRes, nil
}

func ModelToPb(account model.Account) *pb.AccountRes {
	return &pb.AccountRes{
		Id:       int32(account.ID),
		Mobile:   account.Mobile,
		Password: account.Password,
		Nickname: account.NickName,
		Gender:   account.Gender,
		Role:     uint32(account.Role),
	}
}

func (a *AccountServer) GetAccountByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.AccountRes, error) {
	var account model.Account
	result := internal.DB.Where(&model.Account{Mobile: req.Mobile}).First(&account)
	if result.RowsAffected <= 0 {
		return nil, errors.New(errorCode.AccountNotFound)
	}
	res := ModelToPb(account)
	return res, nil
}
func (a *AccountServer) GetAccountByID(ctx context.Context, req *pb.IDRequest) (*pb.AccountRes, error) {
	var account model.Account
	result := internal.DB.First(&account, req.Id)
	if result.RowsAffected <= 0 {
		return nil, errors.New(errorCode.AccountNotFound)
	}
	res := ModelToPb(account)
	return res, nil
}
func (a *AccountServer) AddAccount(ctx context.Context, req *pb.AddAccountRequest) (*pb.AccountRes, error) {
	var account model.Account
	account.Mobile = req.Mobile
	account.Gender = req.Gender
	account.NickName = req.Nickname
	hashedPassword, err := GetMd5(req.Password)
	if err != nil {
		return nil, errors.New(errorCode.HashFailed)
	}
	account.Password = hashedPassword
	result := internal.DB.Create(&account)
	if result.RowsAffected <= 0 {
		return nil, errors.New(errorCode.AddAccountFailed)
	}
	res := ModelToPb(account)
	return res, nil
}
func (a *AccountServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountRes, error) {
	var account model.Account
	result := internal.DB.First(&account, req.Id)
	if result.RowsAffected <= 0 {
		return nil, errors.New(errorCode.AccountNotFound)
	}
	account.Mobile = req.Mobile
	account.Gender = req.Gender
	account.NickName = req.Nickname
	hashedPassword, err := GetMd5(req.Password)
	if err != nil {
		return nil, errors.New(errorCode.HashFailed)
	}
	account.Password = hashedPassword
	result = internal.DB.Save(&account)
	if result.RowsAffected <= 0 {
		return nil, errors.New(errorCode.UpdateAccountFailed)
	}
	res := &pb.UpdateAccountRes{
		Result: true,
	}
	return res, nil
}
func (a *AccountServer) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordRes, error) {
	var account model.Account
	result := internal.DB.First(&account, req.Id)
	if result.RowsAffected <= 0 {
		return nil, errors.New(errorCode.AccountNotFound)
	}
	hashed, err := GetMd5(req.Password)
	if err != nil {
		return nil, errors.New(errorCode.HashFailed)
	}
	if hashed != account.Password {
		return nil, errors.New(errorCode.CheckAccountFailed)
	}
	return &pb.CheckPasswordRes{Result: true}, nil
}
