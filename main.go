package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/handler"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/repository"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/router"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/services"
	"github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/utils"
)

func main() {
	var templates = template.Must(template.New("").ParseGlob("views/**/*.html"))
	for _, t := range templates.Templates() {
		fmt.Println("template:", t.Name())
	}

	// initdb
	// conn, err := database.InitDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer conn.Close(context.Background())

	// init
	logger, err := utils.InitLogger("./logs/apps-", true)
	if err != nil {
		fmt.Println(err)
	}
	repo := repository.NewRepository()
	svc := services.NewService(repo)
	handler := handler.NewHandler(svc, templates, logger)
	router := router.NewRouter(svc, handler, logger)

	// public folder permission
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fs))

	fmt.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
