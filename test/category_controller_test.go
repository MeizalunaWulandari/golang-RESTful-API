package test

import (
	"database/sql"
	"golang-resful-api/app"
	"golang-resful-api/controller"
	"golang-resful-api/helper"
	"golang-resful-api/middleware"
	"golang-resful-api/repository"
	"golang-resful-api/service"
	"net/http"
	"time"

	"github.com/go-playground/validator"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3360)/golang_resful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}
