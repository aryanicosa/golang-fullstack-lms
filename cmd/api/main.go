package main

import (
	"github.com/aryanicosa/golang-fullstack-lms/internal/config"
	"github.com/aryanicosa/golang-fullstack-lms/internal/handlers"
	"github.com/aryanicosa/golang-fullstack-lms/internal/middleware"
	"github.com/aryanicosa/golang-fullstack-lms/internal/repository"
	"github.com/aryanicosa/golang-fullstack-lms/internal/repository/postgres"
	"github.com/aryanicosa/golang-fullstack-lms/internal/route"
	"github.com/aryanicosa/golang-fullstack-lms/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// init config
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// init database & migrations
	db, err := postgres.NewPostgresDB(config)
	if err != nil {
		panic(err)
	}
	defer postgres.ClosePostgresDB(db)

	// init dependency injection
	// jwt / auth service
	var jwtService service.JWTService = service.NewJWTService()

	// user related
	var userRepository repository.UserRepositoryInterface = repository.NewUserRepository(db)
	var userService service.UserServiceInterface = service.NewUserService(userRepository, jwtService)
	var userHandlers handlers.UserHandlersInterface = handlers.NewUserHandlers(userService)

	// init gin
	gin := gin.Default()
	gin.Use(middleware.CORSMiddleware())

	// init routes, service & utils
	routes.User(gin, userHandlers, jwtService)

	// init static asset
	gin.Static("/web", "./web")

	// run
	gin.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }