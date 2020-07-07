package ylyOpenApi

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
)

type Config struct {
	clientId      string
	clientSecret  string
	requestUrl    string
	token         Token
	machine       Machine
	loggor    	  YlySdkLogger
	bLogger       bool
}

type YlySdkLogger interface {
	Info(message string)
	Error(message string)
}

type Token struct {
	AccessToken   string
	RefreshToken  string
	Expires_in    int
	MachineCode   string
}

type Machine struct {
	MachineCode   string
	Msign         string
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

func (conf *Config) SetMachine(machine Machine) {
	conf.machine = machine
}

func (conf *Config) SetToken(token Token) {
	conf.token = token
}

func (conf *Config) SetLogger(logger YlySdkLogger) {
	conf.loggor = logger
	conf.bLogger = true
}

func (conf *Config) GetSign(timestamp string) string {
	h := md5.New()
	h.Write([]byte(conf.clientId + timestamp + conf.clientSecret))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (conf *Config) GetUUID4() string {
	u4, err := uuid.NewV4()
	if err != nil {
		conf.error(string(err.Error()))
		return err.Error()
	}
	return u4.String()
}

func (conf *Config) GetHost() string {
	if conf.requestUrl != "" {
		return conf.requestUrl
	}
	return RequestUrl()
}

func (conf *Config) diagnosis() {
	if conf.loggor != nil {
		conf.bLogger = true
	}else {
		conf.bLogger = false
	}
}

func (conf *Config) info(info string) {
	conf.diagnosis()
	if conf.bLogger {
		conf.loggor.Info(info)
	}
}

func (conf *Config) error(err string) {
	conf.diagnosis()
	if conf.bLogger {
		conf.loggor.Error(err)
	}
}

func RequestUrl() string {
	return "https://open-api.10ss.net"
}
