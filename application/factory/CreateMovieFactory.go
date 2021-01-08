package factory

import (
	"github.com/jalozanot/demoCeiba/application/commands"
	"github.com/jalozanot/demoCeiba/domain/model"
)

func CreateMovie(movieCommand commands.MovieCommand) (model.Movie, error) {
	var movie model.Movie
	movieDto, err := movie.CreateMovil(movieCommand.Nombre, movieCommand.Categoria, movieCommand.CodigoBarras)
	return movieDto, err
}
