package service

import (
	"net/http"
	"time"

	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/febriliankr/go-cfstore-api/internal/helper"
	"github.com/labstack/echo/v4"
)

func (s *KantinService) CreateUserHandler(c echo.Context) error {
	var req entities.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.CreateUser(req)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (s *KantinService) LoginUserHandler(c echo.Context) error {
	var req entities.GetUserRequest

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.GetUser(req)

	if err != nil {
		return ResponseError(c, http.StatusUnauthorized, err)
	}

	token, err := helper.GenerateUserJWT(res)

	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	cookie := generateCookie(s.config.Server.UserTokenName, token)

	c.SetCookie(cookie)

	return ResponseOK(c, http.StatusOK, token)
}

func generateCookie(tokenKey string, tokenValue string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = tokenKey
	cookie.Value = tokenValue
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(3 * 24 * time.Hour)
	return cookie
}
