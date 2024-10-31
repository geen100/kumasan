package main

import (
	"backend/api/app"
	"backend/api/controller"
	"backend/api/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Println("Failed to ping database:", err)
		return
	}

	service := goa.New("Sightings-api")
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	Repo := models.NewSightingsRepo(db)
	userController := controller.NewSightingsController(service, *Repo)
	app.MountSightingsController(service, userController)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	} else {
		service.LogInfo("Seaver started succsesfull on port 8080")
	}
}
