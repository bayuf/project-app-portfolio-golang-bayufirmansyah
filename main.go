package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/db"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/handler"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/router"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/utils"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load()

	// initLogger
	logger, err := utils.InitLogger("./logs/apps-", true)
	if err != nil {
		fmt.Println(err)
	}

	//initdb
	conn, err := db.Connect(logger)
	if err != nil {
		logger.Fatal("cant init connect to database", zap.Error(err))
	}
	defer conn.Close()

	// init template
	var templates = template.Must(template.New("").
		ParseGlob("views/**/*.html"))

	// init layer
	repo := repository.NewRepository(conn, logger)
	svc := services.NewService(repo, logger)
	handler := handler.NewHandler(svc, templates, logger)
	router := router.NewRouter(svc, handler, logger)

	// public folder permission
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fs))

	// start server
	fmt.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
