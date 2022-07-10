package usecases

import (
	"errors"

	"github.com/febriliankr/go-cfstore-api/internal/entities"
	"github.com/febriliankr/go-cfstore-api/internal/helper"
)

func (uc *KantinUsecase) CreateUser(in entities.CreateUserRequest) (entities.CreateUserResponse, error) {
	var res entities.CreateUserResponse

	existingUser, _ := uc.repo.GetUser(entities.GetUserRequest{
		StudentID: in.StudentID,
		Password:  in.Password,
	})

	if existingUser.StudentID != 0 {
		return res, errors.New("user already exists")
	}

	if err := validateStudentID(in.StudentID); err != nil {
		return res, err
	}

	hashedPW, err := helper.Generate(in.Password)

	if err != nil {
		return res, err
	}

	in.Password = hashedPW

	return uc.repo.CreateUser(in)
}

func validateStudentID(studentID int64) error {

	if studentID > 99999 || studentID < 10000 {
		return errors.New("student ID length is invalid, must be 5 digits")
	}

	twoLastDigits := studentID % 100
	firstDigit := studentID % 100000 / 10000
	secondDigit := studentID % 10000 / 1000
	thirdDigit := studentID % 1000 / 100

	validID := firstDigit+secondDigit+thirdDigit == twoLastDigits

	if !validID {
		return errors.New("student ID is invalid")
	}

	return nil
}

func (uc *KantinUsecase) GetUser(in entities.GetUserRequest) (entities.GetUserResponse, error) {

	var res entities.GetUserResponse

	res, err := uc.repo.GetUser(in)

	if err != nil {
		return res, errors.New("user does not exist")
	}

	// hashedInputPW, err := helper.Generate(in.Password)

	if err != nil {
		return res, err
	}

	err = helper.Compare(res.Password, in.Password)

	if err != nil {
		return res, err
	}

	return res, nil
}
