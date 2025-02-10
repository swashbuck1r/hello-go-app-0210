package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func start(port int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Handle ctrl-C signals
	var SignalChan = make(chan os.Signal, 1)
	signal.Notify(SignalChan, os.Interrupt)
	defer func() {
		signal.Stop(SignalChan)
		cancel()
	}()
	go func() {
		select {
		case <-SignalChan: // first signal, cancel context
			log.Printf("Received an interrupt, stopping services...")
			cancel()
		case <-ctx.Done():
		}
		os.Exit(2)
	}()

	r := setupRouter(ctx)
	log.Printf("Server running on port %d", port)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal("could not run server", err)
	}
}

func setupRouter(ctx context.Context) *gin.Engine {
	ginEngine := gin.New()

	ginEngine.SetTrustedProxies(nil)
	ginEngine.Use(gin.Recovery())

	ginEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world3")
	})

	return ginEngine
}

func main() {
	start(8080)
}
