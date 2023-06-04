package repository

import (
	"context"
	"database/sql"

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
	query := `INSERT INTO users (user_name,password)VALUES($1,$2);`
	err := c.DB.QueryRow(query, user.UserName, user.Password)

	return err.Err()

}
