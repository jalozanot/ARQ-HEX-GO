package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/domain/ports"
	"github.com/jalozanot/demoCeiba/domain/service"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/models"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/movie_rep"
	"github.com/jalozanot/demoCeiba/infrastructure/app/middlewares/error_handler"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ( rout = MapUrls())


func performRequest(r http.Handler, method, path string, body *model.Movie)  *httptest.ResponseRecorder {

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, _ := http.NewRequest(method,path,payloadBuf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPinController(t *testing.T)  {

	w := performRequest(rout, "GET", "/ping", nil)
	fmt.Println("linea numero 28 ",w.Body,  " ---- ", w.Code)
	valor := fmt.Sprintf("%v", w.Body)
	fmt.Println(valor)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", valor)

}

func TestMovieControllerGets(t *testing.T)  {


	router.Use(error_handler.ErrorHandler())
	UserRepository = InitGetUsersRepository()
	service.Repo = &UserRepository
	service.Han = CreateHandler(UserRepository)
	w := performRequest(rout, "GET", "/peliculas", nil)
	fmt.Println("linea numero 28 ",w.Body,  " ---- ", w.Code)
	valor := fmt.Sprintf("%v", w.Body)
	fmt.Println(valor)
	assert.Equal(t, http.StatusOK, w.Code)

}


func TestMovieControllerSave(t *testing.T)  {

	value := &model.Movie{Id:1,Nombre: "Spiderman",Categoria: "ficcion", CodigoBarras: "1"}
	router.Use(error_handler.ErrorHandler())
	UserRepository = InitGetUsersRepository()
	service.Repo = &UserRepository
	service.Han = CreateHandler(UserRepository)
	w := performRequest(rout, "POST", "/peliculas", value)

	fmt.Println("linea numero 28 ",w.Body,  " ---- ", w.Code)
	valor := fmt.Sprintf("%v", w.Body)
	fmt.Println(valor)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestMovieControllerUpdate(t *testing.T)  {

	value := &model.Movie{Id:1,Nombre: "Spiderman",Categoria: "ficcion", CodigoBarras: "10"}
	router.Use(error_handler.ErrorHandler())
	UserRepository = InitGetUsersRepository()
	service.Repo = &UserRepository
	service.Han = CreateHandler(UserRepository)
	w := performRequest(rout, "PUT", "/peliculas/1", value)

	fmt.Println("linea numero 28 ",w.Body,  " ---- ", w.Code)
	valor := fmt.Sprintf("%v", w.Body)
	fmt.Println(valor)
	assert.Equal(t, http.StatusOK, w.Code)

}


func TestMovieControllerDelete(t *testing.T)  {

	router.Use(error_handler.ErrorHandler())
	UserRepository = InitGetUsersRepository()
	service.Repo = &UserRepository
	service.Han = CreateHandler(UserRepository)
	w := performRequest(rout, "DELETE", "/peliculas/1", nil)

	fmt.Println("linea numero 28 ",w.Body,  " ---- ", w.Code)
	valor := fmt.Sprintf("%v", w.Body)
	fmt.Println(valor)
	assert.Equal(t, http.StatusNoContent, w.Code)

}

func TestMovieControllerGet(t *testing.T)  {

	router.Use(error_handler.ErrorHandler())
	UserRepository = InitGetUsersRepository()
	service.Repo = &UserRepository
	service.Han = CreateHandler(UserRepository)
	w := performRequest(rout, "GET", "/peliculas/1", nil)
	fmt.Println("linea numero 28 ",w.Body,  " ---- ", w.Code)
	valor := fmt.Sprintf("%v", w.Body)
	fmt.Println(valor)
	assert.Equal(t, http.StatusOK, w.Code)

}


func GetDatabaseInstanceTest() *gorm.DB {

	db, err := gorm.Open("mysql", "ceiba:ceiba@tcp(localhost:3306)/movie_ceiba_test?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		_ = db.Close()
		panic("database not working")
	}
	db.SingularTable(true)
	migrateDatabase(db)

	return db
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.MovieEntity{})
}

func InitGetUsersRepository() ports.MoviesRepository {
	return &movie_rep.UserMysqlRepository{
		Db: GetDatabaseInstanceTest(),
	}
}

