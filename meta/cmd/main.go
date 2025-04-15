package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/Xebec19/lms/common/pkg/logger"
	"github.com/Xebec19/lms/common/utils"
	"github.com/Xebec19/lms/meta/internal"

	"go.uber.org/zap"
)

func main() {

	logger.Init()
	logger.Log.Info("Setup logger")
	defer logger.Log.Sync()

	utils.InitConfig()

	server := internal.CreateServer()

	go func() {
		logger.Log.Info("Meta running", zap.String("PORT", utils.GetConfig().Port))

		err := server.ListenAndServe()

		if err != nil {
			logger.Log.Error("Server stopped", zap.Error(err))
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
