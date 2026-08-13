package server

import (
	configs "cleanarch/config"
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
	app := fiber.New()
	once.Do(func() { //ทำเพื่อทำครั้งเดียวแล้วใช้ได้เลย
		server = &Server{
			app:    app,
			config: config,
			db:     db,
		}
		fmt.Println("LUK PEE PALM")
	})
	fmt.Println("LUK PEE PALM2")
	return server
}

func (s *Server) Start() {
	api := s.app.Group("/api/v1")
	api.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("P PALM HANSOME MAK MAK")
	})
	log.Printf("Sever start at port = %s", s.config.App.Port)
	if err := s.app.Listen(fmt.Sprintf(":%v", s.config.App.Port)); err != nil {
		log.Fatal("Fail to start Server ", err)
	}

}
