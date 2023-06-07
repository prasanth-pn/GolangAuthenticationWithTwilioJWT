package usecase

import (
	"context"

	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
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
func (c *Authusecase)Register(ctx context.Context,users domain.User)error{
	err:=c.authrepo.Register(ctx,users)

	return err
}
func (c *Authusecase)FindUser(ctx context.Context,username string)(domain.User,error){
	user,err:=c.authrepo.FindUser(ctx,username)

	return user,err

}
func (c *Authusecase)StoreOtp(ctx context.Context,otp string)error{
	err:=c.authrepo.StoreOtp(ctx,otp)
	return err

}
