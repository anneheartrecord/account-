package biz_test

import (
	"account/biz"
	"account/internal"
	pb "account/proto"
	"context"
	"testing"
)

func TestAccountServer_AddAccount(t *testing.T) {
	internal.InitDB()
	account := biz.AccountServer{}
	req := &pb.AddAccountRequest{
		Mobile:   "12345678912",
		Password: "123456",
		Gender:   "male",
		Nickname: "happyday",
	}
	res, err := account.AddAccount(context.Background(), req)
	if err != nil || res == nil {
		t.Error("add account failed,err", err)
	}
}

func TestAccountServer_UpdateAccount(t *testing.T) {
	internal.InitDB()
	account := biz.AccountServer{}
	req := &pb.UpdateAccountRequest{
		Id:       1,
		Mobile:   "15179854333",
		Nickname: "newname",
		Password: "654321",
	}
	res, err := account.UpdateAccount(context.Background(), req)
	if err != nil || res == nil {
		t.Error("update account failed, err", err)
	}
}
