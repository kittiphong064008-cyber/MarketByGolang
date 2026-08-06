package main

import (
	configs "cleanarch/config"
	"cleanarch/databases"
	"cleanarch/server"
	"log"
)

func main() {
	cfg := configs.Load()
	db, err := databases.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app := server.NewServer(db)
	log.Fatal(app.Listen(cfg.App.Host + ":" + cfg.App.Port))
}
