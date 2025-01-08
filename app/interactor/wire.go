package interactor

import (
	"backend/app/usecase"

	"github.com/google/wire"
)

var ConsoleSet = wire.NewSet(
	NewBearCreater,
	wire.Struct(new(BearCreatorConfig), "*"), wire.Bind(new(usecase.BearCreator), new(*bearCreater)),
)
