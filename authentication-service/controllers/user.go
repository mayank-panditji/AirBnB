package controllers

import (
	"Authingo/services"
	"Authingo/utils"
	"fmt"
	"net/http"
	"strconv"
	"Authingo/dto"
	"Authingo/middlewares" 
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
	idParam:=chi.URLParam(r,"id")
	id,err:=strconv.ParseInt(idParam,10,64)
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"invalid id",err)
		return
	}
	fmt.Println("user fetching endpoint, id:",id)
	user,err:=uc.UserService.GetUserByID(id)
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusNotFound,"user not found",err)
		return
	}
	response:=dto.UserResponseDTO{
		Id:int64(user.Id),
		Username:user.Username,
		Email:user.Email,
		CreatedAt:user.CreatedAt,
		UpdatedAt:user.UpdatedAt,
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"user fetched succesfully",response)
}
func (uc *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	fmt.Println("get all users endpoint")
	users,err:=uc.UserService.GetAllUsers()
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"error fetching users",err)
		return
	}
	responses:=make([]dto.UserResponseDTO,0,len(users))
	for _,user:=range users{
		responses=append(responses,dto.UserResponseDTO{
			Id:int64(user.Id),
			Username:user.Username,
			Email:user.Email,
			CreatedAt:user.CreatedAt,
			UpdatedAt:user.UpdatedAt,
		})
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"users fetched succesfully",responses)
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
func (uc *UserController) CreateUser(w http.ResponseWriter,r *http.Request){
	fmt.Println("create user endpoint")
	// var payload dto.SignupRequestDTO
	payload := r.Context().Value(middlewares.PayloadContextKey).(dto.SignupRequestDTO)
	user,err:=uc.UserService.CreateUser(&payload)
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"error creating user",err)
		return
	}
	response:=dto.UserResponseDTO{
		Id:int64(user.Id),
		Username:user.Username,
		Email:user.Email,
		CreatedAt:user.CreatedAt,
		UpdatedAt:user.UpdatedAt,
	}
	utils.WriteJsonSuccessResponse(w,http.StatusCreated,"user created succesfully",response)
}
func (uc *UserController) LoginUser(w http.ResponseWriter,r *http.Request){
	fmt.Println("login user endpoint")
	// var payload dto.LoginRequestDTO
	// if jsonErr:=utils.ReadJson(r,&payload); jsonErr!=nil{
	// 	utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Something went wrong while logging",jsonErr)
	// 	return
	// }
	// if validationErr:=utils.Validator.Struct(payload);validationErr!=nil{
	// 	utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"invalid input data",validationErr)
	// 	return
	// }
	payload := r.Context().Value(middlewares.PayloadContextKey).(dto.LoginRequestDTO)
	jwtToken,err:=uc.UserService.LoginUser(&payload)
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"error logging in user",err)
		return
	}
	
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"user logged in succesfully",jwtToken,)
}