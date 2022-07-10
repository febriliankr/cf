package usecases

import (
	"github.com/febriliankr/go-cfstore-api/internal"
	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func NewKantinUsecase(repo internal.KantinRepo) *KantinUsecase {
	return &KantinUsecase{
		repo: repo,
	}
}

type KantinUsecase struct {
	repo internal.KantinRepo
}

func (uc *KantinUsecase) CreateProduct(in entities.CreateProductRequest) (entities.CreateProductResponse, error) {
	// create slug first
	suffix := "_" + uuid.New().String()
	productSlug := slug.Make(in.Name + suffix)
	in.ProductSlug = productSlug
	return uc.repo.CreateProduct(in)
}
func (uc *KantinUsecase) GetProduct(in entities.GetProductRequest) (entities.GetProductResponse, error) {
	return uc.repo.GetProduct(in)
}
func (uc *KantinUsecase) GetProductList(in entities.GetProductListRequest) (entities.GetProductListResponse, error) {
	return uc.repo.GetProductList(in)
}
func (uc *KantinUsecase) DeleteProduct(in entities.DeleteProductRequest) (entities.DeleteProductResponse, error) {
	return uc.repo.DeleteProduct(in)
}
