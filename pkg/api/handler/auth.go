package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/usecase/interfaces"
)

type AuthHandler struct {
	authusecase interfaces.Authusecase
}

func NewAuthHandler(usecase interfaces.Authusecase) *AuthHandler {
	return &AuthHandler{usecase}
}

func (cr *AuthHandler) Register(c *gin.Context) {
var user domain.User

if err:=c.BindJSON(&user);err!=nil{
	c.JSON(401,gin.H{
		"message":"error from geting user registration details",
		"error":err.Error(),
	})
}
user.Password=cr.authusecase.HashPassword(user.Password)
c.JSON(201,gin.H{
	"message":"sucessfully registered",
	"data":user,
})
}
