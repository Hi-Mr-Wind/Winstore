package model

import (
	"encoding/json/v2"
)

type BaseModel struct {
	// 状态码
	Code int `json:"code"`
	// 状态
	Message string `json:"message"`
	// 数据
	Data any `json:"data,omitempty"`
}

func (base *BaseModel) New() string {
	marshal, err := json.Marshal(base)
	if err != nil {
		return "error:" + err.Error()
	}
	return string(marshal)
}
