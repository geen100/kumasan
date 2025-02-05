package mysql

import (
	"backend/app/model"
	"backend/app/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/gofrs/uuid"
)

type bearCreator struct{}

var _ repository.BearCreator = (*bearCreator)(nil)

func NewBearCreater() *bearCreator {
	return &bearCreator{}
}

func (*bearCreator) CreateBear(ctx context.Context, execer repository.Execer, input *repository.CreateBearInput) (*repository.CreateBearOutput, error) {
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

type bearGetter struct{}

var _ repository.BearGetter = (*bearGetter)(nil)

func NewBearGetter() *bearGetter {
	return &bearGetter{}
}

func (*bearGetter) GetBear(ctx context.Context, queryer repository.Queryer, input *repository.GetBearInput) (*repository.GetBearOutput, error) {
	var bear model.Bear

	query :=
		"SELECT" +
			"`id`, `latitude`, `longitude`, `city`, `address`, `date`, `time`, `details` " +
			"FROM `bear` " +
			"WHERE `id`  = ?"

	if err := queryer.QueryRowContext(ctx, query, input.ID).Scan(
		&bear.ID, &bear.Latitude, &bear.Longitude, &bear.City, &bear.Address,
		&bear.Date, &bear.Time, &bear.Details,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrBearNotFound
		}
		return nil, err
	}
	return &repository.GetBearOutput{Bear: model.Bear{
		ID:        bear.ID,
		Latitude:  bear.Latitude,
		Longitude: bear.Longitude,
		City:      bear.City,
		Address:   bear.Address,
		Date:      bear.Date,
		Time:      bear.Time,
		Details:   bear.Details,
	}}, nil
}
