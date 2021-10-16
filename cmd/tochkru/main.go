package main

import (
	"flag"
	"net/http"
	"tochkru-golang/internal/app/api"
	"tochkru-golang/internal/app/config"
	"tochkru-golang/internal/app/metrics"
	"tochkru-golang/internal/app/repository"
	"tochkru-golang/internal/app/service"
	"tochkru-golang/internal/app/web"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	logLevelFlag = flag.String("log", "info", "log level")
)

func main() {
	log.Info("starting tochk.ru web server")
	flag.Parse()
	cfg := config.Config{}
	envconfig.MustProcess("tochkru", &cfg)

	logLevel, err := log.ParseLevel(*logLevelFlag)
	if err != nil {
		log.Warningln("can't parse provided log level, using info level")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
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
		log.Info("listening internal on: " + cfg.InternalListenAddr)
		if err := http.ListenAndServe(cfg.InternalListenAddr, internalRouter); err != nil {
			panic(err)
		}
	}()

	box := rice.MustFindBox("../../internal/static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.Handle("/static/{path:.*}", staticFileServer)

	router.HandleFunc("/", w.Wrapper(w.IndexPage)).Methods("GET")

	router.HandleFunc("/project/{id}", w.Wrapper(w.ProjectPage)).Methods("GET")

	log.Info("listening on: " + cfg.ListenAddr)
	if err := http.ListenAndServe(cfg.ListenAddr, router); err != nil {
		panic(err)
	}
}
