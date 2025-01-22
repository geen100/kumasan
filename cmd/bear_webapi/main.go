package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/webapi"
	"backend/webapi/app"

	goa "github.com/shogo82148/goa-v1"
)

func main() {

	// Goaのサービスを初期化
	service := goa.New("BearAPI")
	// BearControllerを初期化 (Wireを利用)
	bearController, err := webapi.InitializeBearController(service)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize BearController: %v", err)
	}
	log.Println("[DEBUG] BearController initialized.") // コントローラー初期化ログ

	// Goaで生成されたルーティングコードにBearControllerをマウント
	app.MountBearController(service, bearController)
	log.Println("[DEBUG] BearController mounted to Goa service.") // ルーティングコードにマウントログ

	// サーバーを構成
	addr := ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: service.Mux,
	}

	// サーバーをバックグラウンドで起動
	go func() {
		log.Printf("[DEBUG] Starting server on %s", addr) // サーバー起動ログ
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] Failed to start server: %v", err)
		}
	}()

	// シグナルをキャッチ
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	log.Println("[DEBUG] Waiting for shutdown signal...") // シャットダウンシグナル待ちログ
	<-chSignal

	// サーバーのシャットダウン処理
	log.Println("[DEBUG] Shutting down server...") // シャットダウン開始ログ
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] Failed to gracefully shutdown server: %v", err)
	}
	log.Println("[DEBUG] Server gracefully stopped.") // シャットダウン完了ログ
}
