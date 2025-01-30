package mysql

import (
	"backend/app/model"
	"backend/app/repository"
	"context"

	"github.com/gofrs/uuid"
)

type bearCreater struct{}

var _ repository.BearCreator = (*bearCreater)(nil)

func NewBearCreator() *bearCreater {
	return &bearCreater{}
}

func (*bearCreater) CreateBear(ctx context.Context, execer repository.Execer, input *repository.CreateBearInput) (*repository.CreateBearOutput, error) {
	id := uuid.Must(uuid.NewV7()).String()

	_, err := execer.ExecContext(ctx,
		"INSERT INTO `bear` "+
			"(`id`,`latitude`,`longitude`,`city`,`address`,`date`,`time`,`details`)"+
			"VALUES (?,?,?,?,?,?,?,?)",
		id, input.Latitude, input.Longitude, input.City, input.Address, input.Date, input.Time, input.Details,
	)
	if err != nil {
		return nil, err
	}

	return &repository.CreateBearOutput{
		ID: model.BearID(id),
	}, nil
}
