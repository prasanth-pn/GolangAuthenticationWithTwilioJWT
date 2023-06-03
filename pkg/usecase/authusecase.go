package usecase

import (
	repointerfaces "github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/repository/interfaces"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type Authusecase struct {
	authrepo repointerfaces.Authrepositoryinterface
}

func NewAuthUsecase(authrep repointerfaces.Authrepositoryinterface) interfaces.Authusecase {
	return &Authusecase{authrep}
}
func (cr Authusecase) HashPassword(password string) string {

	hashedPassword, _:= bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	
	return string(hashedPassword)
}
