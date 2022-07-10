package service

import (
	"net/http"

	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/labstack/echo/v4"
)

func (s *KantinService) UpdateCanteenBalanceHandler(c echo.Context) error {
	var req entities.UpdateCanteenBalanceRequest

	user, err := getUserByBererToken(c)

	if err != nil {
		return ResponseError(c, http.StatusUnauthorized, err)
	}

	req.StudentID = user.StudentID

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.UpdateCanteenBalance(req)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return ResponseOK(c, http.StatusOK, res)

}
func (s *KantinService) CheckCanteenBalanceHandler(c echo.Context) error {
	var req entities.CheckCanteenBalanceRequest

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.CheckCanteenBalance(req)

	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return ResponseOK(c, http.StatusOK, res)

}
