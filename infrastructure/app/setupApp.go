package app

import (
	"github.com/jalozanot/demoCeiba/application/usescases"
	"github.com/jalozanot/demoCeiba/domain/ports"
	"github.com/jalozanot/demoCeiba/domain/service"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/database_client"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/movie_rep"
)

func CreateHandler(userRepository ports.MoviesRepository) service.RedirectMovieHandler {

	return newHandler(newCreatesUseCase(userRepository), newGetMovieUseCase(userRepository),
		newUpdateMovieUseCase(userRepository), newDeleteMovieUseCase(userRepository), newGetMoviesUseCase(userRepository))
}
func newCreatesUseCase(repository ports.MoviesRepository) usescases.CreatesMoviePort {
	return &usescases.UseCaseMovieCreate{
		MovieRepository: repository,
	}
}

func newGetMovieUseCase(repository ports.MoviesRepository) usescases.GetMovieUseCase {
	return &usescases.UseCaseGetMovie{
		UserRepository: repository,
	}
}

func newUpdateMovieUseCase(repository ports.MoviesRepository) usescases.UpdateMovieUseCase {
	return &usescases.UseCaseUpdateMovie{
		UserRepository: repository,
	}
}

func newDeleteMovieUseCase(repository ports.MoviesRepository) usescases.DeleteMovieUseCase {
	return &usescases.UseCaseDeleteMovie{
		UserRepository: repository,
	}
}

func newGetMoviesUseCase(repository ports.MoviesRepository) usescases.GetMoviesUseCase {
	return &usescases.UseCaseGetMovies {
		UserRepository: repository,
	}
}

func newHandler(createMovie usescases.CreatesMoviePort, getMovieUseCase usescases.GetMovieUseCase, updateMovieUseCase usescases.UpdateMovieUseCase,
	deleteMovieUseCase usescases.DeleteMovieUseCase,getMoviesUseCase usescases.GetMoviesUseCase) service.RedirectMovieHandler {
	return &service.Handler{CreatesUseCase: createMovie, GetMovieUseCase: getMovieUseCase, UseCaseUpdateMovie: updateMovieUseCase,
		UseCaseDeleteMovie: deleteMovieUseCase, GetMoviesUseCase: getMoviesUseCase,
	}
}

func GetUsersRepository() ports.MoviesRepository {
	return &movie_rep.UserMysqlRepository{
		Db: database_client.GetDatabaseInstance(),
	}
}