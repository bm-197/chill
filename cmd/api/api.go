package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type application struct {
    config config
}

type config struct {
    addr string
}

func (app *application) mount() http.Handler {
    r := chi.NewMux()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.RealIP)
    r.Use(middleware.RequestID)
    r.Use(middleware.Timeout(time.Second * 60))

    r.Route("/v1", func(r chi.Router) {
        r.Get("/health", app.healthCheckHandler)
    })
    return r
}

func (app *application) run(mux http.Handler) error {
    srv := &http.Server{
        Addr: app.config.addr,
        Handler: mux,
        WriteTimeout: 10 * time.Second,
        ReadTimeout: 3 * time.Second,
        IdleTimeout: 5 * time.Second,
    }

    log.Printf("Server is running on port %s", srv.Addr)
    return srv.ListenAndServe()
}
