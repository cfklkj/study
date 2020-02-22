package define

type S2CBody_http struct {
	Code int         //代码
	Data interface{} //数据
}

type S2CBody struct {
	Act  int
	Data interface{}
}
type MsgInfo struct {
	Act  int
	Data interface{}
}

type LoginInfo struct {
	Name, Pwd string
}

type DownInfo struct {
	DownUrl  string
	KeepPath string
	FileType string
}
type CertInfo struct {
	Domain string
}
