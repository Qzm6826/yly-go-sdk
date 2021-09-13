package ylyOpenApi

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type apiResponse struct {
	Error string
	Error_description string
	Body  interface{}
}

type APIError struct {
	Code    string
	Message string
}

func (e APIError) Error() string {
	return e.Code + " " + e.Message
}

func APIInterface(config *Config, strAction string, params map[string]interface{}) (interface{}, error) {
	tokenBody := config.token.Body
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	sign := config.GetSign(timestamp)
	reqParams := url.Values{}
	params["client_id"] = config.clientId
	params["access_token"] = tokenBody.Access_token
	params["timestamp"] = timestamp
	params["sign"] = sign
	params["id"] = GetUUID4()
	for k, v := range params {
		v := v.(string)
		reqParams.Set(k, v)
	}
	config.info(strAction + "[req]：" + reqParams.Encode())
	host := config.GetHost()
	reqUrl := host + strAction
	resp, _ := http.PostForm(reqUrl, reqParams)
	body, _ := ioutil.ReadAll(resp.Body)
	config.info(strAction + "[res]：" + string(body))
	defer resp.Body.Close()
	var apiResp apiResponse
	json.Unmarshal([]byte(body), &apiResp)
	if apiResp.Error_description != "success" {
		errorInfo := make(map[string]string)
		errorInfo["code"] = apiResp.Error
		errorInfo["message"] = apiResp.Error_description
		tmp, _ := json.Marshal(errorInfo)
		config.error(string(tmp))
		var apiErr APIError
		json.Unmarshal(tmp, &apiErr)
		return nil, apiErr
	}
	return apiResp.Body , nil
}

//获取UUID4
func GetUUID4() string {
	u4, err := uuid.NewV4()
	if err != nil {
		return err.Error()
	}
	return u4.String()
}
