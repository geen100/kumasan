package repository

import (
	"backend/app/model"
	"context"
)

type (
	// CreateNotificationInput CreateNotification の入力
	CreatebearInput struct {
		Latitude  float64 // 緯度
		Longitude float64 // 経度
		City      string  // 都道府県
		Address   string  // 住所
		Date      string  // 西暦年月日
		Time      string  //  時間
		Details   string  // 詳細
	}

	// CreateNotificationOutput CreateNotification の出力
	CreatebearOutput struct {
		ID model.BearID
	}

	// NotificationCreator 通知を作成するインターフェース
	BearCreator interface {
		// CreateNotification 通知を作成する。
		CreateBear(context.Context, Execer, *CreatebearInput) (*CreatebearOutput, error)
	}
)
