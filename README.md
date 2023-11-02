# Golang SDK 接入指南

## 接入指南

  1. Go version >= 1.8
  2. 通过 go get 命令安装 SDK
  4. 使用 SDK 提供的接口进行开发调试

 ## 安装
 
 ```go
     go get github.com/Qzm6826/yly-go-sdk
 ```
 
 ## 基本用法
 
 ```go
     import openapi "github.com/Qzm6826/yly-go-sdk"
 
     // 新建一个配置实例
     conf := openApi.NewConfig(cid, secret)
     
     // 获取 token 并设置
     oauth := openApi.NewAuthClient(conf)
     tokenData := oauth.GetAccessToken()
     conf.SetToken(tokenData)
     
     // 新建一个 API 实例
     client := openApi.NewClient(conf)
     
     // 添加一个打印机。未绑定打印机，需先调用此方法
     res, _ := client.SetPrinter.AddPrinter(machineCode, mSign, printName)
     fmt.Println(res)
     
     // 调用服务 API
     res, _ := client.PrintService.TextPrint(machineCode, content, originId)
     fmt.Println(res)
 
 ```
 
 ## Token获取
 开放型应用与自有型应用的 token 获取方法略有不同。
 
 实际使用过程中，在 token 获取成功后，该 token 可以使用较长一段时间，需要缓存起来，请勿每次请求都重新获取 token。
 
 ### 开放型应用
 
 
 ```go
     import openapi "github.com/Qzm6826/yly-go-sdk"
     
     // 新建一个配置实例
     conf := openApi.NewConfig(cid, secret)
 
     // 新建 oauth 客户端实例
     oauth := openApi.NewAuthClient(conf)
     
     // 根据 OAuth 2.0 中的对应 state 和 redirectUri，获取授权 URL, 跳转访问，通过回调获取AuthCode
     authURL := oauth.GetAuthUrl(redirectUri, state)
    
 ```
 
 商家打开授权URL，同意授权后，跳转到您的回调页面，并返回code
 
 ```go
     ...
     // 通过授权得到的 code，获取token
     tokenData := oauth.GetAccessTokenByAuthCode(code)
     conf.SetToken(tokenData)
     ...
 ```
 
 自有型应用可以看看`基本用法`
 
 ## Demo使用方法
 
 该 demo 主要用来演示自有型应用的授权流程和打印
 
 1. 在开发者中心创建自有型应用
 
 2. 在 demo 的同一目录clientApplication.go文件中 配置应用信息和redis信息，否则无法运行 demo。
 
 3. 运行 demo。
 
 ## 更新日志
 #### [v2.0.3]
 * Release Date : 2023-11-01
 1. [Feature]v2新增[K8推送开关设置](https://www.kancloud.cn/ly6886/oauth-api/3208323)接口。
 2. [Feature]v2新增[K8高级设置](https://www.kancloud.cn/ly6886/oauth-api/3208324)接口。
 ### [v2.0.1]
 1. [Feature]v2新增[K8关键词设置](https://www.kancloud.cn/ly6886/oauth-api/3198288)接口。
 ### [v2.0]
 * Release Date : 2023-06-07
 1. [Feature]更新接口v2.0版本，[文档](https://www.kancloud.cn/ly6886/oauth-api/3170299)
 2. [Feature]v2新增[订单重打（单订单）](https://www.kancloud.cn/ly6886/oauth-api/3170332)接口。
 3. [Feature]v2新增[面单取消](https://www.kancloud.cn/ly6886/oauth-api/3170326)
 
