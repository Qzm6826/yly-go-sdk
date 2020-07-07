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

//获取授权url
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

//自有应用获取调用凭证
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
	resp, _ := http.PostForm(reqUrl, params)
	body, _ := ioutil.ReadAll(resp.Body)
	oauth.config.info("GetAccessToken[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//开放型应用获取调用凭证需（授权码code）
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
	resp, _ := http.PostForm(reqUrl, params)
	body, _ := ioutil.ReadAll(resp.Body)
	oauth.config.info("GetAccessTokenByAuthCode[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//自有应用&开放型应用刷新调用凭证（需刷新令牌refresh_token）
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
	resp, _ := http.PostForm(reqUrl, params)
	body, _ := ioutil.ReadAll(resp.Body)
	oauth.config.info("GetAccessTokenByRefreshToken[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//开放型应用极速授权获取调用凭证（需终端号machineCode，特殊密钥qrKey）
func (oauth *OAuthClient) GetAccessTokenByQrKey(machineCode string, qrKey string) Token {
	params := url.Values{}
	t := time.Now().Unix()
	timestamp := strconv.FormatInt(t, 10)
	sign := oauth.config.GetSign(timestamp)
	params.Set("client_id", oauth.config.clientId)
	params.Set("machine_code", machineCode)
	params.Set("qr_key", qrKey)
	params.Set("scope", "all")
	params.Set("timestamp", timestamp)
	params.Set("sign", sign)
	params.Set("id", oauth.GetUUID4())
	oauth.config.info("GetAccessTokenByQrKey[req]：" + string(params.Encode()))
	host := oauth.config.GetHost()
	reqUrl := host + "/oauth/scancodemodel"
	resp, _ := http.PostForm(reqUrl, params)
	body, _ := ioutil.ReadAll(resp.Body)
	oauth.config.info("GetAccessTokenByQrKey[res]：" + string(body))
	defer resp.Body.Close()
	var token Token
	_ = json.Unmarshal(body, &token)
	return token
}

//获取UUID4
func (oauth *OAuthClient) GetUUID4() string {
	u4, err := uuid.NewV4()
	if err != nil {
		oauth.config.error(string(err.Error()))
		return err.Error()
	}
	return u4.String()
}
