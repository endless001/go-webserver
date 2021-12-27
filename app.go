package main

import (
	"log"
	"github.com/justinas/alice"
	"github.com/gorilla/mux"
	"net/http"
	"webserver/middleware"
)

type App struct {
	Router      *mux.Router
	Middlewares *middleware.Middleware
	Config      *Env
}


func (a *App) Initialize(e *Env) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Config = e
	a.Router = mux.NewRouter()
	a.Middlewares = &middleware.Middleware{}

	a.InitializeRoter()
}

//InitializeRoter ...
func (a *App) InitializeRoter() {
	m := alice.New(a.Middlewares.LoggingHandler)

	a.Router.Handle("/version", m.ThenFunc(a.version)).Methods("get")
	a.Router.Handle("/healthz", m.ThenFunc(a.healthz)).Methods("get")
}

//Run ...
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}


