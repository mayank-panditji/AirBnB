package config

import (
	env "Authingo/config/env"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)
func SetupDB() (*sql.DB,error){
cfg:=mysql.NewConfig()
	cfg.User=env.GetString("DB_USER","root")
	cfg.Passwd=env.GetString("DB_PASSWORD","root")
	cfg.DBName=env.GetString("DB_NAME","auth_dev")
	cfg.Net=env.GetString("DB_NET","tcp")
	cfg.Addr=env.GetString("DB_ADDR","127.0.0.1:3306")
	fmt.Println("connecting to database",cfg.DBName,cfg.FormatDSN())
	db,err:=sql.Open("mysql",cfg.FormatDSN())
	if err!=nil{
		fmt.Println("error connecting to database",err)
		return nil,err
	}
	pingErr:=db.Ping()
	if pingErr!=nil{
		fmt.Println("error pinging database",pingErr)
		return nil,err
	}
	fmt.Println("connected to database",cfg.DBName)
	return db,nil
}
