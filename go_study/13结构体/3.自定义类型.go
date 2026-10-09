package main

// 自定义类型，//2222模式
const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 1001
	NetworkErrCode Code = 1002
)

type Code int

func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "成功"
	case ServiceErrCode:
		return "服务错误"
	case NetworkErrCode:
		return "网络错误"
	}
	return ""
}
func (c Code) Ok() (code Code, msg string) {
	return c, c.GetMsg()
}
func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode, ServiceErrCode.GetMsg()
		//return ServiceErrCode.Ok()
	}
	if name == "2" {
		return NetworkErrCode, NetworkErrCode.GetMsg()
		//return NetworkErrCode.Ok()
	}
	return SuccessCode, SuccessCode.GetMsg()
	//return ServiceErrCode.Ok()
}

//111111模式
//用222代替这个
//const(
//	SuccessCode int =0
//	ServiceErrCode int=1001
//	NetworkErrCode int=1002
//)
//func getCodeMsg(code int)(msg string)  {
//	switch code {
//	case SuccessCode:return "成功"
//	case ServiceErrCode:return "服务错误"
//	case NetworkErrCode:return "网络错误"
//	}
//	return ""
//}
//func webServer(name string) (code int,msg string){
//	if name=="1"{
//		return ServiceErrCode,getCodeMsg(ServiceErrCode)
//	}
//	if name=="2"{
//		return NetworkErrCode,getCodeMsg(NetworkErrCode)
//	}
//	return SuccessCode,getCodeMsg(SuccessCode)
//}
