package repository

import (
	"context"
	"database/sql"
	"prrestapi/model/domain"
)

type CategoryRepositoryInterface interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Category
	FindByCategoryId(ctx context.Context, tx *sql.Tx, catId int) []domain.Category
	Delete(ctx context.Context, tx *sql.Tx, data domain.Category)
	Update(ctx context.Context, tx *sql.Tx, data domain.Category)
	Create(ctx context.Context, tx *sql.Tx, data domain.Category) domain.Category
}
