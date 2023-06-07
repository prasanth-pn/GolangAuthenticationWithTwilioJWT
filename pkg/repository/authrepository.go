package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/repository/interfaces"
)

type Database struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) interfaces.Authrepositoryinterface {
	return &Database{db}
}

func (c *Database) Register(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (user_name,password,verification)VALUES($1,$2,$3);`
	err := c.DB.QueryRow(query, user.UserName, user.Password,false)

	return err.Err()

}
func (c *Database) FindUser(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	query := `SELECT user_name,password,verification FROM users WHERE user_name=$1;`
	err := c.DB.QueryRow(query, username).Scan(&user.UserName, &user.Password, &user.Verification)
	ctx.Done()
	return user, err

}
func (c *Database)StoreOtp(ctx context.Context,otp string)error{
	query:=`INSERT INTO storeotp(user_id,otp)VALUES($1,$2);`
	err:=err.
}
