package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/YEgorLu/time-tracker/internal/controllers"
)

type Server struct {
	http.Server
}

type ServerConfig struct {
	Port string
}

var defaultConfig = ServerConfig{
	Port: ":8080",
}

func NewServer(conf *ServerConfig) *Server {
	conf = parseConfig(conf)
	fmt.Println(conf)
	server := &Server{
		Server: http.Server{
			Addr: conf.Port,
		},
	}
	return server
}

func (s *Server) Configure() *Server {
	router, err := controllers.GetRoutes()
	if err != nil {
		panic(err)
	}
	s.Handler = router
	return s
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	println("server is listening on ", s.Addr)

	return s.Server.Serve(l)
}

func parseConfig(c *ServerConfig) *ServerConfig {
	conf := *c
	if conf.Port == "" {
		conf.Port = defaultConfig.Port
	} else if !strings.HasPrefix(conf.Port, ":") {
		conf.Port = ":" + conf.Port
	}
	return &conf
}
