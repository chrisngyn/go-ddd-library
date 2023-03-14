//go:build component

package service_test

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/chiennguyen196/go-library/internal/common/server"
	"github.com/chiennguyen196/go-library/internal/common/tests"
	"github.com/chiennguyen196/go-library/internal/lending/ports"
	"github.com/chiennguyen196/go-library/internal/lending/service"
)

const (
	parentRoute = "/api/v1"
)

var (
	httpAddress = ""
)

func TestMain(m *testing.M) {
	httpAddress = fmt.Sprintf("localhost:%d", rand.Intn(9_000)+10_000)
	ok, cleanFn := startService()
	defer cleanFn()
	if !ok {
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func startService() (bool, func()) {
	app, cleanFn := service.NewApplication()

	go server.RunHTTPServerOnAddr(httpAddress, parentRoute, func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})

	ok := tests.WaitForPort(httpAddress)
	if !ok {
		log.Println("Timed out waiting for trainings HTTP to come up")
	}

	return ok, cleanFn
}
