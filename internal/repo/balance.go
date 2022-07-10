package repo

import (
	"github.com/febriliankr/go-cfstore-api/internal/entities"
)

const queryUpdateCanteenBalance = `INSERT INTO canteen_balance (student_id, amount, type) VALUES ($1, $2, $3)`

const queryCheckCanteenBalance = `SELECT SUM(amount) FROM canteen_balance WHERE type = $1`

func (r *KantinDB) UpdateCanteenBalance(in entities.UpdateCanteenBalanceRequest) (entities.UpdateCanteenBalanceResponse, error) {
	var res entities.UpdateCanteenBalanceResponse

	_, err := r.Db.Exec(queryUpdateCanteenBalance, in.StudentID, in.Amount, in.Type)

	if err != nil {

		return res, err
	}

	return res, nil
}

func (r *KantinDB) CheckCanteenBalance(in entities.CheckCanteenBalanceRequest) (entities.CheckCanteenBalanceResponse, error) {
	var res entities.CheckCanteenBalanceResponse
	var deposited int64
	err := r.Db.QueryRowx(queryCheckCanteenBalance, entities.UpdateBalance.Deposit).Scan(&deposited)

	if err != nil {
		return res, err
	}

	var withdrawn int64

	err = r.Db.QueryRowx(queryCheckCanteenBalance, entities.UpdateBalance.Withdrawal).Scan(&withdrawn)

	res.Balance = deposited - withdrawn

	return res, nil
}
