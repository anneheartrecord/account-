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

const GrpcAddr = "127.0.0.1:9095"

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

type AddAccountReq struct {
	Mobile   string `json:"mobile,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Gender   string `json:"gender,omitempty"`
}

type UpdateAccountReq struct {
	Id       uint32 `json:"id,omitempty"`
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
	conn, err := grpc.Dial(GrpcAddr, grpc.WithInsecure())
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
		accountList[k] = RpcAccountToWebAccount(v)
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

func GetAccountByMobileHandler(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")
	conn, err := grpc.Dial(GrpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Logger.Warn("grpc_dial_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GRPCDIALFAILED, errorCode.GrpcWrong)
		return
	}
	client := pb.NewAccountServiceClient(conn)
	res, err := client.GetAccountByMobile(context.Background(), &pb.MobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		log.Logger.Warn("get_account_by_mobile_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GETACCOUNTBYMOBILEFAILER, errorCode.GetAccountByMobileErr)
		return
	}
	data := RpcAccountToWebAccount(res)
	response.ResponseSuccessful(c, data)
	return
}

func GetAccountByIDHandler(c *gin.Context) {
	IDStr := c.DefaultQuery("id", "")
	id, err := strconv.ParseInt(IDStr, 10, 32)
	if err != nil {
		log.Logger.Warn("parse_id_to_int_failed,err_" + err.Error())
	}
	conn, err := grpc.Dial(GrpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Logger.Warn("grpc_dial_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GRPCDIALFAILED, errorCode.GrpcWrong)
		return
	}
	client := pb.NewAccountServiceClient(conn)
	res, err := client.GetAccountByID(context.Background(), &pb.IDRequest{
		Id: uint32(id),
	})
	if err != nil {
		log.Logger.Warn("get_account_by_id_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GETACCOUNTBYIDFAILED, errorCode.GetAccountByIDErr)
		return
	}
	data := RpcAccountToWebAccount(res)
	response.ResponseSuccessful(c, data)
	return
}

func AddAccountHandler(c *gin.Context) {
	var req AddAccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Logger.Warn("bind_params_failed,err_" + err.Error())
	}
	conn, err := grpc.Dial(GrpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Logger.Warn("grpc_dial_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GRPCDIALFAILED, errorCode.GrpcWrong)
		return
	}
	client := pb.NewAccountServiceClient(conn)
	res, err := client.AddAccount(context.Background(), &pb.AddAccountRequest{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: req.Password,
		Gender:   req.Gender,
	})
	if err != nil {
		log.Logger.Warn("add_account_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.ADDACCOUNTFAILED, errorCode.AddAccountErr)
		return
	}
	data := RpcAccountToWebAccount(res)
	response.ResponseSuccessful(c, data)
	return
}

func UpdateAccountHandler(c *gin.Context) {
	var req UpdateAccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Logger.Warn("bind_params_failed,err_" + err.Error())
	}
	conn, err := grpc.Dial(GrpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Logger.Warn("grpc_dial_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.GRPCDIALFAILED, errorCode.GrpcWrong)
		return
	}
	client := pb.NewAccountServiceClient(conn)
	res, err := client.UpdateAccount(context.Background(), &pb.UpdateAccountRequest{
		Id:       req.Id,
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: req.Password,
		Gender:   req.Gender,
		Role:     req.Role,
	})
	if err != nil {
		log.Logger.Warn("update_account_failed,err_" + err.Error())
		response.ResponseWrong(c, errorCode.UPDATEACCOUNTFAILED, errorCode.UpdateAccountErr)
		return
	}
	data := res.Result
	response.ResponseSuccessful(c, data)
	return
}

func RpcAccountToWebAccount(res *pb.AccountRes) AccountInfo {
	var a AccountInfo
	a.Id = res.Id
	a.Mobile = res.Mobile
	a.Role = res.Role
	a.Password = res.Password
	a.Gender = res.Gender
	a.Nickname = res.Nickname
	return a
}
