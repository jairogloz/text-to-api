package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers/middleware"
	"text-to-api/internal/handlers/request_context"
	"text-to-api/internal/handlers/stripe"
	"text-to-api/internal/handlers/translations"
	"text-to-api/internal/repositories/mongo"
	"text-to-api/internal/repositories/mongo/user"
	"text-to-api/internal/repositories/postgres"
	"text-to-api/internal/repositories/postgres/client"
	"text-to-api/internal/server"
	"text-to-api/internal/services/auth"
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

	pgxPool, disconnectFunc, err := postgres.Connect(context.Background(), os.Getenv("POSTGRES_URI"))
	if err != nil {
		panic(fmt.Sprintf("could not connect to postgres: %s", err))
	}
	defer disconnectFunc()

	pgxClientRepo, err := client.NewClientRepository(pgxPool)
	if err != nil {
		panic(fmt.Sprintf("could not create client repository: %s", err))
	}

	mongoClient, disconnect, err := mongo.ConnectMongoDB(os.Getenv("MONGO_URI"))
	if err != nil {
		panic(fmt.Sprintf("could not connect to mongodb: %s", err))
	}
	defer disconnect()

	translator, err := openai.NewOpenAITranslator(logger, os.Getenv("OPENAI_APIKEY"), os.Getenv("OPENAI_OBJECT_TRANSLATOR_ASSISTANT_ID"))
	if err != nil {
		panic(fmt.Sprintf("could not create translator: %s", err))
	}

	userRepo, err := user.NewUserRepository(mongoClient, logger, os.Getenv("MONGO_DB_NAME"))
	if err != nil {
		panic(fmt.Sprintf("could not create user repository: %s", err))
	}

	service, err := translationsService.NewTranslationsService(translator, logger, userRepo)
	if err != nil {
		panic(fmt.Sprintf("could not create translations service: %s", err))
	}

	hdl, err := translations.NewTranslationsHandler(service, logger)
	if err != nil {
		panic(fmt.Sprintf("could not create translations handler: %s", err))
	}

	stripeHdl, err := stripe.NewStripeHandler(os.Getenv("STRIPE_API_KEY"),
		os.Getenv("STRIPE_SUCCESS_URL"),
		os.Getenv("STRIPE_CANCEL_URL"))
	if err != nil {
		panic(fmt.Sprintf("could not create stripe handler: %s", err))
	}

	srv := server.New()

	jwtSecretKey := os.Getenv("SUPABASE_JWT_SECRET")
	if jwtSecretKey == "" {
		panic("SUPABASE_JWT_SECRET is required")
	}
	authSrv, err := auth.NewAuthService(pgxClientRepo, []byte(jwtSecretKey), logger)
	if err != nil {
		panic(fmt.Sprintf("could not create auth service: %s", err))
	}

	// Register the CORS middleware globally
	srv.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	reqCtxHdl := request_context.NewRequestContextHandler()
	authMdlw, err := middleware.NewAuthMdlwHdl(authSrv, logger, reqCtxHdl)
	if err != nil {
		panic(fmt.Sprintf("could not create auth middleware handler: %s", err))
	}

	headersMdlw, err := middleware.NewHeadersMdlwHdl(logger, reqCtxHdl)
	if err != nil {
		panic(fmt.Sprintf("could not create headers middleware handler: %s", err))
	}

	apiKeyAuthGroup := srv.App.Group("/v1/translations", authMdlw.Auth(domain.AuthTypeAPIKey), headersMdlw.ForceHeaders([]string{"User-Id"}))

	tokenAuthGroup := srv.App.Group("/v1/checkout-session", authMdlw.Auth(domain.AuthTypeToken), headersMdlw.ForceHeaders([]string{"Environment"}))

	// Register the translations handler
	apiKeyAuthGroup.Post("/", hdl.Create)

	// Register the Stripe checkout-session handler
	tokenAuthGroup.Post("/", stripeHdl.CreateCheckoutSession)

	// Todo: potentially delete the following handlers
	srv.App.Get("/", srv.HelloWorldHandler)
	srv.App.Get("/health", srv.HealthHandler)

	go func() {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			logger.Fatal(context.Background(), "could not parse port", "error", err.Error())
		}
		fmt.Println("Listening on port", port)
		err = srv.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err))
		}
	}()

	gracefulShutdown(srv)
}
