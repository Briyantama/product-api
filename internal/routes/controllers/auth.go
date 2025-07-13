package controllers

import (
	"net/http"
	"test-case-vhiweb/internal/models"
	"test-case-vhiweb/internal/routes/usecase"
	"test-case-vhiweb/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUC usecase.UserUsecase
}

func NewUserController(uc usecase.UserUsecase) *UserController {
	return &UserController{UserUC: uc}
}

func (ctl *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err)
		return
	}

	if err := ctl.UserUC.Register(c, &user); err != nil {
		c.Error(err)
		return
	}

	utils.JSONResponse(c, http.StatusOK, gin.H{"message": "User registered"})
}

func (ctl *UserController) Login(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}

	id, err := ctl.UserUC.Login(c, body.Email, body.Password)
	if err != nil {
		c.Error(err)
		return
	}

	token, err := utils.GenerateJWT(*id)
	if err != nil {
		c.Error(err)
		return
	}

	c.SetCookie("token", token, 3600*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
	})
}
