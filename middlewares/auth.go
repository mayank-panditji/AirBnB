package middlewares

import (
	"Authingo/config/env"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		authHeader:=r.Header.Get("Authorization")
		if authHeader==""{
			http.Error(w,"Authorization header required",http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader,"Bearer") {
			http.Error(w,"Authorization header must startwith bearer",http.StatusUnauthorized)
			return
		}
		token:=strings.TrimPrefix(authHeader,"Bearer ")
		if token==""{
			http.Error(w,"Authorization header must startwith bearer",http.StatusUnauthorized)
			return
		}
		claims:=jwt.MapClaims{}
		_,err:=jwt.ParseWithClaims(token,&claims,func(token *jwt.Token)(interface{},error){
			return []byte(env.GetString("JWT_SECRET","TOKEN")),nil
		})
		if err!=nil{
			http.Error(w,"Invalid token"+err.Error(),http.StatusUnauthorized)
			return
		}
		userId,okId:=claims["id"].(float64)
			email,okEmail:=claims["email"].(string)
		if !okId || !okEmail{
			http.Error(w,"Invalid token",http.StatusUnauthorized)
			return
		}
		fmt.Println("Authentication user ID:",int64(userId),"Email:",email)
		ctx:=context.WithValue(r.Context(),"userId",int64(userId))
		ctx=context.WithValue(ctx,"email",email)	
		next.ServeHTTP(w,r)
	
	})
}