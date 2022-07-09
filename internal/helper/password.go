package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//Generate a salted hash for the input string
func Generate(unhashedPw string) (string, error) {

	if len(unhashedPw) == 0 {
		return "", errors.New("password cannot be empty")
	}

	saltedBytes := []byte(unhashedPw)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func Compare(existingHash string, incomingHash string) error {
	incoming := []byte(incomingHash)
	existing := []byte(existingHash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
