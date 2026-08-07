package main

import (
	configs "cleanarch/config"
	"cleanarch/databases"
	"cleanarch/server"
	"fmt"
	"log"
)

func test() {
	defer fmt.Println("HELLO1") //first in last out
	defer fmt.Println("HELLO2")
}

func main() {
	cfg := configs.Load()
	db, err := databases.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	test()
	defer db.Close()
	app := server.NewServer(db)
	log.Fatal(app.Listen(cfg.App.Host + ":" + cfg.App.Port))
}
