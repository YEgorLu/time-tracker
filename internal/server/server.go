package server

import (
	"net"
	"net/http"
	"strings"

	"github.com/YEgorLu/time-tracker/internal/controllers"
	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	http.Server
	router *http.ServeMux
	log    logger.Logger
}

type ServerConfig struct {
	Port string
	Log  logger.Logger
}

var defaultConfig = ServerConfig{
	Port: ":8080",
}

func NewServer(conf *ServerConfig) *Server {
	conf = parseConfig(conf)
	server := &Server{
		Server: http.Server{
			Addr: conf.Port,
		},
		log: conf.Log,
	}
	return server
}

func (s *Server) Configure() *Server {
	router, err := controllers.GetRoutes()
	if err != nil {
		s.log.Error(err)
		panic(err)
	}
	s.router = router
	return s
}

func (s *Server) WithSwagger() *Server {
	if s.router == nil {
		s.log.Error("WithSwagger must be placed after Configure()")
		panic("WithSwagger must be placed after Configure()")
	}
	s.router.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost"+s.Addr+"/swagger/doc.json")))
	return s
}

func (s *Server) Run() error {
	s.Handler = middleware.Logger(nil)(s.router)
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		s.log.Error(err)
		return err
	}

	s.log.Info("server is listening on ", s.Addr)

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
