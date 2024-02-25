package controllers

import (
	//	"fmt"
	//	"net/http"
	//	"strings"
	//	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	// "github.com/mccuskero/go-user-management-sandbox/pkg/utils"
	// "github.com/mccuskero/go-user-management-sandbox/pkg/models"
	// "github.com/mccuskero/go-user-management-sandbox/pkg/config"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) AddEntity(ctx *gin.Context) {
	// how do we get the user id?
	// where is the token? can it hold the user id?
	//
}

func (uc *UserController) GetEntities(ctx *gin.Context) {

}
