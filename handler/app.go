package handler

import (
	"project2/docs"
	"project2/infra/config"
	"project2/infra/database"
	"project2/repository/comments_repository/comments_pg"
	"project2/repository/photo_repository/photo_pg"
	"project2/repository/socialmedia_repository/socialmedia_pg"
	"project2/repository/user_repository/user_pg"
	"project2/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	config.LoadAppConfig()

	database.InitiliazeDatabase()

	var port = config.GetAppConfig().Port

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	photoRepo := photo_pg.NewPhotoPG(db)

	photoService := service.NewPhotoService(photoRepo)

	photoHandler := NewPhotoHandler(photoService)

	commentRepo := comments_pg.NewCommentPG(db)

	commentService := service.NewCommentService(commentRepo)

	commentHandler := NewCommentHandler(commentService)

	socialRepo := socialmedia_pg.NewSocialMediaPG(db)

	socialService := service.NewSocialMediaService(socialRepo)

	socialHandler := NewSocialMediaHandler(socialService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialRepo)

	route := gin.Default()

	docs.SwaggerInfo.Title = "Project 2"
	docs.SwaggerInfo.Description = "Ini adalah Project ke 2 dari kelas Kampus Merdeka"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "project2-production-8bf5.up.railway.app"
	docs.SwaggerInfo.Schemes = []string{"https"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.Use(authService.Authentitaction())
		userRoute.PUT("/", userHandler.UpdateUser)
		userRoute.DELETE("/", userHandler.Delete)
	}

	photoRoute := route.Group("/photos")
	{
		photoRoute.Use(authService.Authentitaction())
		photoRoute.POST("/", photoHandler.Posting)
		photoRoute.GET("/", photoHandler.GetPhoto)
		photoRoute.PUT("/:photoId", authService.AuthorizationPhoto(), photoHandler.UpdatePhoto)
		photoRoute.DELETE("/:photoId", authService.AuthorizationPhoto(), photoHandler.DeletePhoto)
	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.Use(authService.Authentitaction())
		commentRoute.POST("/", commentHandler.MakeComment)
		commentRoute.GET("/", commentHandler.GetComments)
		commentRoute.PUT("/:commentId", authService.AuthorizationComment(), commentHandler.UpdateComment)
		commentRoute.DELETE("/:commentId", authService.AuthorizationComment(), commentHandler.DeleteComment)
	}

	socialMediaRoute := route.Group("/socialmedias")
	{
		socialMediaRoute.Use(authService.Authentitaction())
		socialMediaRoute.POST("/", socialHandler.PostSocialMedia)
		socialMediaRoute.GET("/", socialHandler.GetSocialMedia)
		socialMediaRoute.PUT("/:socialMediaId", authService.AuthorizationSocialMedia(), socialHandler.UpdateSocialMedia)
		socialMediaRoute.DELETE("/:socialMediaId", authService.AuthorizationSocialMedia(), socialHandler.DeleteSocialMedia)
	}

	route.Run(":" + port)
}
