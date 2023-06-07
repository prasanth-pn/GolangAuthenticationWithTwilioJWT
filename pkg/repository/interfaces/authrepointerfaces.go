package interfaces

import (
	"context"

	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
)

type Authrepositoryinterface interface {
	Register(ctx context.Context, user domain.User) error
	FindUser(ctx context.Context,username string)(domain.User,error)
	StoreOtp(ctx context.Context,otp string)error
}
