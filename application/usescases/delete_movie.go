package usescases

import (
	"github.com/jalozanot/demoCeiba/domain/ports"
)

type DeleteMovieUseCase interface {
	Handler(id int64) error
}

type UseCaseDeleteMovie struct {
	UserRepository ports.MoviesRepository
}

func (useCaseDeleteMovie *UseCaseDeleteMovie) Handler(id int64) error {
	movie, getUserError := useCaseDeleteMovie.UserRepository.Get(id)
	if getUserError != nil {
		return getUserError
	}
	err := useCaseDeleteMovie.UserRepository.Delete(movie.Id)
	return err
}
