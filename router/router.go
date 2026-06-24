package router

import (
	"Authingo/controllers"
	"Authingo/middlewares"
	"Authingo/utils"

	"github.com/go-chi/chi/v5"
)
type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router, RoleRouter Router) *chi.Mux {
	chirouter := chi.NewRouter()
	chirouter.Use(middlewares.RequestLogger)
	chirouter.Use(middlewares.RateLimitMiddleware)
	chirouter.Get("/ping", controllers.PingHandler)
	chirouter.HandleFunc("/fakestoreservice/*", utils.ProxytoService("https://fakestoreapi.com", "/fakestoreservice/api"))

	// for _, r := range routers {
	// 	r.Register(chirouter)
	// }
	UserRouter.Register(chirouter)
	RoleRouter.Register(chirouter)

	return chirouter
}