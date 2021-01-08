package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jalozanot/demoCeiba/infrastructure/controllers"
)

func MapUrls(/*handler controllers.RedirectMovieHandler*/) *gin.Engine {

	ping := router.Group("/")
	{
		ping.GET("/ping", controllers.Ping)
	}

	movie := router.Group("/")
	{
		movie.POST("/peliculas", controllers.Create)
		movie.GET("/peliculas/:id", controllers.Get)
		movie.GET("/peliculas", controllers.Gets)
		movie.PUT("/peliculas/:id", controllers.Update)
		movie.DELETE("/peliculas/:id", controllers.Delete)
	}

	//cliente := ...

	return router
}
