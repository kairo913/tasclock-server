package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/infra"
)

type Config struct {
	Port         string `env:"PORT" envDefault:"8080"`
	IsProduction bool   `env:"PRODUCTION" envDefault:"false"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r, err := infra.SetUpRouter(ctx)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	sqlHandler, err := infra.NewSqlHandler()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	log.Println("Connected to database")

	sqlHandler.Close()

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Println("Starting server on port:", cfg.Port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown forced: %s\n", err)
	}

	log.Println("Server exiting")
}
