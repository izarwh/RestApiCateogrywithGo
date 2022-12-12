package services

import (
	"context"
	"prrestapi/model/request"
	"prrestapi/model/response"
)

type ServiceCategory interface {
	FindAll(ctx context.Context) []response.ResponseProducts
	FindById(ctx context.Context, id int) response.ResponseProducts
	FindByCategoryId(ctx context.Context, catId int) []response.ResponseProducts
	Delete(ctx context.Context, req request.RequestDeleteProducts) response.ResponseProducts
	Update(ctx context.Context, req request.RequestUpdateProducts) response.ResponseProducts
	Create(ctx context.Context, req request.RequestCreateProducts) response.ResponseProducts
}
