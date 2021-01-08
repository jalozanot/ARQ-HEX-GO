package usescases

import (
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/domain/ports"
)

type GetMovieUseCase interface {
	Handler(userId int64) (model.Movie, error)
}

type UseCaseGetMovie struct {
	UserRepository ports.MoviesRepository
}

func (useCaseGetMovie *UseCaseGetMovie) Handler(id int64) (model.Movie, error) {

	movieDto, err := useCaseGetMovie.UserRepository.Get(id)
	if err != nil {
		return model.Movie{}, err
	}
	return movieDto, nil
}
