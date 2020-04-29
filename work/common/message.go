package common

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-edu/work/serializer"
)

func ValidateResponse(err error) *serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			fmt.Println(e)
			return &serializer.Response{
				Code: 40001,
				Msg:    "parameter is invalid",
				//Msg:    fmt.Sprintf("%s%s", e.Field, e.Tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return &serializer.Response{
			Code: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return &serializer.Response{
		Code: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
