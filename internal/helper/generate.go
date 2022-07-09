package helper

import (
	"os"
	"time"

	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/golang-jwt/jwt"
)

func GenerateUserJWT(user entities.GetUserResponse) (string, error) {

	issuedAt := time.Now()
	expiredAt := time.Now().Add(time.Hour * 24 * 7)

	claim := jwt.MapClaims{
		"student_id": user.StudentID,
		"iat":        issuedAt.Unix(),
		"exp":        expiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	var hmacSecret = []byte(os.Getenv("JWT_PRIVATE_KEY"))
	tokenString, err := token.SignedString(hmacSecret)

	return tokenString, err
}
