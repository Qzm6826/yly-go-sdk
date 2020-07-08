package ylyOpenApi

func (PrintService *PrintService) TextPrint(machineCode string, content string, originId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	params["origin_id"] = originId
	return APIInterface(PrintService.config, "/print/index", params)
}

func (PrintService *PrintService) PicturePrint(machineCode string, pictureUrl string, originId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["picture_url"] = pictureUrl
	params["origin_id"] = originId
	return APIInterface(PrintService.config, "/pictureprint/index", params)
}

func (PrintService *PrintService) ExpressOrderPrint(machineCode string, content string, originId string) (interface{}, error) {
	params := make(map[string]interface{})
	params["machine_code"] = machineCode
	params["content"] = content
	params["origin_id"] = originId
	return APIInterface(PrintService.config, "/expressprint/index", params)
}