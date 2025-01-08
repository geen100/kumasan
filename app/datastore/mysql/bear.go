package mysql

import (
	"backend/app/model"
	"backend/app/repository"
	"context"

	"github.com/shogo82148/goa-v1/uuid"
)

type bearCreater struct{}

var _ repository.BearCreator = (*bearCreater)(nil)

func NewBearCreator() *bearCreater {
	return &bearCreater{}
}

func (*bearCreater) CreateBear(ctx context.Context, execer repository.Execer, input *repository.CreatebearInput) (*repository.CreatebearOutput, error) {
	id := uuid.NewV4().String()

	_, err := execer.ExecContext(ctx,
		"INSERT INTO `bear` "+
			"(`id`,`latitude`,`longitude`,`city`,`address`,`date`,`time`,`details`)"+
			"VALUES (?,?,?,?,?,?,?,?)",
		id, input.Latitude, input.Longitude, input.City, input.Address, input.Date, input.Time, input.Details,
	)
	if err != nil {
		return nil, err
	}

	return &repository.CreatebearOutput{
		ID: model.BearID(id),
	}, nil
}
