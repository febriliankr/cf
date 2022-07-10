package entities

type (
	UpdateCanteenBalanceRequest struct {
		StudentID int64  `json:"student_id,omitempty" db:"student_id"`
		Amount    int64  `json:"amount,omitempty" db:"amount"`
		Type      string `json:"type,omitempty" db:"type"`
	}

	UpdateCanteenBalanceResponse struct {
		Balance int64 `json:"balance,omitempty"`
	}
)

type (
	CheckCanteenBalanceRequest struct {
	}

	CheckCanteenBalanceResponse struct {
		Balance int64 `json:"balance,omitempty"`
	}
)
