package router

import (
	"Authingo/controllers"
	"Authingo/dto"
	"Authingo/middlewares"
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
	
	r.With(middlewares.JWTAuthMiddleware).Get("/users/{id}",ur.UserController.GetUserByID)
	r.Get("/users",ur.UserController.GetAllUsers)
	r.Delete("/users/{id}",ur.UserController.DeleteUserById)
	
	r.Post("/signup", middlewares.ValidateBody[dto.SignupRequestDTO](ur.UserController.CreateUser))
	r.Post("/login", middlewares.ValidateBody[dto.LoginRequestDTO](ur.UserController.LoginUser))
}