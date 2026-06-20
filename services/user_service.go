package services

import (
	db "Authingo/db/repositories"
	"Authingo/models"
	"fmt"
)
type UserService interface {
	GetUserByID() error
	GetAllUsers() ([]*models.User,error)
	DeleteUserById(id int64) error
}
type UserServiceImpl struct {
	userRepository db.UserRepository
}
func NewUserService(_userRepository db.UserRepository) UserService{
	return &UserServiceImpl{
		userRepository:_userRepository,
	}
}
func (u *UserServiceImpl) GetUserByID() error{
	fmt.Println("creating user in userservice")
	u.userRepository.GetByID()
	return nil
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