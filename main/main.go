package main

import (
    "context"
    "github.com/gorilla/mux"
    "github.com/joshinjohnson/authentication-engine/pkg/api"
    "github.com/joshinjohnson/authentication-engine/pkg/models"
    "github.com/joshinjohnson/authentication-manager/internal"
    "github.com/joshinjohnson/authentication-manager/pkg/errors"
    "github.com/joshinjohnson/authentication-manager/tokenengine"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
)

const hostAddress = "127.0.0.1:9090"

func main() {
    mode := models.Mode(2)
    engine, err := api.New(models.Config{Mode: mode})
    if err != nil {
        log.Println(errors.ErrInternalServer)
        os.Exit(1)
    }

    tokenManager := tokenengine.TokenGeneratorEngine{}
    manager := internal.Manager{AuthenticationEngine: engine, TokenGeneratorEngine: tokenManager}
    r := mux.NewRouter()
    r.Use(manager.VerifyTokenMiddleware)
    r.HandleFunc("/login", manager.LoginHandler)
    r.HandleFunc("/register", manager.RegisterHandler)
    r.HandleFunc("/home", manager.HomeHandler)

    server := &http.Server{
        Addr:    hostAddress,
        Handler: r,
    }

    go func() {
        if err := server.ListenAndServe(); err != nil {
            log.Println(errors.ErrInternalServer)
            os.Exit(1)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c

    ctx, cancel := context.WithTimeout(context.Background(), time.Minute*15)
    defer cancel()
    server.Shutdown(ctx)
    log.Println("shutting down server")
    os.Exit(0)
}
