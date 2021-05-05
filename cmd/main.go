package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Manuel9550/DungeonFetcher/pkg/dal"
	"github.com/Manuel9550/DungeonFetcher/pkg/dungeonFetcher"
	"github.com/Manuel9550/DungeonFetcher/pkg/environment"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

const logPath = "./../log"
const logFilePath = "./../log/DungeonFetcher.log"

func main() {

	// Setting up the logger
	logger := log.NewLogfmtLogger(os.Stderr)

	// If Logging folder doesn't exist, create it
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, os.ModeDir)
	}

	// Setting up the log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		level.Error(logger).Log("exit", err)
	}

	defer logFile.Close()

	// Want to write to terminal and file, is possible
	mw := io.MultiWriter(os.Stdout, logFile)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(mw))

	// Get the environment variable we need
	env, ok := environment.GetEnvironmentVariables(logger)

	if !ok {
		os.Exit(-1)
	}

	// Create the database connection
	ctx := context.Background()

	// Create the service
	dataManager, err := dal.CreateDBManager(env.ConnectionString, logger, env.DBType)

	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	service := dungeonFetcher.NewService(logger, dataManager)

	// Make the endpoints
	endpoints := dungeonFetcher.MakeEndpoints(&service)

	// Run server in a goroutine so that it doesn't block
	go func() {
		handler := dungeonFetcher.NewHTTPServer(ctx, endpoints, logger)
		http.ListenAndServe(env.Address, handler)
	}()

	level.Info(logger).Log("msg", "DungeonFetcher service started")

	// Shut down on SIGINT

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c

	level.Info(logger).Log("msg", "DungeonFetcher service shutting down")
	os.Exit(0)

}
