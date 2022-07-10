package repo

import (
	"fmt"

	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/jmoiron/sqlx"
)

func NewKantinDB(db *sqlx.DB) *KantinDB {
	return &KantinDB{
		Db: db,
	}
}

type KantinDB struct {
	Db *sqlx.DB
}

const queryGetProduct = `SELECT product_id, name, description, product_slug, price, image_url FROM product WHERE product_slug = $1 AND hidden = false`

const queryDeleteProduct = `UPDATE product SET hidden = true WHERE product_slug = $1`

const queryCreateProduct = `INSERT INTO 
	product (name, price, description, image_url, product_slug, student_id)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id, product_slug`

func (r *KantinDB) CreateProduct(in entities.CreateProductRequest) (entities.CreateProductResponse, error) {
	var res entities.CreateProductResponse

	err := r.Db.QueryRowx(queryCreateProduct, in.Name, in.Price, in.Description, in.ImageURL, in.ProductSlug, in.StudentID).StructScan(&res)

	if err != nil {

		return res, err
	}

	return res, nil
}

func (r *KantinDB) GetProduct(in entities.GetProductRequest) (entities.GetProductResponse, error) {
	var res entities.GetProductResponse

	err := r.Db.QueryRowx(queryGetProduct, in.ProductSlug).StructScan(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *KantinDB) GetProductList(in entities.GetProductListRequest) (entities.GetProductListResponse, error) {
	var res entities.GetProductListResponse

	offset := (in.Page - 1) * in.Limit

	validSortBy := map[string]bool{
		"name":       true,
		"price":      true,
		"created_at": true,
	}

	if !validSortBy[in.SortBy] {
		in.SortBy = "created_at"
	}

	validSortOrder := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	if !validSortOrder[in.SortOrder] {
		in.SortOrder = "desc"
	}

	var queryGetProductList = fmt.Sprintf(`
	SELECT product_id, name, description, product_slug, price, image_url 
	FROM product
	WHERE hidden = false
	ORDER BY %v, %v LIMIT $1 OFFSET $2  `, in.SortBy, in.SortBy)

	rows, err := r.Db.Queryx(queryGetProductList, in.Limit, offset)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var product entities.GetProductResponse
		err := rows.StructScan(&product)

		if err != nil {
			return res, err
		}

		res.Products = append(res.Products, product)
	}

	return res, nil
}

func (r *KantinDB) DeleteProduct(in entities.DeleteProductRequest) (entities.DeleteProductResponse, error) {
	var res entities.DeleteProductResponse

	_, err := r.Db.Exec(queryDeleteProduct, in.ProductSlug)

	if err != nil {

		return res, err
	}

	return res, nil
}
