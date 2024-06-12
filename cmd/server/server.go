package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start the map to hold schedulers
	// Start the API handler
	// Start the ticker

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool { return true } // allow all origins for our exercise
	router.Use(cors.New(config))

	g := router.Group("/api/v1/")
	addRoutes(g)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		err := srv.ListenAndServe()
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	})
	eg.Go(func() error {
		<-egCtx.Done() // This context is done when the base context (ctx) is called.
		timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := srv.Shutdown(timeout)
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	})

	if err := eg.Wait(); err != nil {
		fmt.Printf("exit reason: %s\n", err.Error())
	}

}

func addRoutes(g *gin.RouterGroup) {
	g.GET("/ping", handlePing)
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
