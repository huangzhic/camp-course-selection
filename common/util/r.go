package util


type R struct {
	Code int       `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

//通用正确处理

func Ok(data interface{}) *R {
	r := &R{
		Code: 0,
		Data: data,
	}
	return r
}

//通用错误处理

func Error(code int) *R {
	r := &R{
		Code: code,
	}
	return r
}
