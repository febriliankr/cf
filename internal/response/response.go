package response

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type SuccessRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PageRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
}

type PageIn struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func SuccessPaged(message string, data []interface{}, paging PageIn) PageRes {
	response := PageRes{
		Success: true,
		Message: message,
		Data:    data,
		Page:    paging.Page,
		Limit:   paging.Limit,
		Total:   len(data),
	}
	return response
}

func Success(message string, data interface{}) SuccessRes {
	response := SuccessRes{
		Success: true,
		Message: message,
		Data:    data,
	}
	return response
}

func Err(message string) ErrorRes {
	response := ErrorRes{
		Success: false,
		Message: message,
	}
	return response
}

func SendResponse(code int, c echo.Context, data interface{}) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set(echo.HeaderCookie, "")

	c.Response().WriteHeader(code)
	return json.NewEncoder(c.Response()).Encode(data)
}
