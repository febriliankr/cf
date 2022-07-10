package usecases

import "github.com/febriliankr/go-cfstore-api/internal/entities"

func (uc *KantinUsecase) CheckCanteenBalance(in entities.CheckCanteenBalanceRequest) (entities.CheckCanteenBalanceResponse, error) {
	return uc.repo.CheckCanteenBalance(in)
}

func (uc *KantinUsecase) UpdateCanteenBalance(in entities.UpdateCanteenBalanceRequest) (entities.UpdateCanteenBalanceResponse, error) {
	var res entities.UpdateCanteenBalanceResponse
	_, err := uc.repo.UpdateCanteenBalance(in)

	if err != nil {
		return res, err
	}

	balance, err := uc.repo.CheckCanteenBalance(entities.CheckCanteenBalanceRequest{})

	res.Balance = balance.Balance

	return res, nil

}
