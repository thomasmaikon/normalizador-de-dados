package E2E

import (
	"log"
	"net/http"
	"os"
	"strings"

	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/setup"
	"hubla/desafiofullstack/utils"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/steinfletcher/apitest"
)

var app *gin.Engine

func TestMain(m *testing.M) {

	dockerInfo := StartPostgresDockerFormTest()
	strs := strings.Split(dockerInfo.HostAndPort, ":")
	port := strs[1]

	dbConnection := utils.GetDBWithParams(port, dockerInfo.User, dockerInfo.Password, dockerInfo.Dbname)

	// initialize app
	app = setup.NewAppEngine().InitializeRoutes().RunMigrations(dbConnection).Router

	//Run tests
	code := m.Run()

	if err := dockerInfo.Pool.Purge(dockerInfo.Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSimpleRequestForCreateUser(t *testing.T) {
	inputData := dtos.CreateUseDTO{
		Name: "test",
		Login: dtos.LoginDTO{
			Email:    "simple-Email",
			Password: "test",
		},
	}

	apitest.New().
		Handler(app).
		Post("/signup").
		JSON(inputData).
		Expect(t).
		HeaderPresent("Authorization").
		Status(http.StatusCreated).
		End()
}

func TestSimpleRequestForCreateUserthatAlredyExist(t *testing.T) {
	inputData := dtos.CreateUseDTO{
		Name: "test",
		Login: dtos.LoginDTO{
			Email:    "simple-Email2",
			Password: "test",
		},
	}

	apitest.New().
		Handler(app).
		Post("/signup").
		JSON(inputData).
		Expect(t).
		HeaderPresent("Authorization").
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(app).
		Post("/signup").
		JSON(inputData).
		Expect(t).
		HeaderNotPresent("Authorization").
		Status(http.StatusConflict).
		End()
}

func TestCreateAccoutAndSignIn(t *testing.T) {
	login := dtos.LoginDTO{
		Email:    "simple-Email3",
		Password: "test",
	}

	inputData := dtos.CreateUseDTO{
		Name:  "test",
		Login: login,
	}

	apitest.New().
		Handler(app).
		Post("/signup").
		JSON(inputData).
		Expect(t).
		HeaderPresent("Authorization").
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(app).
		Post("/signin").
		JSON(login).
		Expect(t).
		HeaderPresent("Authorization").
		Status(http.StatusAccepted).
		End()
}
