package exception

import (
	"fmt"
	"net/http"
	"prrestapi/model/api"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func(c echo.Context) error {
			err := recover()
			//validation  
			validationErrors, ok := err.(validator.ValidationErrors)
			fmt.Println("test0", validationErrors)

			if ok {
				return c.JSON(http.StatusNotFound, api.ApiResponse{
					Data:   validationErrors.Error(),
					Status: "INVALID",
					Code:   http.StatusNotFound,
				})
			}

			notFoundErr, ok := err.(NotFoundError)
			fmt.Println("test1", notFoundErr)
			if ok {
				return c.JSON(http.StatusNotFound, api.ApiResponse{
					Data:   notFoundErr.Err.Error(),
					Status: "NOT FOUND",
					Code:   http.StatusNotFound,
				})
			}

			internalServerError, ok := err.(InternalServerError)
			fmt.Println("test2", internalServerError)
			if ok {
				return c.JSON(http.StatusInternalServerError, api.ApiResponse{
					Data:   internalServerError.Err.Error(),
					Status: "INTERNAL SERVER ERROR",
					Code:   http.StatusInternalServerError,
				})
			}

			return c.NoContent(http.StatusInternalServerError)
		}(c)
		return next(c)
	}
}
