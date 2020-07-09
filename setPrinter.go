package ylyOpenApi

//授权绑定打印机。参数（终端号machineCode，终端密钥mSign，打印机别名printName）注意：仅支持自有型应用
func (SetPrinter *SetPrinter) AddPrinter(machineCode string, mSign string, printName string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["msign"] = mSign
	params["print_name"] = printName
	return APIInterface(SetPrinter.config, "/printer/addprinter", params)
}

func (SetPrinter *SetPrinter) SetVoice(machineCode string, content string, isFile string, aId int) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	params["is_file"] = isFile
	params["aid"] = aId
	return  APIInterface(SetPrinter.config, "/printer/setvoice", params)
}

func (SetPrinter *SetPrinter) DelVoice(machineCode string, aId int) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["aid"] = aId
	return  APIInterface(SetPrinter.config, "/printer/deletevoice", params)
}

func (SetPrinter *SetPrinter) DelPrinter(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/deleteprinter", params)
}

func (SetPrinter *SetPrinter) SetPrinterMenu(machineCode string, content string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	return APIInterface(SetPrinter.config, "/printmenu/addprintmenu", params)
}

func (SetPrinter *SetPrinter) ShutdownOrRestart(machineCode string, responseType string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	return APIInterface(SetPrinter.config, "/printer/shutdownrestart", params)
}

func (SetPrinter *SetPrinter) SetSound(machineCode string, responseType string, voice string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	params["voice"] = voice
	return APIInterface(SetPrinter.config, "/printer/setsound", params)
}

func (SetPrinter *SetPrinter) GetPrinterInfo(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/printinfo", params)
}

func (SetPrinter *SetPrinter) GetPrinterVersion(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/getversion", params)
}

func (SetPrinter *SetPrinter) CancelAllPrintOrders(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/cancelall", params)
}

func (SetPrinter *SetPrinter) CancelAPrintOrder(machineCode string, orderId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["order_id"] = orderId
	return APIInterface(SetPrinter.config, "/printer/cancelone", params)
}

func (SetPrinter *SetPrinter) SetIcon(machineCode string, imgUrl string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["img_url"] = imgUrl
	return APIInterface(SetPrinter.config, "/printer/seticon", params)
}

func (SetPrinter *SetPrinter) DelIcon(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/deleteicon", params)
}

func (SetPrinter *SetPrinter) SetPrintMode(machineCode string, responseType string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	return APIInterface(SetPrinter.config, "/printer/btnprint", params)
}

func (SetPrinter *SetPrinter) SetOrderConfirmation(machineCode string, responseType string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["response_type"] = responseType
	return APIInterface(SetPrinter.config, "/printer/getorder", params)
}

func (SetPrinter *SetPrinter) SetPushUrlByClient(cmd string, url string, status string) (interface{}, error) {
	params := make(map[string]interface{})
	params["cmd"] = cmd
	params["url"] = url
	params["status"] = status
	return APIInterface(SetPrinter.config, "/oauth/setpushurl", params)
}

func (SetPrinter *SetPrinter) SetPushUrlByAuth(machineCode string, cmd string, url string, status string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["cmd"] = cmd
	params["url"] = url
	params["status"] = status
	return APIInterface(SetPrinter.config, "/oauth/setpushurl", params)
}

func (SetPrinter *SetPrinter) GetOrderStatus(machineCode string, orderId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["order_id"] = orderId
	return APIInterface(SetPrinter.config, "/printer/getorderstatus", params)
}

func (SetPrinter *SetPrinter) GetOrderList(machineCode string, pageIndex int, pageSize int) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["page_index"] = pageIndex
	params["page_size"] = pageSize
	return APIInterface(SetPrinter.config, "/printer/getorderpaginglist", params)
}

func (SetPrinter *SetPrinter) GetPrintStatus(machineCode string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	return APIInterface(SetPrinter.config, "/printer/getprintstatus", params)
}
