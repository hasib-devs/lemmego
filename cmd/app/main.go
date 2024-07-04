package main

import (
	"log/slog"
	"pressebo/api"
	_ "pressebo/internal/config"
	"pressebo/internal/handlers"
	"pressebo/internal/plugins"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
)

func main() {
	// Print config
	slog.Info("App will start using the following config:\n", "config", api.ConfMap())

	// Load plugins
	registry := plugins.Load()

	// Create application
	app := api.NewApp(
		api.WithPlugins(registry),
	)
	logger := httplog.NewLogger("lemmego", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		TimeFieldFormat:  "[15:04:05.000]",
		// Tags: map[string]string{
		// 	"version": "v1.0-81aa4244d9fc8076a",
		// 	"env":     "dev",
		// },
		// QuietDownRoutes: []string{
		// 	"/",
		// 	"/ping",
		// },
		// QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})
	// Register global middleware
	app.Use(httplog.RequestLogger(logger), middleware.Recoverer)

	// Register routes
	handlers.Register(app)

	// Handle signals
	go app.HandleSignals()

	// Run application
	app.Run()
}
