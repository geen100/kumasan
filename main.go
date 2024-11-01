package main

import (
	"kumasan/api/app"
	"kumasan/api/controller"
	"kumasan/api/models"
	"kumasan/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

func main() {
	// データベース接続の初期化
	db, err := database.NewDB()
	if err != nil {
		log.Println("データベース接続エラー:", err)
		return
	}
	defer database.CloseDB(db) // CloseDBを使用

	// サービスの設定
	service := goa.New("Sightings-api")
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// コントローラとエンドポイント設定
	repo := models.NewSightingsRepo(db)
	userController := controller.NewSightingsController(service, *repo)
	app.MountSightingsController(service, userController)

	// サーバー起動
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	} else {
		service.LogInfo("Server started successfully on port 8080")
	}
}
