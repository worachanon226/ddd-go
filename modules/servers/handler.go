package servers

import (
	_usersHttp "ddd-go/modules/users/controllers"
	_usersRepositories "ddd-go/modules/users/repositories"
	_usersUsecases "ddd-go/modules/users/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
v1:=
	s.App.Group("/v1")

	usersGroup := v1.Group("/user")
	userRepository := _usersRepositories.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecases.NewUsersUsecase(userRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	s.App.Use(func(c *fiber.Ctx) error{
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})
	
	return nil
}