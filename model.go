package hnAlidayu

import (
	"strconv"
)

type UserParams struct {
	AccessKeyId     string
	AppSecret       string
	PhoneNumbers    string // 接收手机号
	SignName        string // 短信签名
	TemplateCode    string // 短信模板ID
	TemplateParam   string // 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,否则会导致JSON在服务端解析失败
	SmsUpExtendCode string // 上行短信扩展码,无特殊需要此字段的用户请忽略此字段
	OutId           string // 外部流水扩展字段
}

type Params map[string]string

func (p Params) Get(key string) string {
	v, _ := p[key]
	return v
}

func (p Params) Set(key, value string) {
	p[key] = value
}

func (p Params) SetInterface(key string, value interface{}) {
	if value == nil {
		return
	}

	switch value.(type) {
	case int8, int16, int32, int64:
		v, _ := value.(int64)
		s := strconv.FormatInt(v, 10)
		p[key] = s
		break

	case uint8, uint16, uint32, uint64:
		v, _ := value.(uint64)
		s := strconv.FormatUint(v, 10)
		p[key] = s
		break

	case float32, float64:
		v, _ := value.(float64)
		s := strconv.FormatFloat(v, 'f', 0, 64)
		p[key] = s
		break

	case bool:
		v, _ := value.(bool)
		s := strconv.FormatBool(v)
		p[key] = s
		break

	case string:
		v, _ := value.(string)
		p[key] = v
		break
	}

	return
}
