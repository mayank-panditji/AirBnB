package app

import (
	"Authingo/config/env"
	db "Authingo/db/repositories"
	"Authingo/router"
	"fmt"
	"net/http"
	"time"
	"Authingo/controllers"
	"Authingo/services"
	dbConfig "Authingo/config/env/db"
	repo "Authingo/db/repositories"
)

// config holds the cofiguration for the server
type Config struct{
	Addr string
}
type Application struct{
    Config Config
	Store db.Storage
}
func NewConfig() Config{
	port:=env.GetString("PORT","8080")
	return Config{
		Addr:port,
	}
}
func NewApplication (cfg Config) *Application{
	return &Application{
		Config:cfg,
		Store:*db.NewStorage(),
	}
}

func (app *Application) Run() error{
	db,err:=dbConfig.SetupDB()
	if err!=nil{
		fmt.Println("Error setting up database",err)
	}
		ur := repo.NewUserRepository(db)
	rr := repo.NewRoleRepository(db)
	rpr := repo.NewRolePermissionRepository(db)
	urr := repo.NewUserRoleRepository(db)
	us := services.NewUserService(ur)
	rs := services.NewRoleService(rr, rpr, urr)
	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)
	uRouter := router.NewUserRouter(uc)
	rRouter := router.NewRoleRouter(rc)
	server:=&http.Server{
			Addr:app.Config.Addr,
			Handler:router.SetupRouter(uRouter,rRouter), //TODO setup a chi router and put it here
			ReadTimeout: 10*time.Second,
			WriteTimeout: 10*time.Second,
	}
	fmt.Println("Starting server on",app.Config.Addr)
	return server.ListenAndServe()
}