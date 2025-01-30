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
	service := goa.New("BearAPI")

	bearController, err := webapi.InitializeBearController(service)
	if err != nil {
		log.Fatalf("Failed to initialize BearController: %v", err)
	}

	app.MountBearController(service, bearController)

	addr := ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: service.Mux,
	}

	// サーバーをバックグラウンドで起動
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// シグナルをキャッチ
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	<-chSignal

	// サーバーのシャットダウン処理
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown server: %v", err)
	}
}
