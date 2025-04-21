package httpserver

import (
	"medtest/config"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Server *gin.Engine
	Port   string
}

func NewServer(Config *config.Config) *HttpServer {
	r := gin.Default()
	return &HttpServer{Server: r, Port: Config.Server.Port}
}

func (S *HttpServer) MapGet(path string, Handler func(c *gin.Context)) {
	S.Server.GET(path, Handler)
}

func (S *HttpServer) MapPost(path string, Handler func(c *gin.Context)) {
	S.Server.POST(path, Handler)
}

func (S *HttpServer) Run() {
	S.Server.Run(S.Port)
}
