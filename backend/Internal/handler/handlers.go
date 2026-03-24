package handler

import (
	"github.com/ShiunduZachariah/go-boilerplate-test/internal/server"
	"github.com/ShiunduZachariah/go-boilerplate-test/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
