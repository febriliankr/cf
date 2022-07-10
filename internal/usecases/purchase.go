package usecases

import "github.com/febriliankr/go-cfstore-api/internal/entities"

func (uc *KantinUsecase) CreatePurchase(in entities.CreatePurchaseRequest) (entities.CreatePurchaseResponse, error) {
	return uc.repo.CreatePurchase(in)
}
