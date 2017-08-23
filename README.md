# 阿里大于 golang-SDK

> 无意间看大阿里大于被整合进了阿里云  
> 然后就去看了眼文档，惊呆了！最新的阿里大于API版本是2017-05-25  
> 旧的SDK虽然没停但我感觉也差不多了，最近又在学习golang，就用golang照着官方的示例写了一个新版的SDK

### 功能列表  
- 短信发送  

### 下载  
> go get github.com/holdno/alidayu

### 快速使用  
短信发送示例
``` golang
package main

import (
	dayu "github.com/holdno/alidayu"
	"fmt"
)

func main() {
	userInput := &dayu.UserParams{
		AccessKeyId:   "阿里云的AccessKeyId",
		AppSecret:     "阿里云的AppSecret",
		PhoneNumbers:  "接收短信的手机号码",
		SignName:      "审核通过的签名，直接写名称",
		TemplateCode:  "审核通过的模板号",
		 // 模板变量赋值，一定是json格式，注意转义
		TemplateParam: "{\"code\": \"123456\"}",
	}
	ok, msg, err := dayu.SendMessage(userInput)
	if ok {
		fmt.Println("短信发送成功")
	} else {
	    // 根据业务进行错误处理
		fmt.Println(msg, err)
	}
}

```
