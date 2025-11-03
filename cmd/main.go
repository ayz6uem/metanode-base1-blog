package main

import (
	"base1-blog/config"
	"base1-blog/models"
	"base1-blog/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Warn(".env file not found, relying on environment variables")
	}
	logrus.SetFormatter(&logrus.TextFormatter{})

	logrus.Info("Initializing database...")
	config.InitDatabase()
	models.InitModels()

	r := routes.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	logrus.Info("Starting application... on port ", port)
	if err := r.Run(port); err != nil {
		logrus.Fatal(err)
		return
	}
}
