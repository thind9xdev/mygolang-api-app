package routes

import (
	"mygolangapp/controllers"
	"mygolangapp/middleware"
	"net/http"
)

func SetupRoutes() {
	http.Handle("/user", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetUser)))
	http.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetAllUsers)))
}
