package services

import (
	"Authingo/config/env"
	db "Authingo/db/repositories"
	"Authingo/models"
	"Authingo/utils"
	"fmt"
		"Authingo/dto"
	"github.com/golang-jwt/jwt/v4"
)
type UserService interface {
	GetUserByID(id int64) (*models.User,error)
	GetAllUsers() ([]*models.User,error)
	DeleteUserById(id int64) error
	CreateUser(payload *dto.SignupRequestDTO) (*models.User,error)
	LoginUser(payload *dto.LoginRequestDTO) (string,error)
}
type UserServiceImpl struct {
	userRepository db.UserRepository
}
func NewUserService(_userRepository db.UserRepository) UserService{
	return &UserServiceImpl{
		userRepository:_userRepository,
	}
}
func (u *UserServiceImpl) GetUserByID(id int64) (*models.User,error){
	fmt.Println("fetching user in userservice, id:",id)
	user,err:=u.userRepository.GetByID(id)
	if err!=nil{
		fmt.Println("error fetching user in userservice",err)
		return nil,err
	}
	return user,nil
}
func (u *UserServiceImpl) GetAllUsers() ([]*models.User,error){
	fmt.Println("fetching all users in userservice")
	users,err:=u.userRepository.GetAll()
	if err!=nil{
		fmt.Println("error fetching all users in userservice",err)
		return nil,err
	}
	return users,nil
}
func (u *UserServiceImpl) DeleteUserById(id int64) error{
	fmt.Println("deleting user in userservice, id:",id)
	err:=u.userRepository.DeleteById(id)
	if err!=nil{
		fmt.Println("error deleting user in userservice",err)
		return err
	}
	return nil
}
func (u *UserServiceImpl) CreateUser(payload *dto.SignupRequestDTO) (*models.User,error){
	fmt.Println("creating user in userservice")
	hashedPassword,err:=utils.HashPassword(payload.Password)
	if err!=nil{
		fmt.Println("error hashing password",err)
		return nil,err
	}
	newId,err:=u.userRepository.Create(
		payload.Username,
		payload.Email,
		hashedPassword,
	)
	if err!=nil{
		fmt.Println("error creating user in userservice",err)
		return nil,err
	}
	user,err:=u.userRepository.GetByID(newId)
	if err!=nil{
		fmt.Println("error fetching newly created user in userservice",err)
		return nil,err
	}
	return user,nil
}
func (u *UserServiceImpl) LoginUser(payload *dto.LoginRequestDTO) (string,error){
	email:=payload.Email
	password:=payload.Password
user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return "", err
	}

	if user == nil {
		fmt.Println("No user found with the given email")
		return "", fmt.Errorf("no user found with email: %s", email)
	}

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", nil
	}

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil
}