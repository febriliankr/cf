package service

import (
	"net/http"

	"github.com/febriliankr/go-cfstore-api/internal"
	"github.com/febriliankr/go-cfstore-api/internal/app"
	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/labstack/echo/v4"
)

type KantinService struct {
	uc internal.KantinUC
}

func NewKantinService(app *app.SeminarApp) *KantinService {
	return &KantinService{
		uc: app.Usecases.KantinUC,
	}
}

func (s *KantinService) CreateProductHandler(c echo.Context) error {
	var req entities.CreateProductRequest

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.CreateProduct(req)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (s *KantinService) GetProductHandler(c echo.Context) error {
	var req entities.GetProductRequest

	product_slug := c.Param("product_slug")

	req.ProductSlug = product_slug

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.GetProduct(req)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return ResponseOK(c, http.StatusOK, res)

}
func (s *KantinService) GetProductListHandler(c echo.Context) error {
	var req entities.GetProductListRequest

	page, limit := getPagination(c)

	sortBy, orderBy := getSortBy(c)

	req.Page = page
	req.Limit = limit
	req.SortBy = sortBy
	req.SortOrder = orderBy

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.GetProductList(req)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return ResponseOK(c, http.StatusOK, res)
}

func (s *KantinService) DeleteProductHandler(c echo.Context) error {
	var req entities.DeleteProductRequest
	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.DeleteProduct(req)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return ResponseOK(c, http.StatusOK, res)

}
func (s *KantinService) UpdateCanteenBalanceHandler(c echo.Context) error {
	var req entities.UpdateCanteenBalanceRequest

	if err := c.Bind(&req); err != nil {
		return ResponseError(c, http.StatusBadRequest, err)
	}

	res, err := s.uc.UpdateCanteenBalance(req)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, err)
	}

	return ResponseOK(c, http.StatusOK, res)

}
