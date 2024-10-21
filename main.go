package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/util/config"
	"github.com/kairo913/tasclock-server/app/web/handler"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg := config.NewServerConfig(ctx)

	if cfg.ProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r, err := handler.SetUpRouter(ctx)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Println("Starting server on port:", cfg.Port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("server shutdown forced: %s\n", err))
	}

	log.Println("Server exiting")
}
