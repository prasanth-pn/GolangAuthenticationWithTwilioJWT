//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/api"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/api/handler"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/db"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/repository"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/usecase"
)

func InitializeEvent() (*http.ServeHTTP, error) {
	wire.Build(db.ConnectDB,
		repository.NewAuthRepository,
		usecase.NewAuthUsecase,
		handler.NewAuthHandler,
		http.NewServeHTTP)
	return &http.ServeHTTP{}, nil

}
