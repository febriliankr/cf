package repo

import "github.com/febriliankr/go-cfstore-api/internal/entities"

const queryCreatePurchase = `INSERT INTO
	user_purchase (product_slug, student_id)
	VALUES ($1, $2)`

const queryHiddenProduct = `UPDATE product SET hidden = true WHERE product_slug = $1`

func (r *KantinDB) CreatePurchase(in entities.CreatePurchaseRequest) (entities.CreatePurchaseResponse, error) {
	var res entities.CreatePurchaseResponse

	_, err := r.Db.Exec(queryCreatePurchase, in.ProductSlug, in.StudentID)

	if err != nil {
		return res, err
	}

	_, err = r.Db.Exec(queryHiddenProduct, in.ProductSlug)

	if err != nil {
		return res, err
	}

	return res, nil

}
