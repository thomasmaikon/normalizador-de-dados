package E2E

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
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
	inputData := dtos.UserDTO{
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
	inputData := dtos.UserDTO{
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

	inputData := dtos.UserDTO{
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

func TestCreateAnCreator(t *testing.T) {
	login := dtos.LoginDTO{
		Email:    "simple-Email4",
		Password: "test",
	}

	inputData := dtos.UserDTO{
		Name:  "test4",
		Login: login,
	}

	creator := dtos.UserDTO{
		Name: "SImple Test Creator",
	}
	result := apitest.New().
		Handler(app).
		Post("/signup").
		JSON(inputData).
		Expect(t).
		HeaderPresent("Authorization").
		Status(http.StatusCreated).
		End()

	jwt := result.Response.Header.Get("Authorization")

	apitest.New().
		Handler(app).
		Post("/creator").
		JSON(creator).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestExpectedErrorWhenCreateAnCreatorWithInvalidToken(t *testing.T) {
	creator := dtos.UserDTO{
		Name: "SImple Test Creator",
	}

	jwt := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJleGFtcGxlNUBob3RtYWlsLmNvbSIsImV4cCI6MTY3MDk4MTg3N30.9p1Q0p2uM7WTRToTizR2GcF_9JxVJdELxZvJDSWnPJw"

	apitest.New().
		Handler(app).
		Post("/creator").
		JSON(creator).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusUnauthorized).
		Body(``).
		End()
}

func TestCreateAnCreatorAndGetIt(t *testing.T) {
	login := dtos.LoginDTO{
		Email:    "simple-Email5",
		Password: "test",
	}

	inputData := dtos.UserDTO{
		Name:  "test5",
		Login: login,
	}

	creator := dtos.UserDTO{
		Name: "SImple Test Creator 2",
	}

	creatorExpected := models.CreatorModel{
		CreatorId: 2,
		Name:      creator.Name,
	}
	outputExpectedJsonCreator := fmt.Sprintf(`{"Amount":0, "Info":{"CreatorId":%d,"Name":"%s"}}`,
		creatorExpected.CreatorId, creatorExpected.Name,)

	result := apitest.New().
		Handler(app).
		Post("/signup").
		JSON(inputData).
		Expect(t).
		HeaderPresent("Authorization").
		Status(http.StatusCreated).
		End()

	jwt := result.Response.Header.Get("Authorization")

	apitest.New().
		Handler(app).
		Post("/creator").
		JSON(creator).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(app).
		Get("/creator").
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusOK).
		Body(outputExpectedJsonCreator).
		End()
}

func TestCreateAnCreatorAndGetItWithInvalidToken(t *testing.T) {
	login := dtos.LoginDTO{
		Email:    "simple-Email6",
		Password: "test",
	}

	inputData := dtos.UserDTO{
		Name:  "test6",
		Login: login,
	}

	creator := dtos.UserDTO{
		Name: "SImple Test Creator 3",
	}

	result := apitest.New().
		Handler(app).
		Post("/signup").
		JSON(inputData).
		Expect(t).
		HeaderPresent("Authorization").
		Status(http.StatusCreated).
		End()

	jwt := result.Response.Header.Get("Authorization")

	apitest.New().
		Handler(app).
		Post("/creator").
		JSON(creator).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusCreated).
		End()

	jwt = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJleGFtcGxlNUBob3RtYWlsLmNvbSIsImV4cCI6MTY3MDk4MTg3N30.9p1Q0p2uM7WTRToTizR2GcF_9JxVJdELxZvJDSWnPJw"

	apitest.New().
		Handler(app).
		Get("/creator").
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}
