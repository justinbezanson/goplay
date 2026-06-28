package main

import (
	"fmt"
	"log/slog"
	"os"

	"goplay/models"
	"goplay/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.InitDB("data")
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	r := gin.Default()

	routes.RegisterWebRoutes(r)
	routes.RegisterAPIRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info(fmt.Sprintf("server starting on :%s", port))
	r.Run(":" + port)
}
