package service

import (
	"net/http"

	"github.com/febriliankr/go-cfstore-api/internal"
	"github.com/febriliankr/go-cfstore-api/internal/app"
	"github.com/febriliankr/go-cfstore-api/internal/app/config"
	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/labstack/echo/v4"
)

type KantinService struct {
	uc     internal.KantinUC
	config config.Config
}

func NewKantinService(app *app.SeminarApp) *KantinService {
	return &KantinService{
		uc:     app.Usecases.KantinUC,
		config: app.Cfg,
	}
}

func (s *KantinService) CreateProductHandler(c echo.Context) error {
	var req entities.CreateProductRequest

	user, err := getUserByBererToken(c)

	if err != nil {
		return ResponseError(c, http.StatusUnauthorized, err)
	}

	req.StudentID = user.StudentID
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

	sortBy, order := getSortByAndOrder(c)

	req.Page = page
	req.Limit = limit
	req.SortBy = sortBy
	req.SortOrder = order

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
