package define

//版本
const (
	Version  = "/v1"
	TokenKey = "WSLKL10191112Fly"
)

// 目录
const (
	CSS_CLIENT_PATH   = "/css/"
	DART_CLIENT_PATH  = "/js/"
	IMAGE_CLIENT_PATH = "/image/"
	CSS_SVR_PATH      = "./web/"
	DART_SVR_PATH     = "./web/"
	IMAGE_SVR_PATH    = "./web/"
	FILE_PATH         = "/download/"
	FILE_SVR_PATH     = "./"
	FILE_upName       = "uploadfile"
)

//Err
const (
	Err_null       = 200  //无错误
	Err_req        = 1001 //请求错误
	Err_MethodGet  = 1002 //协议错误
	Err_MethodPost = 1003 //协议错误
	Err_Ummarshal  = 1004 //解析错误
	Err_guidNoOpt  = 1005 //guid没有该类型
	Err_TypeGuid   = 1006 //guid类型错误
	Err_dataLen    = 1007 //数据长度错误
	Err_addDir     = 1008 //添加目录错误
	Err_addFile    = 1009 //添加文件错误
	Err_svrToken   = 1010 //服务器token信息错误
	Err_token      = 1011 //验证toen失败
	Err_parame     = 1012 //参数错误
	Err_post       = 1013 //post 错误
	Err_postDo     = 1014 //postdo 错误
	Err_open       = 1015 //打开文件
	Err_logined    = 1016 //已登入
)

//MsgAct
const (
	Msg_set      = 2001
	Msg_get      = 2002
	Msg_del      = 2003
	Msg_getLenth = 2004
	Msg_fileUp   = 2005
	Msg_login    = 2006
	Msg_pub      = 2007
	Msg_route    = 2008
)

//返回信息通用结构体
type S2CBody struct {
	Code    int         //代码
	CodeMsg string      //代码说明
	Data    interface{} //数据
}

type ProxInfo struct {
	Url         string
	Data        string
	ContentType string
}

type MsgInfo struct {
	Act            int
	ConversationId string
	Index          int
	Data           interface{}
}

//---act
const (
	Act_user = 3001 //users
)