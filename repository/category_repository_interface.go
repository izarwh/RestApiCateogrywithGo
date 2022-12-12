package repository

import (
	"context"
	"database/sql"
	"prrestapi/model/domain"
)

type RepositoryCategory interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	Delete(ctx context.Context, tx *sql.Tx, dom domain.Category)
	Update(ctx context.Context, tx *sql.Tx, dom domain.Category)
	Insert(ctx context.Context, tx *sql.Tx, dom domain.Category) domain.Category
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Category
}
