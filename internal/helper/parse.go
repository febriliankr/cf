package helper

import (
	"errors"
	"fmt"
	"os"

	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt"
)

func ParseUserCookie(c echo.Context, db *sqlx.DB) (entities.GetUserResponse, error) {
	var user entities.GetUserResponse
	cookies := c.Cookies()

	if len(cookies) == 0 {
		return user, errors.New("no cookie found")
	}
	cookie, err := c.Cookie("compfest_user")
	if err != nil {
		return user, err
	}
	cookieVal := cookie.Value
	user, err = ParseUserByJWT(cookieVal)
	if err != nil {
		return user, err
	}
	return user, nil
}

func ParseUserByJWT(tokenStr string) (entities.GetUserResponse, error) {
	var res entities.GetUserResponse

	jwtToken, err := jwt.Parse(tokenStr, keyFunction)

	if err != nil {
		return res, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok || !jwtToken.Valid {
		return res, fmt.Errorf("invalid claims")
	}

	res, err = parseClaims(claims)

	if err != nil {
		return res, err
	}

	return res, nil
}

func keyFunction(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	var hmacSecret = []byte(os.Getenv("JWT_PRIVATE_KEY"))

	if !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return hmacSecret, nil
}

func parseClaims(claims jwt.MapClaims) (entities.GetUserResponse, error) {
	var user entities.GetUserResponse

	parsedUser, ok := claims["student_id"].(int64)

	if !ok {
		return user, errors.New("email not found")
	}

	user.StudentID = parsedUser

	return user, nil
}
