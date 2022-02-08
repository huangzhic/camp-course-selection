package util

import "camp-course-selection/common/exception"

type R struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

//通用正确处理

func Ok(data interface{}) *R {
	r := &R{
		Code:    0,
		Data:    data,
		Message: "操作成功",
	}
	return r
}

//通用错误处理

func Error(code int) *R {
	message := exception.Code2String[code]
	r := &R{
		Code:    code,
		Message: message,
	}
	return r
}
