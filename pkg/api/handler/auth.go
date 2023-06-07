package handler

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/config"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	authusecase interfaces.Authusecase
}

func NewAuthHandler(usecase interfaces.Authusecase) *AuthHandler {
	return &AuthHandler{usecase}
}

func (cr *AuthHandler) Register(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(401, gin.H{
			"message": "error from geting user registration details",
			"error":   err.Error(),
		})
	}
	user.Password = cr.authusecase.HashPassword(user.Password)
	err := cr.authusecase.Register(c, user)
	if err != nil {
		c.JSON(401, gin.H{
			"message ": "failed to reginster",
			"error":    err.Error(),
		})
	}
	c.JSON(201, gin.H{
		"message": "sucessfully registered",
		"data":    user,
	})
}
func (cr *AuthHandler) Login(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(401, gin.H{
			"message ": "failed to fetch the data from user",
			"error":    err.Error(),
		})
		return
	}
	User, err := cr.authusecase.FindUser(c, user.UserName)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "faile to login user credentials are not valid",
			"err":     err.Error(),
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(user.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Entered password is incorrect",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message ": "sucessfully login",
		"data":     user,
	})

}
func (cr *AuthHandler) SendOTP(c *gin.Context) {
	mobile:=c.Query("mobilenumber")
	k, err := config.TwilioSendOtp(mobile)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "otp is not send please try again",
			"error":   err.Error(),
		})
		return
	}
	reg:=`\d{6}`
	regex,_:=regexp.Compile(reg)
	code:=regex.FindString(k)
	err:=cr.authusecase.StoreOtp(c,code)
	c.JSON(200, gin.H{
		"messsage": "otp send sucessfully",
	})
	c.Abort()
}
