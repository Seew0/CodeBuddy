package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/seew0/DoubtBuddy/handler"
	"github.com/seew0/DoubtBuddy/middleware"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func NewServer(port string, engine *gin.Engine) *Server {
	return &Server{
		port:   port,
		engine: engine,
	}
}

func (s *Server) Run() {
	gin.SetMode(gin.ReleaseMode)

	s.engine.Use(middleware.CORSmanager)

	s.engine.POST("/api/answer", func(ctx *gin.Context){handler.Getanswer(ctx)})
	
	err := s.engine.Run(s.port)
	if err != nil {
		log.Printf("Some Unexpected Error Occured: %v", err)
	}
	fmt.Println("Server is running at port:", s.port)
}
