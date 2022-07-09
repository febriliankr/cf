package usecases

import (
	"github.com/febriliankr/go-cfstore-api/internal"
	"github.com/febriliankr/go-cfstore-api/internal/entities"
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
func (uc *KantinUsecase) UpdateCanteenBalance(in entities.UpdateCanteenBalanceRequest) (entities.UpdateCanteenBalanceResponse, error) {
	return uc.repo.UpdateCanteenBalance(in)
}
