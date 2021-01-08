package users_mapper

import (
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SetupMovie() model.Movie{

	movie := model.Movie{Nombre:"Spider Man", Categoria: "Ficcion", CodigoBarras: "1"}
	return movie
}

func SetupMovieEntity() models.MovieEntity {

	return models.MovieEntity{Nombre: "super-man", Categoria: "ficcion", CodigoBarras: "2"}

}

func TestMovieDtoToMovieEntity(t *testing.T) {

	movieEntity := MovieDtoToMovieEntity(SetupMovie())

	assert.Equal(t,movieEntity.Nombre,"Spider Man")

}

func TestMovieEntityToMovieDto(t *testing.T) {

	movie := MovieEntityToMovieDto(SetupMovieEntity())
	assert.Equal(t, movie.Nombre,"super-man")
	assert.Equal(t, movie.Categoria,"ficcion")
}

func TestMovieEntityToMovilesDto(t *testing.T) {
	movies := make([]models.MovieEntity, 5)
	movies[0] = SetupMovieEntity()
	movie := MovieEntityToMovilesDto(movies)

	assert.Equal(t,movie[0].Nombre,"super-man")

}