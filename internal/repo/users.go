package repo

import "github.com/febriliankr/go-cfstore-api/internal/entities"

const queryGetUser = `
SELECT
	student_id,
	password
FROM
	users
WHERE
	student_id = $1
LIMIT 1
`

const queryCreateUser = `
INSERT INTO
	users (student_id, password, created_at)
VALUES
	($1, $2, now())
RETURNING
	student_id
`

func (r *KantinDB) CreateUser(in entities.CreateUserRequest) (entities.CreateUserResponse, error) {

	var res entities.CreateUserResponse

	err := r.Db.QueryRowx(queryCreateUser, in.StudentID, in.Password).StructScan(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *KantinDB) GetUser(in entities.GetUserRequest) (entities.GetUserResponse, error) {
	var res entities.GetUserResponse

	err := r.Db.QueryRowx(queryGetUser, in.StudentID).StructScan(&res)

	if err != nil {
		return res, err
	}

	return res, nil

}
