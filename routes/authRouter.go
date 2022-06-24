package routes

import (
	controller "github.com/vikas-gautam/golang-jwt-project/controllers"
)

func AuthRoutes() {
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())

}
