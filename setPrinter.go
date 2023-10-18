package ylyOpenApi

import "strconv"

//绑定打印机API
//machineCode 终端号
//mSign       终端密钥
//printName   打印机别名
func (SetPrinter *SetPrinter) AddPrinter(machineCode string, mSign string, printName string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["msign"] = mSign
	params["print_name"] = printName
	return APIInterface(SetPrinter.config, "/printer/addprinter", params)
}

//设置内置语音API
//machineCode 终端号
//content (播报内容，音量(1~9)，声音类型(0,1,3,4)组成json字符串) 或者 (在线语音链接，语音内容请小于24kb)
//isFile "true" or "false" , 判断content是否为在线语音链接，格式MP3
//aid 定义需设置的语音编号0~9
func (SetPrinter *SetPrinter) SetVoice(machineCode string, content string, isFile string, aId int) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	params["is_file"] = isFile
	params["aid"] = strconv.Itoa(aId)
	return  APIInterface(SetPrinter.config, "/printer/setvoice", params)
}

//删除内置语音API
//machineCode 终端号
//aid 内置语音编号0~9
func (SetPrinter *SetPrinter) DelVoice(machineCode string, aId int) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["aid"] = strconv.Itoa(aId)
	return  APIInterface(SetPrinter.config, "/printer/deletevoice", params)
}

//删除已授权的打印机API
//machineCode 终端号
func (SetPrinter *SetPrinter) DelPrinter(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/deleteprinter", params)
}

//设置应用菜单API
//machineCode 终端号
//content json字符串["菜单名","菜单url"]
func (SetPrinter *SetPrinter) SetPrinterMenu(machineCode string, content string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	return APIInterface(SetPrinter.config, "/printmenu/addprintmenu", params)
}

//关机重启API
//machineCode 终端号
//responseType 重启"restart"，关闭"shutdown" 仅支持k4机型，子版本中带字母A的机器
func (SetPrinter *SetPrinter) ShutdownOrRestart(machineCode string, responseType string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	return APIInterface(SetPrinter.config, "/printer/shutdownrestart", params)
}

//声音调节API
//machineCode 终端号
//responseType 蜂鸣器"buzzer"，喇叭"horn"
//voice 音量大小0~4
func (SetPrinter *SetPrinter) SetSound(machineCode string, responseType string, voice string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	params["voice"] = voice
	return APIInterface(SetPrinter.config, "/printer/setsound", params)
}

//获取机型打印宽度API
//machineCode 终端号
func (SetPrinter *SetPrinter) GetPrinterInfo(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/printinfo", params)
}

//获取机型软硬件版本API
//machineCode 终端号
func (SetPrinter *SetPrinter) GetPrinterVersion(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/getversion", params)
}

//取消所有未打印订单API
//machineCode 终端号
func (SetPrinter *SetPrinter) CancelAllPrintOrders(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/cancelall", params)
}

//取消单条未打印订单API
//machineCode 终端号
//orderId 易联云打印id
func (SetPrinter *SetPrinter) CancelAPrintOrder(machineCode string, orderId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["order_id"] = orderId
	return APIInterface(SetPrinter.config, "/printer/cancelone", params)
}

//设置LOGO API
//machineCode 终端号
//imgUrl logo图标地址,图片宽度最大为350px,文件大小不能超过40Kb
func (SetPrinter *SetPrinter) SetIcon(machineCode string, imgUrl string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["img_url"] = imgUrl
	return APIInterface(SetPrinter.config, "/printer/seticon", params)
}

//取消LOGO API
//machineCode 终端号
func (SetPrinter *SetPrinter) DelIcon(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/deleteicon", params)
}

//设置打印方式API
//machineCode 终端号
//responseType 开启"btnopen"，关闭"btnclose"
func (SetPrinter *SetPrinter) SetPrintMode(machineCode string, responseType string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	return APIInterface(SetPrinter.config, "/printer/btnprint", params)
}

//设置订单确认API
//machineCode 终端号
//responseType 开启"open"，关闭"close"
func (SetPrinter *SetPrinter) SetOrderConfirmation(machineCode string, responseType string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	return APIInterface(SetPrinter.config, "/printer/getorder", params)
}

//设置推送url API (自有型)
//cmd 打印完成标识"oauth_finish"，接单拒单标识"oauth_getOrder"，终端状态标识"oauth_printStatus"， 按键请求标识"oauth_request"
//url 推送地址填写必须以http://或https://开头的地址。推送地址需要支持GET访问，当GET请求访问时，请直接返回{"data":"OK"}，用于推送地址的可用性测试
//status 开启"open"，关闭"close"
func (SetPrinter *SetPrinter) SetPushUrlByClient(cmd string, url string, status string) (interface{}, error) {
	params := make(map[string]interface{})
	params["cmd"] = cmd
	params["url"] = url
	params["status"] = status
	return APIInterface(SetPrinter.config, "/oauth/setpushurl", params)
}

//设置推送url API (开放型)
//machineCode 终端号
//cmd 打印完成标识"oauth_finish"，接单拒单标识"oauth_getOrder"，终端状态标识"oauth_printStatus"， 按键请求标识"oauth_request"
//url 推送地址填写必须以http://或https://开头的地址。推送地址需要支持GET访问，当GET请求访问时，请直接返回{"data":"OK"}，用于推送地址的可用性测试
//status 开启"open"，关闭"close"
func (SetPrinter *SetPrinter) SetPushUrlByAuth(machineCode string, cmd string, url string, status string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["cmd"] = cmd
	params["url"] = url
	params["status"] = status
	return APIInterface(SetPrinter.config, "/oauth/setpushurl", params)
}

//获取订单状态API
//machineCode 终端号
//orderId 易联云打印id
func (SetPrinter *SetPrinter) GetOrderStatus(machineCode string, orderId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["order_id"] = orderId
	return APIInterface(SetPrinter.config, "/printer/getorderstatus", params)
}

//获取订单列表API
//machineCode 终端号
//pageIndex 查询条件—当前页码,暂只提供前100页数据
//pageSize 查询条件—每页显示条数,每页最大条数100
func (SetPrinter *SetPrinter) GetOrderList(machineCode string, pageIndex int, pageSize int) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["page_index"] = strconv.Itoa(pageIndex)
	params["page_size"] = strconv.Itoa(pageSize)
	return APIInterface(SetPrinter.config, "/printer/getorderpaginglist", params)
}

//订单重打（单订单）API
//machineCode 终端号
//orderId 易联云打印id
func (SetPrinter *SetPrinter) Reprint(machineCode string, orderId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["order_id"] = orderId
	return APIInterface(SetPrinter.config, "/printer/reprintorder", params)
}

//获取终端状态API
//machineCode 终端号
func (SetPrinter *SetPrinter) GetPrintStatus(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/getprintstatus", params)
}

//K8关键词设置API
//machineCode 终端号
//keys 关键词-key
//type 黑白名单类型
//content 关键词-value 数据类型 json数组
func (SetPrinter *SetPrinter) SetKeywords(machineCode string, keys string, keyType string, content string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["keys"] = keys
	params["type"] = keyType
	params["content"] = content
	return APIInterface(SetPrinter.config, "/printer/setkeywords", params)
}
