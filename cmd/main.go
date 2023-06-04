package main

import (
	"log"

	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/db"
	"github.com/prasanth-pn/GolangAuthenticationWithTwilioJWT/pkg/di"
)

const (
	Port = ":9090"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("error from initialize the server %s", err)
		}
	}()
	_ = db.ConnectGorm()

	server, err := di.InitializeEvent()
	if err != nil {
		panic(err)
	} else {
		server.Start(Port)
	}

}
