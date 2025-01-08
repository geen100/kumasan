package mysql

import (
	"backend/app/repository"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	New,
	wire.Bind(new(repository.DB), new(*DB)),
	wire.Bind(new(repository.TxBeginner), new(*DB)),
	wire.Bind(new(repository.Queryer), new(*DB)),
	NewBearCreator,
	wire.Bind(new(repository.BearCreator), new(*bearCreater)),
)
