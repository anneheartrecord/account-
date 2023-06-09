package handler

import (
	"account/errorCode"
	"account/log"
	pb "account/proto"
	"account/response"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"strconv"
)

type AccountListHandlerResp struct {
	Total int32       `json:"total"`
	Data  AccountList `json:"data"`
}

type AccountList struct {
	List []AccountInfo `json:"list"`
}

type AccountInfo struct {
	Id       int32  `json:"id,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Role     uint32 `json:"role,omitempty"`
}

func AccountListHandler(c *gin.Context) {
	pageNumberStr := c.DefaultQuery("pagenumber", "1")
	pageCntStr := c.DefaultQuery("pagecnt", "10")
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		log.Logger.Warn("pageNumber_strconv_failed,err_" + err.Error())
	}
	pageCnt, err := strconv.Atoi(pageCntStr)
	if err != nil {
		log.Logger.Warn("pageCnt_strconv_failed,err_" + err.Error())
	}
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		log.Logger.Warn("grpc_dial_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GRPCDIALFAILED, errorCode.GrpcWrong)
		return
	}
	client := pb.NewAccountServiceClient(conn)
	res, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   uint32(pageNumber),
		PageSize: uint32(pageCnt),
	})
	if err != nil {
		log.Logger.Warn("get_account_list_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GETACCOUNTLISTFAILED, errorCode.GetAccountListErr)
		return
	}
	accountList := make([]AccountInfo, res.Total)
	for k, v := range res.AccountList {
		accountList[k].Id = v.Id
		accountList[k].Nickname = v.Nickname
		accountList[k].Gender = v.Gender
		accountList[k].Password = v.Password
		accountList[k].Role = v.Role
		accountList[k].Mobile = v.Mobile
	}
	result := AccountListHandlerResp{
		Total: res.Total,
		Data: AccountList{
			List: accountList,
		},
	}
	log.Logger.Info("AccountListHandler success")
	response.ResponseSuccessful(c, result)
	return
}
