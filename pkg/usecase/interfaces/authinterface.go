package interfaces

import (
	"context"

	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
)

type Authusecase interface {
	Register(ctx context.Context, user domain.User) error
	HashPassword(password string) string
	FindUser(ctx context.Context,username string)(domain.User,error)
	StoreOtp(ctx context.Context,otp string)error
}
