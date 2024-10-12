package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"text-to-api/internal/handlers/translations"
	"text-to-api/internal/server"
	translationsService "text-to-api/internal/services/translations"
	"text-to-api/internal/translators/openai"
	"text-to-api/internal/zap"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func gracefulShutdown(fiberServer *server.FiberServer) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
}

func main() {

	logger, err := zap.NewLogger("development")
	if err != nil {
		panic(fmt.Sprintf("could not create logger: %s", err))
	}

	translator, err := openai.NewOpenAITranslator(logger, os.Getenv("OPENAI_APIKEY"), os.Getenv("OPENAI_ASSISTANT_ID"))
	if err != nil {
		panic(fmt.Sprintf("could not create translator: %s", err))
	}

	service, err := translationsService.NewTranslationsService(translator, logger)
	if err != nil {
		panic(fmt.Sprintf("could not create translations service: %s", err))
	}

	hdl, err := translations.NewTranslationsHandler(service)
	if err != nil {
		panic(fmt.Sprintf("could not create translations handler: %s", err))
	}

	srv := server.New()
	srv.App.Post("/v1/translations", hdl.Create)

	// Todo: potentially delete the following handlers
	srv.App.Get("/", srv.HelloWorldHandler)
	srv.App.Get("/health", srv.HealthHandler)

	go func() {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		err := srv.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err))
		}
	}()

	gracefulShutdown(srv)
}
