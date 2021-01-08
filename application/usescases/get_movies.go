package usescases

import (
	"fmt"
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/domain/ports"
)

type GetMoviesUseCase interface {
	Handler() ([]model.Movie, error)
}

type UseCaseGetMovies struct {
	UserRepository ports.MoviesRepository
}

func (useCaseGetMovies *UseCaseGetMovies) Handler() ([]model.Movie, error) {
	fmt.Println("linea numero 18")
	movieDto, err := useCaseGetMovies.UserRepository.Gets()
	fmt.Println("linea numero 20")
	if err != nil {
		return []model.Movie{}, err
	}
	return movieDto, nil
}
