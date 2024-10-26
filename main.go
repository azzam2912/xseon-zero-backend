package main

import (
	"fmt"
	"log"
	"xseon-zero/controller"
	"xseon-zero/handler"
	configlib "xseon-zero/lib/config"
	"xseon-zero/lib/db"
	"xseon-zero/repository/files_db"
	"xseon-zero/usecase/filelink"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting File Links Service")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config, err := configlib.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := db.ConnectToDB(config.DB)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	filesDB := files_db.NewFilesDBImpl(db)
	fileLinkUseCase := filelink.NewFileLinkImpl(filesDB)
	fileLinkHandler := handler.NewFileLinkHandler(fileLinkUseCase)
	ctrl := controller.NewController(fileLinkHandler)
	r := gin.Default()

	ctrl.SetupRoutes(r)
	serverAddr := fmt.Sprintf(":%s", config.Service.Port)
	log.Printf("Server starting on port %s", config.Service.Port)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
