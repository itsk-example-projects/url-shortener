package core

import (
	"go.uber.org/fx"

	"url-shortener/internal/config"
	"url-shortener/internal/storage"
	"url-shortener/internal/web"
	"url-shortener/pkg/postgres"
)

func Load() *fx.App {
	return fx.New(
		fx.Invoke(
			config.LoadConfig,
		),
		fx.Provide(
			postgres.NewInstance,
			storage.NewLinkStorage,
			web.NewEngine,
			web.RegisterHandlers,
			web.RegisterTemplates,
		),
		fx.Invoke(
			web.RunServer,
		),
	)
}
