package main

import (
	"tenet-profile/config"
	"tenet-profile/internal/client"
	"tenet-profile/internal/middleware"
	repository "tenet-profile/internal/repositories"
	service "tenet-profile/internal/services"
	"tenet-profile/internal/web/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	db, err := config.InitDataBase()
	if err != nil {
		panic("Failed to connect to database")
	}

	config.RunMigrations(db)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	_, err = dependenciesInit(router, db)
	if err != nil {
		panic("Failed to initialize dependencies")
	}

	router.Run(":8082")
}

func dependenciesInit(router *gin.Engine, db *gorm.DB) (*gin.Engine, error) {

	// Repositories
	profileRepo := repository.NewTenetProfileRepository(db)
	sessionAllowAttributesRepo := repository.NewSessionAllowAttributesRepository(db)

	// Services
	profileService := service.NewTenetProfileService(profileRepo, sessionAllowAttributesRepo)

	authClient := client.NewAuthClient("http://localhost:8080/auth")

	AuthMiddleware := middleware.NewAuthMiddleware(authClient)

	// Handlers
	profileHandler := handlers.NewProfileHandler(profileService)

	// Routes
	protected := router.Group("/").Use(AuthMiddleware.MiddlewareFunc())

	{
		protected.POST("/profile", profileHandler.CreateProfile)
		protected.GET("/profile/by-user/:userId", profileHandler.GetProfileByUserID)
		protected.GET("/profile-attributes/:sessionId/userId/:userId", profileHandler.GetAttributesFiltred)
	}

	return router, nil
}
