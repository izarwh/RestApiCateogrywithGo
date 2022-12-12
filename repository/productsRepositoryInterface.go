package repository

import (
	"context"
	"database/sql"
	"prrestapi/model/domain"
)

type ProductRepositoryInterface interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Products
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Products
	FindByCategoryId(ctx context.Context, tx *sql.Tx, catId int) []domain.Products
	Delete(ctx context.Context, tx *sql.Tx, data domain.Products)
	Update(ctx context.Context, tx *sql.Tx, data domain.Products)
	Create(ctx context.Context, tx *sql.Tx, data domain.Products) domain.Products
}
