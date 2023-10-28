package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

var Version = "1.0.0"

type Config struct{
    env string
    port int
}

type Application struct{
    config Config
    logger *log.Logger
}

func main(){
    var cfg Config
    flag.IntVar(&cfg.port, "port", 4000, "API server port")
    flag.StringVar(&cfg.env, "env", "dev", "Server environment(dev|stag|prod)")
    flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

    app := &Application{
        config: cfg,
        logger: logger,
    }

    addr := fmt.Sprintf(":%d", cfg.port)

    mux := httprouter.New()
    mux.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheck)

    srv := &http.Server{
        Addr: addr,
        Handler: app.routes(),
        WriteTimeout: 1 * time.Minute,
        ReadTimeout: 30 * time.Second,
        IdleTimeout: 30 * time.Second,
    }

    logger.Printf("Starting %s at port %d", cfg.env, cfg.port)

    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}


