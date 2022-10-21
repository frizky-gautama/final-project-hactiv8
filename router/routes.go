package router

import (
	"MyGram/config"
	"MyGram/controller"
	"MyGram/middleware"
	"MyGram/repo"
	"log"

	"github.com/gin-gonic/gin"
)

func Routes() {
	_, port_app, err := config.GetEnv()
	if err != nil {
		log.Fatal("Environment is not Connected", err)
		return
	}

	db, err := config.GetConnection()
	if err != nil {
		log.Fatal("Database not Connected", err)
		return
	}
	// =================================== HANDLER METHOD ===================================================
	// AUTHOR
	authorRepository := repo.NewAuthorRepository(db)
	authorHandler := middleware.NewAuthorController(authorRepository)
	// USER
	userRepository := repo.NewUserRepository(db)
	AuthHandler := controller.NewAuthController(userRepository)
	UserHandler := controller.NewUserController(userRepository)
	// PHOTO
	photoRepository := repo.NewPhotoRepository(db)
	PhotoHandler := controller.NewPhotoController(photoRepository)
	// Comment
	commentRepository := repo.NewCommentRepository(db)
	CommentHandler := controller.NewCommentController(commentRepository)
	// Socmed
	socmedRepository := repo.NewSocmedRepository(db)
	SocmedHandler := controller.NewSocmedController(socmedRepository)
	// =================================== ROUTES ===================================================
	router := gin.Default()
	// AUTH
	router.POST("/register", AuthHandler.Register)
	router.POST("/login", AuthHandler.Login)
	// MIDDLEWARE
	router.Use(middleware.Auth)
	// USER
	router.PUT("/users/:id", middleware.AuthorUser(), UserHandler.UpdateUser)
	router.DELETE("/users/:id", middleware.AuthorUser(), UserHandler.DeleteUser)
	// PHOTO
	router.POST("/photos", PhotoHandler.InsertPhoto)
	router.GET("/photos", PhotoHandler.GetAllFoto)
	router.PUT("/photos/:id", authorHandler.AuthorPhoto(), PhotoHandler.UpdatePhoto)
	router.DELETE("/photos/:id", authorHandler.AuthorPhoto(), PhotoHandler.DeletePhoto)
	// Comment
	router.POST("/comments", CommentHandler.InsertComment)
	router.GET("/comments", CommentHandler.GetAllComment)
	router.PUT("/comments/:id", authorHandler.AuthorComment(), CommentHandler.UpdateComment)
	router.DELETE("/comments/:id", authorHandler.AuthorComment(), CommentHandler.DeleteComment)
	// SOCMED
	router.POST("/socialmedias", SocmedHandler.InsertSocmed)
	router.GET("/socialmedias", SocmedHandler.GetAllSocmed)
	router.PUT("/socialmedias/:id", authorHandler.AuthorSocmed(), SocmedHandler.UpdateSocmed)
	router.DELETE("/socialmedias/:id", authorHandler.AuthorSocmed(), SocmedHandler.DeleteSocmed)
	// RUN
	router.Run(*port_app)
}
