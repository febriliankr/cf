package service

import (
	"net/http"

	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/labstack/echo/v4"
)

func (s *KantinService) CreatePurchaseHandler(c echo.Context) error {
	var req entities.CreatePurchaseRequest

	user, err := getUserByBererToken(c)

	if err != nil {
		return ResponseError(c, http.StatusUnauthorized, err)
	}

	req.StudentID = user.StudentID

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.CreatePurchase(req)

	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}
