package repositories

import (
	"ddd-go/modules/entities"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) entities.UsersRepository {
	return &userRepo{
		Db: db,
	}
}

func (r *userRepo) Register(req *entities.UsersResgisterReq) (*entities.UsersRegisterRes, error) {
	query := `
	INSERT INTO "users"(
		"username",
		"password"
	)
	VALUES ($1, $2)
	RETURNING "id", "username";
	`

	user := new(entities.UsersRegisterRes)

	rows, err := r.Db.Queryx(query, req.Username, req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	defer r.Db.Close()

	return user, nil
}
