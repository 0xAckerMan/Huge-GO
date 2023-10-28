package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() *httprouter.Router{
    router := httprouter.New()
    router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheck)
    return router
}
