package main

import (
    "log"

    "github.com/example/gin-rest-api/internal/config"
    "github.com/example/gin-rest-api/internal/handler"
    "github.com/example/gin-rest-api/internal/middleware"
    "github.com/example/gin-rest-api/internal/repository"
    "github.com/example/gin-rest-api/internal/route"
    "github.com/example/gin-rest-api/internal/service"

    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.Load()

    if cfg.Env == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(middleware.CORS())

    userRepo := repository.NewUserRepository()
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    route.Register(r, userHandler)

    addr := ":" + cfg.Port
    log.Printf("%s listening on %s", cfg.AppName, addr)

    if err := r.Run(addr); err != nil {
        log.Fatal(err)
    }
}
