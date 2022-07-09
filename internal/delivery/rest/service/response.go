package service

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func ResponseOK(c echo.Context, code int, data interface{}) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set(echo.HeaderCookie, "")

	c.Response().WriteHeader(code)
	return json.NewEncoder(c.Response()).Encode(data)
}

func ResponseError(c echo.Context, code int, err error) error {
	return c.JSON(code, map[string]string{
		"message": err.Error(),
	})
}
