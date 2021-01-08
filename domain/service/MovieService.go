package service

import (
	"fmt"
	"github.com/fmcarrero/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/jalozanot/demoCeiba/application/commands"
	"github.com/jalozanot/demoCeiba/application/usescases"
	"github.com/jalozanot/demoCeiba/domain/ports"
	"github.com/jalozanot/demoCeiba/infrastructure/marshallers"
	"net/http"
	"strconv"
)

type RedirectMovieHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Gets(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}


type Handler struct {
	CreatesUseCase      usescases.CreatesMoviePort
	GetMovieUseCase     usescases.GetMovieUseCase
	GetMoviesUseCase    usescases.GetMoviesUseCase
	UseCaseUpdateMovie  usescases.UpdateMovieUseCase
	UseCaseDeleteMovie  usescases.DeleteMovieUseCase
}

var Repo *ports.MoviesRepository
var Han RedirectMovieHandler

func ProxyHandler(codigo int, c *gin.Context)  {

	if codigo == 1 {
		fmt.Println(" handler create ")
		Han.Create(c)
	} else if codigo == 2 {
		fmt.Println(" handler get ")
		Han.Get(c)
	} else if codigo == 3 {
		fmt.Println(" handler gets ")
		Han.Gets(c)
	} else if codigo == 4 {
		fmt.Println(" handler delete ")
		Han.Delete(c)
	} else if codigo == 5 {
		fmt.Println(" handler update ")
		Han.Update(c)
	}

}


func (h *Handler) Create(c *gin.Context) {

	var userCommand commands.MovieCommand
	if err := c.ShouldBindJSON(&userCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, createUserErr := h.CreatesUseCase.Handler(userCommand)

	if createUserErr != nil {
		_ = c.Error(createUserErr)
		return
	}

	fmt.Println("linea numero 48 ::: ", result)
	isPublic := true
	c.JSON(http.StatusCreated, marshallers.Marshall(isPublic, result))
}

func (h *Handler) Get(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	fmt.Println(userId)
	movie, errGet := h.GetMovieUseCase.Handler(userId)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}

	c.JSON(http.StatusOK, marshallers.Marshall(true, movie))
}

func (h *Handler) Gets(c *gin.Context) {

	fmt.Println("linea numero 74")
	movie, errGet := h.GetMoviesUseCase.Handler()

	if errGet != nil {
		_ = c.Error(errGet)
		return
	}

	c.JSON(http.StatusOK,  movie)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var movieCommand commands.MovieCommand
	if err := c.ShouldBindJSON(&movieCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	movieDto, updateErr := h.UseCaseUpdateMovie.Handler(id, movieCommand)
	if updateErr != nil {
		restErr := rest_errors.NewBadRequestError(updateErr.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, &movieDto)
}

func (h *Handler) Delete(c *gin.Context) {
	id, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	errDelete := h.UseCaseDeleteMovie.Handler(id)
	if errDelete != nil {
		restErr := rest_errors.NewBadRequestError(errDelete.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.Status(http.StatusNoContent)
}
