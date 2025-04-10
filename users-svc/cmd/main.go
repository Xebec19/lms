package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/Xebec19/lms/users-svc/internal"
	"github.com/Xebec19/lms/users-svc/internal/logger"
	"github.com/Xebec19/lms/users-svc/internal/utils"
	"go.uber.org/zap"
)

func main() {

	logger.Init()
	logger.Log.Info("Setup logger")
	defer logger.Log.Sync()

	utils.InitConfig()

	server := internal.CreateServer()

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Log.Error("Failed to start server", zap.Error(err))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Log.Info("shutting down")
	os.Exit(0)
}
