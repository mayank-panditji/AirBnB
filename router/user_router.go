package router

import (
	"Authingo/controllers"

	"github.com/go-chi/chi/v5"
)
type UserRouter struct {
	UserController *controllers.UserController
}
func NewUserRouter(_userController *controllers.UserController) Router{
	return &UserRouter{
		UserController:_userController,
	}
}
func (ur *UserRouter) Register(r chi.Router){
	r.Get("/profile",ur.UserController.GetUserByID)
	r.Get("/users",ur.UserController.GetAllUsers)
	r.Delete("/users/{id}",ur.UserController.DeleteUserById)
}