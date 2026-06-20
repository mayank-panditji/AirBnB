package controllers

import (
	"Authingo/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)
type UserController struct{
	UserService services.UserService
}
func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService:_userService,
	}
}
func (uc *UserController) GetUserByID(w http.ResponseWriter,r *http.Request){
	fmt.Println("user fetching end point")
	uc.UserService.GetUserByID()
	w.Write([]byte("user registerd endpoint"))
}
func (uc *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	fmt.Println("get all users endpoint")
	users,err:=uc.UserService.GetAllUsers()
	if err!=nil{
		http.Error(w,"error fetching users",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(users)
}
func (uc *UserController) DeleteUserById(w http.ResponseWriter,r *http.Request){
	idParam:=chi.URLParam(r,"id")
	id,err:=strconv.ParseInt(idParam,10,64)
	if err!=nil{
		http.Error(w,"invalid id",http.StatusBadRequest)
		return
	}
	fmt.Println("delete user endpoint, id:",id)
	err=uc.UserService.DeleteUserById(id)
	if err!=nil{
		http.Error(w,"error deleting user",http.StatusInternalServerError)
		return
	}
	w.Write([]byte("user deleted succesfully"))
}