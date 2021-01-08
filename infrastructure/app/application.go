package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jalozanot/demoCeiba/domain/ports"
	"github.com/jalozanot/demoCeiba/domain/service"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/database_client"
	"github.com/jalozanot/demoCeiba/infrastructure/app/middlewares/error_handler"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

var UserRepository ports.MoviesRepository

//StartApplication inicio de aplicacion
func StartApplication() {

	_ = godotenv.Load()
	router.Use(error_handler.ErrorHandler())
	UserRepository = GetUsersRepository()
	service.Repo = &UserRepository
	service.Han = CreateHandler(UserRepository)
	r := MapUrls()
	database_client.GetConnectionRedis()

	/*r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))*/
	r.Use(cors.Default())
	_ = r.Run(":8084")
}


