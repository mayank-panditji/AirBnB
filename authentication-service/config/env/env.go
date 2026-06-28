package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
func Load(){
 err:=godotenv.Load()
 if err!=nil{
  fmt.Println("Error loading .env file")
 }
}
func GetString(key string ,fallback string ) string{
	
	value,ok:=os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
func GetInt(key string ,fallback int)int{

	value,ok:=os.LookupEnv(key);
	if !ok{
		return fallback
	}
	intvalue,err:=strconv.Atoi(value)
	if err!=nil{
		fmt.Printf("error converting %s to int :%v\n",key,err)
		return fallback
	}
	return intvalue
}
func GetBool(key string ,fallback bool)bool{
	
	value,ok:=os.LookupEnv(key);
	if !ok{
		return fallback
	}
	boolvalue,err:=strconv.ParseBool(value)
	if err!=nil{
		fmt.Printf("error converting %s to bool :%v\n",key,err)
		return fallback
	}
	return boolvalue
}
