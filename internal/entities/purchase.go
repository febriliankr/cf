package entities

type (
	CreatePurchaseRequest struct {
		ProductSlug string `json:"product_slug,omitempty" db:"product_slug" validate:"required"`
		StudentID   int64  `json:"student_id,omitempty" db:"student_id" validate:"required"`
	}

	CreatePurchaseResponse struct {
	}
)
