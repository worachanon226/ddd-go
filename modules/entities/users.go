package entities

type UsersUsecase interface {
	Register(req *UsersResgisterReq) (*UsersRegisterRes, error)
}

type UserRepository interface {
	Register(req *UsersResgisterReq) (*UsersRegisterRes, error)
}

type UsersResgisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRegisterRes struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
