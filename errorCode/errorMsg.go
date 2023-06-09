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
	GrpcWrong             = "调用GRPC失败"
	GetAccountListErr     = "获取用户列表失败"
	GetAccountByMobileErr = "根据手机号获取用户失败"
	GetAccountByIDErr     = "根据ID获取用户失败"
	AddAccountErr         = "添加用户失败"
	UpdateAccountErr      = "更新用户失败"
)
