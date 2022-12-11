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

type categoryHandler struct {
	categoryservice services.ServiceCategory
}

func NewHandler(catService services.ServiceCategory) categoryHandler {
	return categoryHandler{catService}
}

func (handler *categoryHandler) FindAll(c echo.Context) error {
	response := handler.categoryservice.FindAll(c.Request().Context())
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *categoryHandler) FindById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}

	response := handler.categoryservice.FindById(c.Request().Context(), id)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *categoryHandler) FindByCategoryId(c echo.Context) error {
	catidParam := c.Param("category_id")
	catid, err := strconv.Atoi(catidParam)
	if err != nil {
		panic(err)
	}

	response := handler.categoryservice.FindByCategoryId(c.Request().Context(), catid)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *categoryHandler) Delete(c echo.Context) error {
	requestDeleteCategory := new(request.RequestDeleteCategory)
	c.Bind(requestDeleteCategory)
	response := handler.categoryservice.Delete(c.Request().Context(), *requestDeleteCategory)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *categoryHandler) Update(c echo.Context) error {

	requestUpdateCategory := new(request.RequestUpdateCategory)
	c.Bind(requestUpdateCategory)
	response := handler.categoryservice.Update(c.Request().Context(), *requestUpdateCategory)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
func (handler *categoryHandler) Create(c echo.Context) error {
	// fmt.Println("testHandler")
	requestCreateCategory := new(request.RequestCreateCategory)
	c.Bind(requestCreateCategory)
	fmt.Println(requestCreateCategory.CategoryId)
	response := handler.categoryservice.Create(c.Request().Context(), *requestCreateCategory)
	// fmt.Println("testHandler")
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
