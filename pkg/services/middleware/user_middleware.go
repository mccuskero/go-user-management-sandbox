package middleware

import (
	//	"log"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"

	//	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/mccuskero/go-user-management-sandbox/pkg/config"
	"github.com/mccuskero/go-user-management-sandbox/pkg/initializer"
	"github.com/mccuskero/go-user-management-sandbox/pkg/models"
	"github.com/mccuskero/go-user-management-sandbox/pkg/utils"
)

// TODO: do we want to manage DB better? initializer is used to initialize
type UserMiddleware struct {
	DB *gorm.DB
}

func NewUserMiddleware(initializer *initializer.Initializer) *UserMiddleware {
	return &UserMiddleware{
		DB: initializer.DB,
	}
}

func (um *UserMiddleware) DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User

		result := um.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
