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

	//Product Declaration
	prodRepo := repository.NewproductRepository()
	validator_ := validator.New()
	prodSer := services.NewProductService(db, prodRepo, *validator_)
	prodHandler := handler.NewHandler(prodSer)

	//Product new endpoint
	echo_ := echo.New()
	api_ := echo_.Group("/api")
	api_.Use(middleware.AuthMiddleware, exception.ErrorHandlerMiddleware)
	api_.POST("/products", prodHandler.Create)                               //input product
	api_.PUT("/products", prodHandler.Update)                                //update product
	api_.DELETE("/products", prodHandler.Delete)                             //delete product
	api_.GET("/products", prodHandler.FindAll)                       //ambil semua product
	api_.GET("/category/:category_id/products", prodHandler.FindByCategoryId) //show product by categoryid
	api_.GET("/products/:id", prodHandler.FindById)                          //show product by id

	catrep := repository.NewCategoryRepository()
	categoryservice := services.NewCategoryRepository(db, catrep, validator_)
	categoryHandler := handler.NewCategoryHandler(categoryservice)
	//Category endpoint
	api_.GET("/categories", categoryHandler.FindAll)
	api_.POST("/categories", categoryHandler.Insert)
	api_.PUT("/categories", categoryHandler.Update)
	api_.DELETE("/categories", categoryHandler.Delete)
	api_.GET("/categories/:id", categoryHandler.FindById)
	echo_.Start(":3333")
}
