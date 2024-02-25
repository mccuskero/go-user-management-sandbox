package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/mccuskero/go-user-management-sandbox/pkg/services/controllers"
	"github.com/mccuskero/go-user-management-sandbox/pkg/services/middleware"
)

type AuthRouter struct {
	authController controllers.AuthController
	userMiddleware middleware.UserMiddleware
}

func NewAuthRouter(authController controllers.AuthController, userMiddleware middleware.UserMiddleware) AuthRouter {
	return AuthRouter{
		authController: authController,
		userMiddleware: userMiddleware,
	}
}

func (ar *AuthRouter) Initialize(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", ar.authController.RegisterUser)
	router.POST("/login", ar.authController.SignInUser)
	router.GET("/profile", ar.authController.UserProfile)
	router.GET("/refresh", ar.authController.RefreshAccessToken)
	router.GET("/logout", ar.userMiddleware.DeserializeUser(), ar.authController.LogoutUser)
}
