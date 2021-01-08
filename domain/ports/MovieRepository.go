package ports

import "github.com/jalozanot/demoCeiba/domain/model"

type MoviesRepository interface {
	Save(user *model.Movie) (model.Movie, error)
	Get(userId int64) (model.Movie, error)
	Gets() ([]model.Movie, error)
	Update(userId int64, user model.Movie) (*model.Movie, error)
	Delete(userId int64) error
}
