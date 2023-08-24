package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
    port int
    env string
}

type Aplication struct {
    config
    logger *log.Logger
}

func main(){

    var cfg config

    flag.IntVar(&cfg.port, "port", 3000, "API running port")
    flag.StringVar(&cfg.env, "env", "dev", "Current Environment(dev|staging|prod)")
    flag.Parse()

    addr := fmt.Sprintf(":%d", cfg.port)

    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    app := &Aplication{
        config: cfg,
        logger: logger,
    }

    srv := http.Server{
        Addr: addr,
        Handler: app.routes(),
        IdleTimeout: time.Minute,
        WriteTimeout: 30*time.Second,
        ReadTimeout: 10*time.Second,
    }

    logger.Printf("Server is running on %v env on port %v\n",app.env,addr)
    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
