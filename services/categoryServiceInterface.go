package services

import (
	"context"
	"prrestapi/model/request"
	"prrestapi/model/response"
)

type ServiceCategory interface {
	FindAll(ctx context.Context) []response.ResponseCategory
	FindById(ctx context.Context, id int) response.ResponseCategory
	FindByCategoryId(ctx context.Context, catId int) []response.ResponseCategory
	Delete(ctx context.Context, req request.RequestDeleteCategory) response.ResponseCategory
	Update(ctx context.Context, req request.RequestUpdateCategory) response.ResponseCategory
	Create(ctx context.Context, req request.RequestCreateCategory) response.ResponseCategory
}
