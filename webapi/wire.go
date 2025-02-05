//go:build wireinject
// +build wireinject

package webapi

import (
	"backend/app/datastore/mysql"
	"backend/app/interactor"
	"backend/app/usecase"
	"backend/database"
	"backend/webapi/app"

	"github.com/google/wire"
	goa "github.com/shogo82148/goa-v1"
)

// WireSet は webapi の依存性を解決するための Wire セット
var WireSet = wire.NewSet(
	database.NewDB,
	mysql.NewMySQLConfig,
	interactor.ConsoleSet, // interactor の依存関係を注入
	mysql.Set,             // datastore の依存関係を注入

	NewBearControllerConfig,
	NewBearController,
	wire.Bind(new(app.BearController), new(*BearController)),
)

func NewBearControllerConfig(creater usecase.BearCreator, getter usecase.BearGetter) *BearControllerConfig {
	return &BearControllerConfig{
		creater: creater,
		getter:  getter,
	}
}

// InitializeBearController は BearController の依存関係を初期化します。
func InitializeBearController(service *goa.Service) (*BearController, error) {
	wire.Build(WireSet)
	return nil, nil
}
