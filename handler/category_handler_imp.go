package handler

import (
	"fmt"
	"net/http"
	"prrestapi/helper"
	"prrestapi/model/api"
	"prrestapi/model/request"
	"prrestapi/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type categoryHandler struct {
	categoryService services.ServiceCategory
}

func NewCategoryHandler(categoryservice services.ServiceCategory) *categoryHandler {
	return &categoryHandler{categoryService: categoryservice}
}

func (handler *categoryHandler) FindAll(c echo.Context) error {
	responseCategories := handler.categoryService.FindAll(c.Request().Context())

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseCategories,
	},
	)
}

func (handler *categoryHandler) Delete(c echo.Context) error {
	requestCategory := new(request.RequestDeleteCategory)
	c.Bind(requestCategory)

	handler.categoryService.Delete(c.Request().Context(), *requestCategory)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	})

}
func (handler *categoryHandler) Update(c echo.Context) error {
	requestCategory := new(request.RequestUpdateCategory)
	c.Bind(requestCategory)
	responseCategory := handler.categoryService.Update(c.Request().Context(), *requestCategory)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseCategory,
	})
}
func (handler *categoryHandler) Insert(c echo.Context) error {
	requestCategory := new(request.RequestCreateCategory)
	c.Bind(requestCategory)
	fmt.Println(requestCategory)
	responseCategory := handler.categoryService.Insert(c.Request().Context(), *requestCategory)

	fmt.Println("test" + responseCategory.Name)
	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseCategory,
	},
	)
}
func (handler *categoryHandler) FindById(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfError(err)

	responseCategory := handler.categoryService.FindById(c.Request().Context(), id)

	return c.JSON(http.StatusOK, api.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   responseCategory,
	},
	)
}
