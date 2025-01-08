package interactor

import (
	"backend/app/model"
	"backend/app/repository"
	"backend/app/usecase"
	"context"
)

// BearCreatorConfig クマ情報作成ユースケースの設定
type BearCreatorConfig struct {
	TxBeginner  repository.TxBeginner
	BearCreator repository.BearCreator
}

type bearCreater struct {
	cfg *BearCreatorConfig
}

var _ usecase.BearCreator = (*bearCreater)(nil)

// NewbearCreater クマ情報作成ユースケースを生成する
func NewBearCreater(cfg *BearCreatorConfig) *bearCreater {
	return &bearCreater{cfg: cfg}
}

// CreateBear implements usecase.BearCreator interface
func (c *bearCreater) CreateBear(ctx context.Context, input *usecase.CreateBearInput) (*usecase.CreateBearOutput, error) {
	// トランザクションを開始
	tx, err := c.cfg.TxBeginner.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// クマ情報を登録
	created, err := c.cfg.BearCreator.CreateBear(ctx, tx, &repository.CreatebearInput{
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		City:      input.City,
		Address:   input.Address,
		Date:      input.Date,
		Time:      input.Time,
		Details:   input.Details,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// 作成されたクマ情報をモデル化
	bear := model.Bear{
		ID:        created.ID,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		City:      input.City,
		Address:   input.Address,
		Date:      input.Date,
		Time:      input.Time,
		Details:   input.Details,
	}

	return &usecase.CreateBearOutput{
		Bear: bear,
	}, nil
}
