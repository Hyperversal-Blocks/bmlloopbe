package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/bmlloopbe/pkg/adopter"
	"github.com/hyperversal-blocks/bmlloopbe/pkg/auth"
	jwtPkg "github.com/hyperversal-blocks/bmlloopbe/pkg/jwt"
	"github.com/hyperversal-blocks/bmlloopbe/pkg/node"
)

type Services struct {
	logger     *logrus.Logger
	router     *chi.Mux
	auth       auth.Auth
	user       adopter.Service
	node       *node.Node
	jwtService jwtPkg.JWT
}

func New(logger *logrus.Logger, router *chi.Mux, auth auth.Auth, user adopter.Service, node *node.Node, jwtService jwtPkg.JWT) *Services {
	return &Services{
		logger:     logger,
		router:     router,
		auth:       auth,
		user:       user,
		node:       node,
		jwtService: jwtService,
	}
}

func (s *Services) Cors() {
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
}

func (s *Services) Routes() {
	// TODO: setup routes
}

func (s *Services) GetRouter() *chi.Mux {
	return s.router
}
