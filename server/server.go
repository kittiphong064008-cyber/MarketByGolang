package server

import (
	configs "cleanarch/config"
	"cleanarch/server/router"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app    *fiber.App
	config *configs.Configs
	db     *sql.DB
}

var (
	once   sync.Once
	server *Server
)

func NewServer(config *configs.Configs, db *sql.DB) *Server {
	once.Do(func() { //ทำเพื่อทำครั้งเดียวแล้วใช้ได้เลย
		app := fiber.New()
		server = &Server{
			app:    app,
			config: config,
			db:     db,
		}
	})
	return server
}

func (s *Server) Start() {
	api := s.app.Group("/api/v1")
	router.Product(api, s.db)
	router.Catagory(api, s.db)
	log.Printf("Sever start at port = %s", s.config.App.Port)
	if err := s.app.Listen(fmt.Sprintf(":%v", s.config.App.Port)); err != nil {
		log.Fatal("Fail to start Server ", err)
	}

}
