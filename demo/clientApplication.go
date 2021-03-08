package main

import (
	openApi "github.com/Qzm6826/yly-go-sdk"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var conf openApi.Config

func CheckWhetherAuthorizationIsGranted (redisCli redis.Conn) bool {
	tokenJson, _ := redis.Bytes(redisCli.Do("GET", "tokenJson"))
	var token openApi.Token
	_ = json.Unmarshal(tokenJson, &token)
	if token.Error_description != "success" {
		return false
	}
	conf.SetToken(token)
	return true
}

func ConnectRedis() redis.Conn {
	address := "127.0.0.1:6379" //host + port
	password := ""
	db := 1
	conn, _ := redis.Dial("tcp", address, redis.DialPassword(password), redis.DialDatabase(db))
	return conn
}

/**
openApi.NewAuthClient下集成了5个API：
GetAuthUrl、GetAccessToken、GetAccessTokenByAuthCode、GetAccessTokenByRefreshToken、GetAccessTokenBySecret
 */
func GetAccessToken(redisCli redis.Conn) {
	oauth := openApi.NewAuthClient(conf)
	tokenData := oauth.GetAccessToken()
	tokenJson, _ := json.Marshal(tokenData)
	_, err := redis.String(redisCli.Do("SET", "tokenJson", tokenJson))
	if err != nil {
		fmt.Println("Failed to cache token data to redis", err)
	}
	conf.SetToken(tokenData)
}

/**
ApiClient.PrintService下集成了3个API：
TextPrint、PicturePrint、ExpressOrderPrint
 */
func TextPrint(client openApi.ApiClient, machineCode string, content string, originId string)  {
	res, _ := client.PrintService.TextPrint(machineCode, content, originId)
	fmt.Println(res)
}

/**
ApiClient.SetPrinter下集成了20个API：
AddPrinter、SetVoice、DelVoice、DelPrinter、SetPrinterMenu、
ShutdownOrRestart、SetSound、GetPrinterInfo、GetPrinterVersion、CancelAllPrintOrders、
CancelAPrintOrder、SetIcon、DelIcon、SetPrintMode、SetOrderConfirmation、
SetPushUrlByClient、SetPushUrlByAuth、GetOrderStatus、GetOrderList、GetPrintStatus
 */
func AddPrinter(client openApi.ApiClient, machineCode string, mSign string, printName string) {
	res, _ := client.SetPrinter.AddPrinter(machineCode, mSign, printName)
	fmt.Println(res)
}

func main() {
	cid := ""    //你的应用id
	secret := "" //你的应用密钥
	conf = openApi.NewConfig(cid, secret)
	var logger openApi.SimpleLogger
	conf.SetLogger(&logger)
	redisCli := ConnectRedis()
	if !CheckWhetherAuthorizationIsGranted(redisCli) {
		GetAccessToken(redisCli)
	}
	client := openApi.NewClient(conf)
	AddPrinter(client, "", "", "")   //未绑定打印机，需先调用此方法
	TextPrint(client, "", "", "")
}
