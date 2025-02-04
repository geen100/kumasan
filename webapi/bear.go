package webapi

//go:generate go run gen/main.go

import (
	"errors"

	"backend/app/model"
	"backend/app/usecase"
	"backend/webapi/app"

	goa "github.com/shogo82148/goa-v1"
)

type BearControllerConfig struct {
	creater usecase.BearCreator
	getter  usecase.BearGetter
}

// NotificationController はPUSH通知を管理します。
type BearController struct {
	*goa.Controller
	cfg *BearControllerConfig
}

// NewBearController creates a bear controller.
func NewBearController(service *goa.Service, cfg *BearControllerConfig) *BearController {
	return &BearController{
		Controller: service.NewController("BearController"),
		cfg:        cfg,
	}
}

// Add runs the add action.
func (c *BearController) Add(ctx *app.AddBearContext) error {
	// 入力をユースケース層に渡す
	out, err := c.cfg.creater.CreateBear(ctx, &usecase.CreateBearInput{
		Latitude:  ctx.Payload.Latitude,
		Longitude: ctx.Payload.Longitude,
		City:      ctx.Payload.City,
		Address:   ctx.Payload.Address,
		Date:      ctx.Payload.Date,
		Time:      ctx.Payload.Time,
		Details:   ctx.Payload.Details,
	})
	if err != nil {
		var invalidErr *usecase.InvalidInputError
		if errors.As(err, &invalidErr) {
			return ctx.BadRequest()
		}

		return err
	}

	// 成功レスポンスを作成
	res := &app.Sighting{
		ID:        string(out.Bear.ID),
		Latitude:  out.Bear.Latitude,
		Longitude: out.Bear.Longitude,
		City:      out.Bear.City,
		Address:   out.Bear.Address,
		Date:      out.Bear.Date,
		Time:      out.Bear.Time,
		Details:   out.Bear.Details,
	}

	return ctx.Created(res)
}

func (c *BearController) Get(ctx *app.GetBearContext) error {
	out, err := c.cfg.getter.GetBear(ctx, &usecase.GetBearInput{
		ID: model.BearID(ctx.ID),
	})
	if err != nil {
		if errors.Is(err, usecase.ErrResourceNotFound) {
			return ctx.NotFound()
		}
		return err
	}

	return ctx.OK(&app.Sighting{
		ID:        string(out.Bear.ID),
		Latitude:  out.Bear.Latitude,
		Longitude: out.Bear.Longitude,
		City:      out.Bear.City,
		Address:   out.Bear.Address,
		Date:      out.Bear.Date,
		Time:      out.Bear.Time,
		Details:   out.Bear.Details,
	})
}
