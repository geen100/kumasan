package usecase

import (
	"backend/app/model"
	"context"
)

type (
	// CreateBearInput CreateBear の入力
	CreateBearInput struct {
		Latitude  float64 // 緯度
		Longitude float64 // 経度
		City      string  // 都道府県
		Address   string  // 住所
		Date      string  // 西暦年月日
		Time      string  //  時間
		Details   string  // 詳細
	}

	// CreateBearOutput CreateBear の出力
	CreateBearOutput struct {
		Bear model.Bear
	}

	// BearCreator  クマ情報通を作成するインターフェース
	BearCreator interface {
		// CreateBear クマ情報を作成する
		CreateBear(context.Context, *CreateBearInput) (*CreateBearOutput, error)
	}
)
