package main

import (
	"fmt"
	"log"
	"net/http"

	"project-app-bioskop-golang-azwin/internal/wire"
	"project-app-bioskop-golang-azwin/pkg/database"
	"project-app-bioskop-golang-azwin/pkg/utils"

	"go.uber.org/zap"
)

func main() {
// Read configuration
config, err := utils.ReadConfiguration()
if err != nil {
log.Fatalf("failed to read file config: %v", err)
}

// Initialize logger
logger, err := utils.InitLogger(config.PathLogging, config.Debug)
if err != nil {
log.Fatalf("failed to initialize logger: %v", err)
}
defer logger.Sync()

// Initialize database
db, err := database.InitDB(config.DB)
if err != nil {
		logger.Fatal("failed to connect to postgres database", zap.Error(err))
	}
	logger.Info("Successfully connected to the database")

// Initialize app with dependency injection
app := wire.InitializeApp(db, logger, config)

// Start server
port := config.Port
if port == "" {
port = "8080"
}

address := fmt.Sprintf(":%s", port)
logger.Info(fmt.Sprintf("Starting server on %s", address))

if err := http.ListenAndServe(address, app.Router); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
}
}