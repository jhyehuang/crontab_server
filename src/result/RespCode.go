package result

var (
	OK              = response(0, "ok") // 通用成功
	ParamError      = response(1000001, "参数错误")
	Err             = response(1000002, "")           // 通用错误
	SystemError     = response(1000003, "系统异常，请稍候重试") //
	TokenInvalid    = response(1000004, "token失效")
	PermissionError = response(1000005, "权限不足")

	// UserNotFound K.WithData(ret)
	//}用户模块错误吗
	UserNotFound      = response(2000001, "用户不存在")
	UserPasswordError = response(2000002, "用户名或密码错误")
	UserLoginError    = response(2000003, "登录失败")

	// 节点
	NodeNotFound    = response(3000001, "节点不存在")
	MinerNotFound   = response(3000002, "矿工不存在")
	MinerTypeErr    = response(3000003, "节点类型错误")
	OldMinerTypeErr = response(3000004, "存在算力的节点暂不支持，后续开放！")

	//钱包
	TransferError      = response(4000001, "转账失败")
	GasFeeNotFound     = response(4000002, "gas费获取失败")
	ChainStatsNotFound = response(4000003, "filscan汇总数据获取失败")

	// 服务商
	SpNotFound = response(5000001, "服务商不存在")

	// 募集计划
	PlanNotFound         = response(6000001, "募集计划不存在")
	PlanAlreadyExist     = response(6000002, "募集计划已经存在")
	PlanHasTxHash        = response(6000003, "募集计划已经上链")
	DataNotFound         = response(6000004, "数据不存在")
	AssetPackageNotFound = response(6000005, "资产包不存在")

	// 内容
	ContentNotFound = response(7000001, "内容不存在")
	UploadLogoFail  = response(7000002, "上传logo失败")
)
