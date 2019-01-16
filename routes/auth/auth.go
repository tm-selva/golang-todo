package auth

import (
	authservice "../../services/auth"

	"github.com/gin-gonic/gin"
)

/*
auth routes
*/
func Routes(authRoute *gin.Engine) {

	auth := authRoute.Group("/auth")
	{
		auth.POST("/login", authservice.Login)
		auth.POST("/register", authservice.Register)
	}

}
