package handler

import (
	"fmt"
	"net/http"
	"prrestapi/model/api"
	"prrestapi/model/request"
	"prrestapi/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productService services.ServiceProducts
}

func NewHandler(prodService services.ServiceProducts) productHandler {
	return productHandler{prodService}
}

func (handler *productHandler) FindAll(c echo.Context) error {
	response := handler.productService.FindAll(c.Request().Context())
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *productHandler) FindById(c echo.Context) error {
	idParam := c.Param("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}

	response := handler.productService.FindById(c.Request().Context(), id)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *productHandler) FindByCategoryId(c echo.Context) error {
	catidParam := c.Param("category_id")
	fmt.Println("cek",catidParam)
	catid, err := strconv.Atoi(catidParam)
	if err != nil {
		panic(err)
	}

	response := handler.productService.FindByCategoryId(c.Request().Context(), catid)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *productHandler) Delete(c echo.Context) error {
	requestDeleteCategory := new(request.RequestDeleteProducts)
	c.Bind(requestDeleteCategory)
	response := handler.productService.Delete(c.Request().Context(), *requestDeleteCategory)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *productHandler) Update(c echo.Context) error {

	requestUpdateCategory := new(request.RequestUpdateProducts)
	c.Bind(requestUpdateCategory)
	response := handler.productService.Update(c.Request().Context(), *requestUpdateCategory)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *productHandler) Create(c echo.Context) error {
	// fmt.Println("testHandler")
	requestCreateCategory := new(request.RequestCreateProducts)
	c.Bind(requestCreateCategory)
	fmt.Println(requestCreateCategory.CategoryId)
	response := handler.productService.Create(c.Request().Context(), *requestCreateCategory)
	// fmt.Println("testHandler")
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
