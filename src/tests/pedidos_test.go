package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/postech-soat2-grupo16/pedidos-api/api"
	"github.com/postech-soat2-grupo16/pedidos-api/tests/tutils"
)

var baseURL string

func TestFeatures(t *testing.T) {
	server := setup()
	defer server.Close()

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func setup() *http.Server {
	os.Setenv("IS_LOCAL", "true")
	db := api.SetupDB()
	r := api.SetupRouter(db)

	server := http.Server{
		Handler: r,
	}
	serverAddress := tutils.StartNewTestServer(&server)
	baseURL = fmt.Sprintf("http://%s", serverAddress)

	return &server
}
