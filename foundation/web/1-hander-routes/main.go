package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "1.0.0"

type Config struct {
    Port int
    Env string
}

type Application struct {
    Config
    logger *log.Logger
}

func main() {

    var cfg Config

    flag.IntVar(&cfg.Port, "port", 4000, "The API server port")
    flag.StringVar(&cfg.Env, "env", "dev", "The environment(dev|staging|prod)")
    flag.Parse()

    addr := fmt.Sprintf(":%d", cfg.Port)


    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

    app := &Application{
        Config: cfg,
        logger: logger,
    }

    srv := &http.Server{
        Addr: addr,
        Handler: app.route(),
        IdleTimeout: time.Minute,
        ReadTimeout: 10*time.Second,
        WriteTimeout: 30*time.Second,
    }

    logger.Printf("starting %s server on %s\n", cfg.Env, addr)
    err := srv.ListenAndServe()
    logger.Fatal(err)

}


