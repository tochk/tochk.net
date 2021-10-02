package api

import (
	"tochkru-golang/internal/app/metrics"
	"tochkru-golang/internal/app/service"
)

type Api struct {
	s *service.Service
	m *metrics.Metrics
}

func New(s *service.Service, m *metrics.Metrics) *Api {
	return &Api{
		s: s,
		m: m,
	}
}
