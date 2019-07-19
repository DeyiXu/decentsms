# Usage
```bash
go get -u github.com/DeyiXu/decentsms
```
# Import
```go
import "github.com/DeyiXu/decentsms"
```
# Coding
```go
func init() {
	decentsms.AppCode = ""
}

func main() {
	phone := "13400000000"
	code := decentsms.RandomCode(6)
	param := decentsms.Parameter{
		"code": code,
	}
	if err := decentsms.SendSms(phone, TP180XXXX, param); err != nil {
		logger.Errorf("SendRegisterCode:%s", err)
	}
}
```

### 短信模板
```go

const (
	// TP180XXXX 【我惠淘】验证码:#code#，您正在注册会员，请于5分钟内填写，如非本人操作，请忽略本短信，泄露有风险。
	TP180XXXX = "TP180XXXX"
)

```