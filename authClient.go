package ylyOpenApi

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type OAuthClient struct {
	config Config
}

func NewAuthClient(conf Config) OAuthClient {
	auth := OAuthClient{}
	auth.SetConfig(conf)
	return auth
}

func (oauth *OAuthClient) SetConfig(conf Config) {
	oauth.config = conf
}

//获取授权url API
//redirectUri 回调地址
//state 用于保持请求和回调的状态，在回调时，会回传该参数。开发者可以用这个参数验证请求有效性，也可以记录用户请求授权页前的位置。可防止CSRF攻击
func (oauth *OAuthClient) GetAuthUrl(redirectUri string, state string) string{
	params := url.Values{}
	responseType := "code"
	params.Set("client_id", oauth.config.clientId)
	params.Set("response_type", responseType)
	params.Set("redirect_uri", redirectUri)
	params.Set("state", state)
	oauth.config.info("GetAuthUrl[req]：" + string(params.Encode()))
	host := oauth.config.GetHost()
	authUrl := host + "/oauth/authorize?" + params.Encode()
	return authUrl
}

//自有应用获取调用凭证API
func (oauth *OAuthClient) GetAccessToken() Token {
	params := url.Values{}
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	sign := oauth.config.GetSign(timestamp)
	params.Set("client_id", oauth.config.clientId)
	params.Set("grant_type", "client_credentials")
	params.Set("scope","all")
	params.Set("timestamp", timestamp)
	params.Set("sign", sign)
	params.Set("id", oauth.GetUUID4())
	oauth.config.info("GetAccessToken[req]：" + string(params.Encode()))
	host := oauth.config.GetHost()
	reqUrl := host + "/oauth/oauth"
	resp, err := http.PostForm(reqUrl, params)
	if err != nil {
		oauth.config.error(fmt.Sprintf("HTTP 请求失败: %v", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		oauth.config.error(fmt.Sprintf("读取响应体失败: %v", err))
	}
	oauth.config.info("GetAccessToken[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//开放型应用获取调用凭API
//code 授权码
func (oauth *OAuthClient) GetAccessTokenByAuthCode(code string) Token {
	params := url.Values{}
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	sign := oauth.config.GetSign(timestamp)
	params.Set("client_id", oauth.config.clientId)
	params.Set("grant_type", "authorization_code")
	params.Set("code", code)
	params.Set("scope","all")
	params.Set("timestamp", timestamp)
	params.Set("sign", sign)
	params.Set("id", oauth.GetUUID4())
	oauth.config.info("GetAccessTokenByAuthCode[req]：" + string(params.Encode()))
	host := oauth.config.GetHost()
	reqUrl := host + "/oauth/oauth"
	resp, err := http.PostForm(reqUrl, params)
	if err != nil {
		oauth.config.error(fmt.Sprintf("HTTP 请求失败: %v", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		oauth.config.error(fmt.Sprintf("读取响应体失败: %v", err))
	}
	oauth.config.info("GetAccessTokenByAuthCode[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//自有应用&开放型应用刷新调用凭证API
//需刷新令牌refresh_token
func (oauth *OAuthClient) GetAccessTokenByRefreshToken(refreshToken string) Token {
	params := url.Values{}
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	sign := oauth.config.GetSign(timestamp)
	params.Set("client_id", oauth.config.clientId)
	params.Set("grant_type", "refresh_token")
	params.Set("refresh_token", refreshToken)
	params.Set("scope","all")
	params.Set("timestamp", timestamp)
	params.Set("sign", sign)
	params.Set("id", oauth.GetUUID4())
	oauth.config.info("GetAccessTokenByRefreshToken[req]：" + string(params.Encode()))
	host := oauth.config.GetHost()
	reqUrl := host + "/oauth/oauth"
	resp, err := http.PostForm(reqUrl, params)
	if err != nil {
		oauth.config.error(fmt.Sprintf("HTTP 请求失败: %v", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		oauth.config.error(fmt.Sprintf("读取响应体失败: %v", err))
	}
	oauth.config.info("GetAccessTokenByRefreshToken[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//极速授权获取调用凭证API 注意：仅支持开放型应用
//machineCode 终端号
//secret 密钥(可能是msign或qr_key)
func (oauth *OAuthClient) GetAccessTokenBySecret(machineCode string, secret string, secretType int) Token {
	params := url.Values{}
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	sign := oauth.config.GetSign(timestamp)
	params.Set("client_id", oauth.config.clientId)
	params.Set("machine_code", machineCode)
	if secretType == 1 {
		params.Set("qr_key", secret)
	}else {
		params.Set("msign", secret)
	}
	params.Set("scope", "all")
	params.Set("timestamp", timestamp)
	params.Set("sign", sign)
	params.Set("id", oauth.GetUUID4())
	oauth.config.info("GetAccessTokenByQrKey[req]：" + string(params.Encode()))
	host := oauth.config.GetHost()
	reqUrl := host + "/oauth/scancodemodel"
	resp, err := http.PostForm(reqUrl, params)
	if err != nil {
		oauth.config.error(fmt.Sprintf("HTTP 请求失败: %v", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		oauth.config.error(fmt.Sprintf("读取响应体失败: %v", err))
	}
	oauth.config.info("GetAccessTokenByQrKey[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//获取UUID4
func (oauth *OAuthClient) GetUUID4() string {
	u4 := uuid.New()
	return u4.String()
}
