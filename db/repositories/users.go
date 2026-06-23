package db

import (
	"database/sql"
	"fmt"
	"Authingo/models"
)
type UserRepository interface{
		GetByID(id int64) (*models.User,error) 
		Create(username string,email string,hashedPassword string) (int64,error)
		GetAll() ([]*models.User,error)
		DeleteById(id int64) error
		GetByEmail(email string) (*models.User,error)

}
type UsserRepositoryImpl struct{
		db *sql.DB
}
func NewUserRepository(_db *sql.DB) UserRepository{
	return &UsserRepositoryImpl{
			db:_db,
	}
}
func(u *UsserRepositoryImpl) GetAll() ([]*models.User,error){
	fmt.Println("fetching all users in UserRepository")
	query:="SELECT id,username,email,password,created_at,updated_at FROM users"
	rows,err:=u.db.Query(query)
	if err!=nil{
		fmt.Println("error fetching users",err)
		return nil,err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next(){
		user:=&models.User{}
		err:=rows.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)
		if err!=nil{
			fmt.Println("error scanning user",err)
			return nil,err
		}
		users=append(users,user)
	}

	if err:=rows.Err();err!=nil{
		fmt.Println("error iterating rows",err)
		return nil,err
	}

	fmt.Println("fetched users succesfully, count:",len(users))
	return users,nil
}
func(u *UsserRepositoryImpl) DeleteById(id int64) error{
	fmt.Println("deleting user in UserRepository, id:",id)
	query:="DELETE FROM users WHERE id=?"
	result,err:=u.db.Exec(query,id)
	if err!=nil{
		fmt.Println("error deleting user",err)
		return err
	}

	rowsaffected,rowerr:=result.RowsAffected()
	if rowerr!=nil{
		fmt.Println("error getting rows affected",rowerr)
		return rowerr
	}
	if rowsaffected==0{
		fmt.Println("no rows affected, user not found, id:",id)
		return sql.ErrNoRows
	}

	fmt.Println("user deleted succesfully, rows affected:",rowsaffected)
	return nil
}
func(u *UsserRepositoryImpl) Create(username string,email string,hashedPassword string) (int64,error){
	query:="INSERT INTO users(username,email,password) VALUES(?,?,?)"
	result,err:=u.db.Exec(query,username,email,hashedPassword)
	if err!=nil{
		fmt.Println("error creating user",err)
		return 0,err
	}
	rowsaffected,rowerr:=result.RowsAffected()
	if rowerr!=nil{
		fmt.Println("error getting rows affected",rowerr)
		return 0,rowerr
	}
	if rowsaffected==0{
		fmt.Println("no rows affected,user are not created")
		return 0,fmt.Errorf("user was not created")
	}
	newId,idErr:=result.LastInsertId()
	if idErr!=nil{
		fmt.Println("error getting last insert id",idErr)
		return 0,idErr
	}
	fmt.Println("user created succesfully, id:",newId)
	return newId,nil
}
func(u *UsserRepositoryImpl) GetByID (id int64) (*models.User,error) {
	fmt.Println("fetching user in UserRepository, id:",id)
	//prepare quesy
	query:="SELECT id,username,email,password,created_at,updated_at FROM users WHERE id=?"
//wexecute query
	row:=u.db.QueryRow(query,id)
	//process the result
	user:=&models.User{}
	err:=row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)
	if err!=nil{
		if err==sql.ErrNoRows{
			fmt.Println("user not found, id:",id)
			return nil,err
		}else{
			fmt.Println("error scanning user",err)
			return nil,err
		}
	}
	//print the user detail
	fmt.Println("user fetched succesfully",user)
	return user,nil
}

func(u *UsserRepositoryImpl) GetByEmail(email string) (*models.User,error){
	fmt.Println("fetching user by email in UserRepository")
	query:="SELECT id,username,email,password FROM users WHERE email=?"
	row:=u.db.QueryRow(query,email)

	user:=&models.User{}
	err:=row.Scan(&user.Id,&user.Username,&user.Email,&user.Password)
	if err!=nil{
		if err==sql.ErrNoRows{
			fmt.Println("user not found with email:",email)
			return nil,err
		}
		fmt.Println("error scanning user",err)
		return nil,err
	}
	fmt.Println("user fetched succesfully by email",user)
	return user,nil
}