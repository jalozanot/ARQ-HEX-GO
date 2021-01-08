package movie_rep

import (
	"context"
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/database_client"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
)

var (
	userMysqlRepository UserMysqlRepository
)

func TestMain(m *testing.M)  {

	containerMockServer, ctx := load()
	code := m.Run()
	beforeAll(containerMockServer,ctx)
	os.Exit(code)

}

func load() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "mysql",
			"MYSQL_DATABASE":      "movie_ceiba",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),

	}
	mysqlC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")
	port := p.Port()
	_ = os.Setenv("MYSQL_USERS_HOST", host)
	_ = os.Setenv("MYSQL_USERS_PORT", port)
	_ = os.Setenv("MYSQL_USERS_SCHEMA", "movie_ceiba")
	_ = os.Setenv("MYSQL_USERS_USERNAME", "root")
	_ = os.Setenv("MYSQL_USERS_PASSWORD", "mysql")

	userMysqlRepository = UserMysqlRepository{
		Db: database_client.GetDatabaseInstance(),
	}
	return mysqlC, ctx
}


func beforeAll(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}

func TestUserMysqlRepository_Save(t *testing.T) {
	tx := userMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie model.Movie

	movie, _ = movie.CreateMovil("Jesus", "chiste", "1")
	movie, err := userMysqlRepository.Save(&movie)

	assert.Nil(t, err)
	assert.EqualValues(t, movie.Nombre, "spider man", "user names are differences")
}
