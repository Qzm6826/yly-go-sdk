package ylyOpenApi

import (
	"crypto/md5"
	"encoding/hex"
)

type Config struct {
	clientId      string
	clientSecret  string
	requestUrl    string
	token         Token
	logger    	  YlySdkLogger
	bLogger       bool
}

type YlySdkLogger interface {
	Info(message string)
	Error(message string)
}

type Token struct {
	Error             int
	Error_description string
	Body              TokenBody
}

type TokenBody struct {
	Access_token   string
	Refresh_token  string
	Expires_in     int
	Machine_code   string
}

func NewConfig(cid string, secret string) Config {
	conf := Config{}
	conf.SetClientId(cid)
	conf.SetClientSecret(secret)
	return conf
}

func (conf *Config) SetClientId(cid string) {
	conf.clientId = cid
}

func (conf *Config) SetClientSecret(secret string) {
	conf.clientSecret = secret
}

func (conf *Config) SetRequestUrl(reqUrl string) {
	conf.requestUrl = reqUrl
}

func (conf *Config) SetToken(token Token) {
	conf.token = token
}

func (conf *Config) SetLogger(logger YlySdkLogger) {
	conf.logger = logger
	conf.bLogger = true
}

func (conf *Config) GetSign(timestamp string) string {
	h := md5.New()
	h.Write([]byte(conf.clientId + timestamp + conf.clientSecret))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (conf *Config) GetHost() string {
	if conf.requestUrl != "" {
		return conf.requestUrl
	}
	return RequestUrl()
}

func (conf *Config) diagnosis() {
	if conf.logger != nil {
		conf.bLogger = true
	}else {
		conf.bLogger = false
	}
}

func (conf *Config) info(info string) {
	conf.diagnosis()
	if conf.bLogger {
		conf.logger.Info(info)
	}
}

func (conf *Config) error(err string) {
	conf.diagnosis()
	if conf.bLogger {
		conf.logger.Error(err)
	}
}

func RequestUrl() string {
	return "https://open-api.10ss.net"
}
