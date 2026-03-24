package repository

import "github.com/ShiunduZachariah/go-boilerplate-test/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
