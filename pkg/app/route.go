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

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(JWTAuth())
	adminRoutes.GET("/users", GetUsers)
	adminRoutes.GET("/user/:id", GetUser)
	adminRoutes.PUT("/user/:id", UpdateUser)
	adminRoutes.POST("/user/role", CreateRole)
	adminRoutes.GET("/user/roles", GetRoles)
	adminRoutes.PUT("/user/role/:id", UpdateRole)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
