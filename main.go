package main

import (
	"prrestapi/database"
	"prrestapi/exception"
	"prrestapi/handler"
	"prrestapi/middleware"
	"prrestapi/repository"
	"prrestapi/services"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

//buat handler DONE
//Buat error handler
//buat middleware
//done!

func main() {
	db := database.GetConnection()
	categoryRepo := repository.NewCategoryRepository()
	validator_ := validator.New()
	categorySer := services.NewCategoryService(db, categoryRepo, *validator_)
	categoryHandler := handler.NewHandler(categorySer)
	echo_ := echo.New()
	api_ := echo_.Group("/api")
	api_.Use(middleware.AuthMiddleware, exception.ErrorHandlerMiddleware)
	api_.GET("/categories", categoryHandler.FindAll)
	api_.POST("/categories", categoryHandler.Create)
	api_.PUT("/categories", categoryHandler.Update)
	api_.DELETE("/categories", categoryHandler.Delete)
	api_.GET("/categories/:category_id", categoryHandler.FindByCategoryId)

	echo_.Start(":3333")
}
