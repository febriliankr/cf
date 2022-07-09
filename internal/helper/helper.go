package helper

import (
	"math/rand"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

func GenerateRandom3Digits() int {
	return (int)((rand.Float64() * 9999) + 1)

}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func GetCurrentTimestamp() int64 {
	return time.Now().UnixMilli()
}
