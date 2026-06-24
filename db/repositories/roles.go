package db
import (
	"database/sql"
	"Authingo/models"
	"fmt"
)
type RoleRepository interface {
	GetRoleById(id int64) (*models.Role,error)
	GetRoleByName(name string) (*models.Role,error)
	GetAllRoles() ([]*models.Role,error)
	CreateRole(name string,description string) (int64,error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64,name string,description string) error
}
type RoleRepositoryImpl struct {
	db *sql.DB
}
func NewRoleRepository(_db *sql.DB) UserRepository{
	return &UsserRepositoryImpl{
			db:_db,
	}
}

func (r *RoleRepositoryImpl) GetRoleById(id int64) (*models.Role,error){
	query:="SELECT id,name,description,created_at,updated_at FROM roles WHERE id=?"
	row := r.db.QueryRow(query,id)
	role:=&models.Role{}
	err:=row.Scan(&role.Id,&role.Name,&role.Description,&role.CreatedAt,&role.UpdatedAt)
	if err!=nil{
		fmt.Println("error scanning role",err)
		return nil,err
	}
	return role,nil
}
func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role,error){
	query:="SELECT id,name,description,created_at,updated_at FROM roles WHERE name=?"
	row := r.db.QueryRow(query,name)
	role:=&models.Role{}
	err:=row.Scan(&role.Id,&role.Name,&role.Description,&role.CreatedAt,&role.UpdatedAt)
	if err!=nil{
		fmt.Println("error scanning role",err)
		return nil,err
	}
	return role,nil
}
func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role,error){
	query:="SELECT id,name,description,created_at,updated_at FROM roles"
	rows,err:=r.db.Query(query)
	if err!=nil{
		fmt.Println("error fetching roles",err)
		return nil,err
	}
	defer rows.Close()
	var roles []*models.Role
	for rows.Next(){
		role:=&models.Role{}
		err:=rows.Scan(&role.Id,&role.Name,&role.Description,&role.CreatedAt,&role.UpdatedAt)
		if err!=nil{
			fmt.Println("error scanning role",err)
			return nil,err
		}
		roles=append(roles,role)
	}
	if err:=rows.Err();err!=nil{
		fmt.Println("error iterating rows",err)
		return nil,err
	}
	return roles,nil
}
func (r *RoleRepositoryImpl) CreateRole(name string,description string) (*models.Role,error){
	query:="INSERT INTO roles(name,description,created_at,updated_at) VALUES(?,?,NOW(),NOW())"
	result,err:=r.db.Exec(query,name,description)
	if err!=nil{
		fmt.Println("error creating role",err)
		return nil,err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		fmt.Println("error creating role",err)
		return nil,err
	}
	return &models.Role{
		Id: id,
		Name: name,
		Description: description,
		CreatedAt: "",
		UpdatedAt: "",
	},nil
}
func (r *RoleRepositoryImpl) DeleteRoleById(id int64) error{
	query:="DELETE FROM roles WHERE id=?"
	result,err:=r.db.Exec(query,id)
	if err!=nil{
		fmt.Println("error deleting role",err)
		return err
	}
	rowsAffected,err:=result.RowsAffected()
	if err!=nil{
		fmt.Println("error getting rows affected",err)
		return err
	}
	if rowsAffected==0{
		fmt.Println("no rows affected, role not found, id:",id)
		return sql.ErrNoRows
	}
	fmt.Println("role deleted succesfully, rows affected:",rowsAffected)
	return nil
}
func (r *RoleRepositoryImpl) UpdateRole(id int64,name string,description string) (*models.Role,error){
	query:="UPDATE roles SET name=?,description=? update_at=NOW() WHERE id=?"
	_,err:=r.db.Exec(query,name,description,id)
	if err!=nil{
		
		return nil,err
	}
	return &models.Role{
		Id: id,
		Name: name,
		Description: description,
		CreatedAt: "",
		UpdatedAt: "",
	},nil
}