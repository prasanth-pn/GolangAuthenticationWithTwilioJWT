package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/api/handler"
)

type ServeHTTP struct {
	engine *gin.Engine
}

func NewServeHTTP(AuthHandler *handler.AuthHandler) *ServeHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(cors.Default())
	userapi := engine.Group("/user")
	userapi.POST("/register", AuthHandler.Register)
	userapi.GET("/login",AuthHandler.Login)
	userapi.POST("/otp",AuthHandler.SendOTP)

	return &ServeHTTP{engine}
}

func (c *ServeHTTP) Start(n string) {
	c.engine.Run(n)
}
