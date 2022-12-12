package repository

import (
	"context"
	"database/sql"
	"prrestapi/exception"
	"prrestapi/helper"
	"prrestapi/model/domain"
)

type categoryRepository struct {
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{}
}

func (c *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "select id, name from categories"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()
	var categoriesSlice []domain.Category

	for rows.Next() {
		var category domain.Category
		// var id int
		// var name string
		err = rows.Scan(category.GetId(), category.GetName())
		helper.PanicIfError(err)
		// category.SetCategory(id, name)
		categoriesSlice = append(categoriesSlice, category)
	}

	return categoriesSlice
}
func (c *categoryRepository) Delete(ctx context.Context, tx *sql.Tx, dom domain.Category) {
	query := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, query, dom.GetId())
	helper.PanicIfError(err)
}
func (c *categoryRepository) Update(ctx context.Context, tx *sql.Tx, dom domain.Category) {
	query := "update categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, dom.GetName(), dom.GetId())
	helper.PanicIfError(err)
}
func (c *categoryRepository) Insert(ctx context.Context, tx *sql.Tx, dom domain.Category) domain.Category {
	query := "insert into categories (name) values (?)"
	res, err := tx.ExecContext(ctx, query, dom.GetName())
	helper.PanicIfError(err)
	val, _ := res.LastInsertId()
	val_ := int(val)
	dom.SetId(&val_)
	return dom
}
func (c *categoryRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Category {
	query := "select name from categories where id = ?"
	var category domain.Category
	category.SetId(&id)
	err := tx.QueryRowContext(ctx, query, id).Scan(category.GetName())
	if err != nil {
		panic(exception.NewNotFound(err))
	}
	return category
}
