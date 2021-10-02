package main

import (
	"net/http"
	"tochkru-golang/internal/app/api"
	"tochkru-golang/internal/app/config"
	"tochkru-golang/internal/app/metrics"
	"tochkru-golang/internal/app/repository"
	"tochkru-golang/internal/app/service"
	"tochkru-golang/internal/app/web"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting")
	cfg := config.Config{}
	envconfig.MustProcess("tochkru", &cfg)

	db := sqlx.MustConnect("postgres", cfg.DbUrl)
	log.Info("connected to database")

	repo := repository.New(db)
	s := service.New(cfg, repo)
	m := metrics.New()
	w := web.New(s, m)

	router := mux.NewRouter()

	go func() {
		internalRouter := mux.NewRouter()
		a := api.New(s, m)
		internalRouter.Handle("/metrics", promhttp.Handler()).Methods("GET")
		internalRouter.HandleFunc("/ready", a.GetReadyHandler).Methods("GET")
		log.Info("Listening internal on: " + cfg.InternalListenAddr)
		if err := http.ListenAndServe(cfg.InternalListenAddr, internalRouter); err != nil {
			panic(err)
		}
	}()

	router.HandleFunc("/", w.Wrapper(w.IndexPage)).Methods("GET")

	log.Info("Listening on: " + cfg.ListenAddr)
	if err := http.ListenAndServe(cfg.ListenAddr, router); err != nil {
		panic(err)
	}
}
