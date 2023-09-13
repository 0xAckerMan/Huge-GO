package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const Version = "1.0.0"

type config struct{
    port int
    env string
}

type Application struct{
    config
    logger *log.Logger
}

func main (){
    var cfg config

    flag.IntVar(&cfg.port, "port", 4000, "Runing port server")
    flag.StringVar(&cfg.env, "env", "dev", "Running environment(dev|staging|prod)")
    flag.Parse()

    addr := fmt.Sprintf(":%d", cfg.port)

    logg := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    logg.Printf("Starting %s server on port %s\n", cfg.env, addr)
    app := Application{
        config: cfg,
        logger: logg,
    }
 
    srv := http.Server{
        Addr: addr,
        Handler: app.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 20*time.Second,
        WriteTimeout: 20*time.Second,
    }
    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
