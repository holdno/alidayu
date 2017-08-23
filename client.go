package hnAlidayu

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"time"
)

const (
	domain = "dysmsapi.aliyuncs.com"
)

func sendMessage(userInput *UserParams) {
	params := make(Params)
	params.Set("AccessKeyId", userInput.AccessKeyId)
	params.Set("Timestamp", time.Now().UTC().Format("2006-01-02T03:04:05Z")) // 格式为：yyyy-MM-dd’T’HH:mm:ss’Z’；时区为：GMT
	params.Set("SignatureMethod", "HMAC-SHA1")                               // 建议固定值：HMAC-SHA1
	params.Set("SignatureVersion", "1.0")                                    // 建议固定值：1.0
	params.Set("SignatureNonce", GetRandomString(12))                        // 用于请求的防重放攻击，每次请求唯一
	params.Set("Action", "SendSms")                                          // API的命名，固定值，如发送短信API的值为：SendSms
	params.Set("Version", "2017-05-25")                                      // API的版本，固定值，如短信API的值为：2017-05-25
	params.Set("RegionId", "cn-hangzhou")                                    // API支持的RegionID，如短信API的值为：cn-hangzhou
	params.Set("PhoneNumbers", userInput.PhoneNumbers)                       // 短信接收号码,支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式
	params.Set("SignName", userInput.SignName)                               // 短信签名
	params.Set("TemplateParam", userInput.TemplateParam)                     // 短信模板ID
	// 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,
	// 请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,
	// 否则会导致JSON在服务端解析失败
	params.Set("TemplateCode", userInput.TemplateCode)
	// 构造待签名的请求串
	value := params.SortToJoin()
	// 生成签名 签名采用HmacSHA1算法 + Base64
	sign := Sign(userInput.AppSecret, "GET&"+SpecialUrlEncode("/")+"&"+SpecialUrlEncode(value))
	// 增加签名结果到请求参数中，发送请求。签名也要做特殊URL编码
	getParams := "?Signature=" + SpecialUrlEncode(sign) + "&" + value
	fmt.Println("http://" + domain + getParams)
	resp, err := http.Get(getParams)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (this *Params) SortToJoin() string {
	var keyList []string
	for k, _ := range *this {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	sortQueryStringTmp := ""
	for _, v := range keyList {
		sortQueryStringTmp += "&" + SpecialUrlEncode(v) + "=" + SpecialUrlEncode(this.Get(v))
	}
	fmt.Println(sortQueryStringTmp)
	// 字符串转切片截取字符串开头多余的&
	// rune切片类型返回的长度为物理长度 len([]rune(string)) 获取的是肉眼可见的长度
	result := []rune(sortQueryStringTmp)
	return string(result[1:])
}

func Sign(appSecret, paramsEncodeValue string) string {
	mac := hmac.New(sha1.New, []byte(appSecret+"&"))
	mac.Write([]byte(paramsEncodeValue))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func SpecialUrlEncode(str string) string {
	return strings.Replace(strings.Replace(strings.Replace(url.QueryEscape(str), "+", "%20", -1), "*", "%2A", -1), "%7E", "~", -1)
}
