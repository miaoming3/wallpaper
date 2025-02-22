package message

const (
	ACCESSERROR  uint = 500
	CLIENTERROR  uint = 400
	Unauthorized uint = 401
	//1100- 1199 验证码的错误码
	CAPTCHAERROR uint = 1100

	//管理员错误码 1300 - 1600
	ADMINORPASSWORD uint = iota + 1296
	AdminStatusNone
	AdminChangeError
	PasswordError
	NotFoundRow

	//分类信息错误
	CategoryNameErr uint = iota + 1494
	SaveCategoryErr
	NotFoundPid
	CategoryParentError
	NotFoundError
	FoundSUCCESS
	DeleteError
	NotFoundCategory
	CreatedError
	UniqueFindFiledError

	//图片错误

	NotFoundImages uint = 1600

	//upload 文件上传错误
	UploadNotSupportExt = 1800
	UploadMaxSize       = 1801
	UploadCreateDirErr  = 1802
	UploadSaveError     = 1803
)

var messageTxt = map[uint]string{
	ACCESSERROR:          "服务器内部错误",
	CLIENTERROR:          "客户端错误",
	CAPTCHAERROR:         "验证码错误",
	Unauthorized:         "用户未授权",
	ADMINORPASSWORD:      "用户名或密码错误",
	AdminStatusNone:      "超级管理员不能禁用状态",
	AdminChangeError:     "修改失败",
	CreatedError:         "创建失败",
	UniqueFindFiledError: "数据存在",

	CategoryNameErr:     "分类名已经存在",
	SaveCategoryErr:     "保存失败",
	NotFoundPid:         "访问错误",
	CategoryParentError: "父级id不能为自己",
	NotFoundError:       "查询数据不存在",
	FoundSUCCESS:        "当前分类存在子分类",
	DeleteError:         "删除失败",
	NotFoundCategory:    "获取分类失败",

	NotFoundImages: "查找图片数据不存在,修正后重试",

	UploadNotSupportExt: "文件格式错误,不支持该格式",
	UploadMaxSize:       "文件超出最大限制",
	UploadCreateDirErr:  "创建文件夹失败",
	UploadSaveError:     "保存文件失败",
}

func GetMessage(code uint) string {
	if message, exists := messageTxt[code]; exists {
		return message
	}
	return "服务器内部错误"
}
