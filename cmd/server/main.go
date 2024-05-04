package main

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/structx/orchestrator/internal/adapter/port/http/router"
	"github.com/structx/orchestrator/internal/core/domain"
	"github.com/structx/orchestrator/internal/core/service"
)

func main() {
	fx.New(
		fx.Provide(fx.Annotate(service.NewAvailable, fx.As(new(domain.AvailableService)))),
		fx.Provide(fx.Annotate(router.New, fx.As(new(http.Handler)))),
	).Run()
}
