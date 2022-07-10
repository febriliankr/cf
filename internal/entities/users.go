package entities

type (
	CreateUserRequest struct {
		StudentID int64  `json:"student_id,omitempty" db:"student_id" validate:"required"`
		Password  string `json:"password,omitempty" db:"password" validate:"required"`
	}

	CreateUserResponse struct {
		StudentID int64 `json:"student_id,omitempty" db:"student_id"`
	}

	GetUserRequest struct {
		StudentID int64  `json:"student_id" db:"student_id" validate:"required"`
		Password  string `json:"password" db:"password" validate:"required"`
	}

	GetUserResponse struct {
		StudentID int64  `json:"student_id,omitempty" db:"student_id"`
		Password  string `json:"-" db:"password"`
	}

	GetUserJWT struct {
		StudentID int64 `json:"student_id,omitempty" db:"student_id"`
	}
)
