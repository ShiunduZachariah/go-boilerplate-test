package service

import (
	"github.com/ShiunduZachariah/go-boilerplate-test/internal/lib/job"
	"github.com/ShiunduZachariah/go-boilerplate-test/internal/repository"
	"github.com/ShiunduZachariah/go-boilerplate-test/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
