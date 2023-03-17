package models

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func ConnectPostegreSQL() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Addr:     os.Getenv("ADDR"),
		Database: os.Getenv("DATABASE"),
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
