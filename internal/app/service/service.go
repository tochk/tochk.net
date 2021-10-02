package service

import (
	"tochkru-golang/internal/app/config"
	"tochkru-golang/internal/app/repository"
)

type Service struct {
	r   *repository.Repository
	cnf config.Config
}

func New(cnf config.Config, r *repository.Repository) *Service {
	return &Service{
		r:   r,
		cnf: cnf,
	}
}

func (s *Service) Ready() error {
	return s.r.Ready()
}
