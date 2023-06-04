package interfaces

import (
	"context"

	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
)

type Authusecase interface {
	Register(ctx context.Context, user domain.User) error
	HashPassword(password string) string
}
