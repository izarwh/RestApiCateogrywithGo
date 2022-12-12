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
	prodSer := services.NewCategoryService(db, categoryRepo, *validator_)
	prodHandler := handler.NewHandler(prodSer)
	echo_ := echo.New()
	api_ := echo_.Group("/api")
	api_.Use(middleware.AuthMiddleware, exception.ErrorHandlerMiddleware)
	api_.GET("/products", prodHandler.FindAll)
	api_.POST("/products", prodHandler.Create)
	api_.PUT("/products", prodHandler.Update)
	api_.DELETE("/products", prodHandler.Delete)
	api_.GET("/products/:category_id", prodHandler.FindByCategoryId)

	echo_.Start(":3333")
}
