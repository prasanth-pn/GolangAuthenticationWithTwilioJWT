package interfaces

import (
	"context"

	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
)

type Authrepositoryinterface interface {
	Register(ctx context.Context, user domain.User) error
}
