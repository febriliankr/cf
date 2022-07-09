package entities

type (
	CreateProductRequest struct {
		Name        string `json:"name,omitempty" db:"name" validate:"required"`
		Description string `json:"description,omitempty" db:"description" validate:"required"`
		ProductSlug string `json:"product_slug,omitempty" db:"product_slug" validate:"required"`
		Price       int64  `json:"price,omitempty" db:"price" validate:"required"`
		ImageURL    string `json:"image_url,omitempty" db:"image_url" validate:"required"`
		StudentID   int64  `json:"student_id,omitempty" db:"student_id" validate:"required"`
		Hidden      bool   `json:"hidden,omitempty" db:"hidden"`
	}
	
	CreateProductResponse struct {
		ProductID string `json:"product_id,omitempty" db:"product_id"`
	}
)

type (
	GetProductRequest struct {
		ProductSlug string `json:"product_slug,omitempty" db:"product_slug"`
	}

	GetProductResponse struct {
		ProductID   string `json:"product_id,omitempty" db:"product_id"`
		Name        string `json:"name,omitempty" db:"name"`
		Description string `json:"description,omitempty" db:"description"`
		ProductSlug string `json:"product_slug,omitempty" db:"product_slug"`
		Price       int64  `json:"price,omitempty" db:"price"`
		ImageURL    string `json:"image_url,omitempty" db:"image_url"`
	}
)

type (
	GetProductListRequest struct {
		Page      int64  `json:"page,omitempty" db:"page"`
		Limit     int64  `json:"limit,omitempty" db:"limit"`
		SortBy    string `json:"sort_by,omitempty" db:"sort_by"`
		SortOrder string `json:"sort_order,omitempty" db:"sort_order"`
	}
	GetProductListResponse struct {
		Products []GetProductResponse `json:"products,omitempty" db:"products"`
		Total    int64                `json:"total,omitempty" db:"total"`
	}
)

type (
	DeleteProductRequest struct {
		ProductSlug string `json:"product_slug,omitempty" db:"product_slug"`
	}

	DeleteProductResponse struct {
	}
)

type (
	UpdateCanteenBalanceRequest struct {
		UID    string `json:"uid,omitempty" db:"uid"`
		Amount int64  `json:"amount,omitempty" db:"amount"`
		Type   string `json:"type,omitempty" db:"type"`
	}

	UpdateCanteenBalanceResponse struct {
	}
)
