package ports

import (
	"github.com/chiennguyen196/go-library/internal/lending/app"
)

type Job struct {
	app app.Application
}

func NewJob(app app.Application) Job {
	return Job{app: app}
}
