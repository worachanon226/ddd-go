package usecases

import (
	"ddd-go/modules/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (u *usersUse) Register(req *entities.UsersResgisterReq) (*entities.UsersRegisterRes, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	user, err := u.UsersRepo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
