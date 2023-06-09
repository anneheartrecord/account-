package errorCode

const (
	AccountNotFound     = "账户不存在"
	AddAccountFailed    = "添加用户失败"
	UpdateAccountFailed = "更新用户失败"
	CheckAccountFailed  = "用户名或密码错误"
)

const (
	HashFailed = "md5加密失败"
)

const (
	GrpcWrong         = "调用GRPC失败"
	GetAccountListErr = "获取用户列表失败"
)
