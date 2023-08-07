package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ServeApplicationRoute() {
	router := gin.Default()

	authRoutes := router.Group("/auth/user")
	// registration route
	authRoutes.POST("/register", UserRegister)
	// login route
	authRoutes.POST("/login", UserLogin)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
