package interactor

import (
	"backend/app/model"
	"backend/app/repository"
	"backend/app/usecase"
	"context"
	"errors"
	"fmt"
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
	created, err := c.cfg.BearCreator.CreateBear(ctx, tx, &repository.CreateBearInput{
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

	return &usecase.CreateBearOutput{Bear: model.Bear{
		ID:        created.ID,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		City:      input.City,
		Address:   input.Address,
		Date:      input.Date,
		Time:      input.Time,
		Details:   input.Details,
	}}, nil
}

// BearGetterconfig クマ情報の取得ユースケース作成
type BearGetterconfig struct {
	Queryer    repository.Queryer
	bearGetter repository.BearGetter
}

type bearGetter struct {
	cfg *BearGetterconfig
}

var _ usecase.BearGetter = (*bearGetter)(nil)

// NewBearGetter クマ情報の取得ユースケース生成
func NewBearGetter(cfg *BearGetterconfig) *bearGetter {
	return &bearGetter{cfg: cfg}
}

// GetBear implements usecase.BearGetter interface
func (g *bearGetter) GetBear(ctx context.Context, input *usecase.GetBearInput) (*usecase.GetBearOutput, error) {
	bear, err := g.cfg.bearGetter.GetBear(ctx, g.cfg.Queryer, &repository.GetBearInput{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, repository.ErrBearNotFound) {
			return nil, fmt.Errorf("bear not found: %w", usecase.ErrResourceNotFound)
		}
		return nil, err
	}

	output := &usecase.GetBearOutput{
		Bear: bear.Bear,
	}
	return output, nil

}
