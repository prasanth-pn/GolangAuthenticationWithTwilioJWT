package db

import (
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectGorm()*gorm.DB{
db,_:=gorm.Open(postgres.Open(DBsource),&gorm.Config{
	SkipDefaultTransaction: true,
})
db.AutoMigrate(domain.User{})


	return db
}