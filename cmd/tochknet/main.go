package main

import (
	"embed"
	"flag"
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

//go:embed static/*
var static embed.FS

var (
	logLevelFlag = flag.String("log", "info", "log level")
	cache        = flag.Bool("cache", false, "page cache")
)

func main() {
	log.Info("starting tochk.net web server")
	flag.Parse()
	cfg := config.Config{}
	envconfig.MustProcess("tochknet", &cfg)

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
	w := web.New(s, m, *cache)

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

	router.Handle("/static/{path:.*}", http.FileServer(http.FS(static)))

	router.HandleFunc("/", w.Wrapper(w.IndexPageHandler)).Methods("GET")

	router.HandleFunc("/articles/", w.Wrapper(w.ArticlesPageHandler)).Methods("GET")
	router.HandleFunc("/article/{id}", w.Wrapper(w.ArticlePageHandler)).Methods("GET")

	router.HandleFunc("/projects/", w.Wrapper(w.ProjectsPageHandler)).Methods("GET")
	router.HandleFunc("/project/{id}", w.Wrapper(w.ProjectPageHandler)).Methods("GET")

	log.Info("listening on: " + cfg.ListenAddr)
	if err := http.ListenAndServe(cfg.ListenAddr, router); err != nil {
		panic(err)
	}
}
