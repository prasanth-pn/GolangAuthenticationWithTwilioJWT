package repository

import "database/sql"
import "github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/repository/interfaces"

type Database struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) interfaces.Authrepositoryinterface{
	return &Database{db}
}
