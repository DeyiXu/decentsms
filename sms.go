package decentsms

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/nilorg/sdk/random"
)

var (
	// URL 请求URL
	URL = ""
	// AppKey App Key
	AppKey = ""
	// AppSecret App Secret
	AppSecret = ""
	// AppCode App Code
	AppCode = ""
	// Timeout ...
	Timeout    = time.Second * 3
	httpclient = &http.Client{
		Timeout: Timeout,
	}
)

// getSmsErrorMsg 获取短信错误信息
func getSmsErrorMsg(code string) string {
	switch code {
	case "00000":
		return "调用成功"
	case "10000":
		return "参数异常"
	case "10001":
		return "手机号格式不正确"
	case "10002":
		return "模板不存在"
	case "10003":
		return "模板变量不正确"
	case "10004":
		return "变量中含有敏感词"
	case "10005":
		return "变量名称不匹配"
	case "10006":
		return "短信长度过长"
	case "10007":
		return "手机号查询不到归属地"
	case "10008":
		return "产品错误"
	case "10009":
		return "价格错误"
	case "10010":
		return "重复调用"
	case "99999":
		return "系统错误"
	}
	return "自定义未知错误"
}

// SendSms 发送短信
func SendSms(phone, tplID string, param Parameter) (err error) {
	parameter := url.Values{}
	parameter.Add("mobile", phone)
	parameter.Add("param", encodeParameter(param))
	parameter.Add("tpl_id", tplID)
	var req *http.Request
	req, err = http.NewRequest("POST", URL, strings.NewReader(parameter.Encode()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Add("Authorization", "APPCODE "+AppCode)
	var response *http.Response
	response, err = httpclient.Do(req)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}

	resultJSON, err := simplejson.NewFromReader(response.Body)
	if err != nil {
		return
	}
	if retCode, ok := resultJSON.CheckGet("return_code"); ok {
		code := retCode.MustString()
		if code != "00000" {
			return fmt.Errorf("短信发送错误:%s", getSmsErrorMsg(code))
		}
	}
	return
}

// encodeParameter 编码参数
func encodeParameter(vmap Parameter) string {
	var buffer bytes.Buffer
	keys := []string{}
	for k := range vmap {
		keys = append(keys, k)
	}
	length := len(keys)
	for i := 0; i < length; i++ {
		buffer.WriteString(fmt.Sprintf("%s:%s", keys[i], vmap[keys[i]]))
		if i < length-1 {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

// RandomCode 生成验证码
func RandomCode(length int) string {
	return random.SmsCode(length)
}

// Parameter 参数
type Parameter map[string]string
