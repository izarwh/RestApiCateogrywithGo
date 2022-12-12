package repository

import (
	"context"
	"database/sql"
	"prrestapi/exception"
	"prrestapi/helper"
	"prrestapi/model/domain"
)

type categoryRepository struct{}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{}
}

func (cr *categoryRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Products {
	query := "SELECT id, category_id, name FROM pr_categories"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var categories []domain.Products

	for rows.Next() {
		var category domain.Products
		err := rows.Scan(category.GetId(), category.GetCategoryId(), category.GetName())
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func (cr *categoryRepository) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Products {
	query := "SELECT id, category_id, name FROM pr_categories WHERE id = ?"
	var category domain.Products
	err := tx.QueryRowContext(ctx, query, id).Scan(category.GetId(), category.GetCategoryId(), category.GetName())
	if err != nil {
		panic(exception.NewNotFound(err))
	}
	return category
}

func (cr *categoryRepository) FindByCategoryId(ctx context.Context, tx *sql.Tx, catId int) []domain.Products {
	query := "SELECT id, category_id, name FROM pr_categories WHERE category_id = ?"
	rows, err := tx.QueryContext(ctx, query, catId)
	if err != nil {
		panic(exception.NewNotFound(err))
	}
	defer rows.Close()
	var categories []domain.Products

	for rows.Next() {
		var category domain.Products
		err := rows.Scan(category.GetId(), category.GetCategoryId(), category.GetName())
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func (cr *categoryRepository) Delete(ctx context.Context, tx *sql.Tx, data domain.Products) {
	query := "DELETE FROM pr_categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, data.GetId())
	if err != nil {
		panic(err)
	}
}

func (cr *categoryRepository) Update(ctx context.Context, tx *sql.Tx, data domain.Products) {
	query := "UPDATE pr_categories SET name = ?, category_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, data.GetName(), data.GetCategoryId(), data.GetId())
	if err != nil {
		panic(err)
	}
}

func (cr *categoryRepository) Create(ctx context.Context, tx *sql.Tx, data domain.Products) domain.Products {
	query := "INSERT INTO pr_categories (name, category_id) VALUES (?, ?)"
	res, err := tx.ExecContext(ctx, query, data.GetName(), data.GetCategoryId())
	helper.PanicIfError(err)
	lastInserted, _ := res.LastInsertId()
	id := int(lastInserted)
	data.SetId(&id)
	return data
}
