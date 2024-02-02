package main

import (
	"Hanivan/learn-golang-restfull-api/app"
	"Hanivan/learn-golang-restfull-api/controller"
	"Hanivan/learn-golang-restfull-api/helper"
	"Hanivan/learn-golang-restfull-api/middleware"
	"Hanivan/learn-golang-restfull-api/repository"
	"Hanivan/learn-golang-restfull-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
