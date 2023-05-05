package integration

import (
	"hubla/desafiofullstack/setup"
	"hubla/desafiofullstack/utils"
	"log"
	"os"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {

	dockerInfo := StartPostgresDockerFormTest()
	strs := strings.Split(dockerInfo.HostAndPort, ":")
	port := strs[1]

	dbConnection := utils.GetDBWithParams(port, dockerInfo.User, dockerInfo.Password, dockerInfo.Dbname)

	// initialize app
	setup.NewAppEngine().RunMigrations(dbConnection)

	//Run tests
	code := m.Run()

	if err := dockerInfo.Pool.Purge(dockerInfo.Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
