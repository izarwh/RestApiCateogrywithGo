package services

import (
	"context"
	"prrestapi/model/request"
	"prrestapi/model/response"
)

type ServiceCategory interface {
	FindAll(ctx context.Context) []response.ResponseCategory
	Delete(ctx context.Context, request request.RequestDeleteCategory) response.ResponseCategory
	Update(ctx context.Context, request request.RequestUpdateCategory) response.ResponseCategory
	Insert(ctx context.Context, request request.RequestCreateCategory) response.ResponseCategory
	FindById(ctx context.Context, id int) response.ResponseCategory
}
