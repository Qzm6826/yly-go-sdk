package ylyOpenApi

//文本打印。参数（终端号machineCode，打印内容content，商户订单号originId）
func (PrintService *PrintService) TextPrint(machineCode string, content string, originId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	params["origin_id"] = originId
	return APIInterface(PrintService.config, "/print/index", params)
}

//图片打印。参数（终端号machineCode，图片链接pictureUrl，商户订单号originId）
func (PrintService *PrintService) PicturePrint(machineCode string, pictureUrl string, originId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["picture_url"] = pictureUrl
	params["origin_id"] = originId
	return APIInterface(PrintService.config, "/pictureprint/index", params)
}

//面单打印。参数（终端号machineCode，面单数据详情请看文档content，商户订单号originId）
func (PrintService *PrintService) ExpressOrderPrint(machineCode string, content string, originId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	params["origin_id"] = originId
	return APIInterface(PrintService.config, "/expressprint/index", params)
}