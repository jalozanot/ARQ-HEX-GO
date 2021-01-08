package users_mapper

import (
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/models"
)

func MovieDtoToMovieEntity(movieDto model.Movie) models.MovieEntity {

	var movieEntity models.MovieEntity
	movieEntity.Nombre = movieDto.Nombre
	movieEntity.Categoria = movieDto.Categoria
	movieEntity.CodigoBarras = movieDto.CodigoBarras

	return movieEntity
}

func MovieEntityToMovieDto(movieEntity models.MovieEntity) model.Movie {
	var Movie model.Movie
	Movie.Id = movieEntity.ID
	Movie.Nombre = movieEntity.Nombre
	Movie.Categoria = movieEntity.Categoria
	Movie.CodigoBarras = movieEntity.CodigoBarras
	return Movie
}

func MovieEntityToMovilesDto(MovieEntity []models.MovieEntity) []model.Movie {
	var movies []model.Movie
	for _, MovieEntity := range MovieEntity {
		movie := MovieEntityToMovieDto(MovieEntity)
		movies = append(movies, movie)
	}
	return movies
}
