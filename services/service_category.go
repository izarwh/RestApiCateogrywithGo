package services

import (
	"context"
	"database/sql"
	"prrestapi/helper"
	"prrestapi/model/domain"
	"prrestapi/model/request"
	"prrestapi/model/response"
	"prrestapi/repository"

	"github.com/go-playground/validator/v10"
)

type categoryService struct {
	db                 *sql.DB
	categoryRepository repository.RepositoryCategory
	validate           *validator.Validate
}

func NewCategoryRepository(db *sql.DB, categoryRep repository.RepositoryCategory, validator *validator.Validate) *categoryService {
	return &categoryService{db, categoryRep, validator}
}

func (service *categoryService) FindAll(ctx context.Context) []response.ResponseCategory {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	categoreis := service.categoryRepository.FindAll(ctx, tx)
	var responseCategories []response.ResponseCategory

	for _, v := range categoreis {
		responseCategories = append(responseCategories, v.ToResponseCategory())
	}

	return responseCategories
}

func (service *categoryService) Insert(ctx context.Context, request request.RequestCreateCategory) response.ResponseCategory {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var category_ domain.Category
	category_.SetName(&request.Name)
	// fmt.Println("service" + *category_.GetName())

	category_ = service.categoryRepository.Insert(ctx, tx, category_)

	return category_.ToResponseCategory()
}
func (service *categoryService) Delete(ctx context.Context, request request.RequestDeleteCategory) response.ResponseCategory {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var category_ = service.categoryRepository.FindById(ctx, tx, request.Id)
	service.categoryRepository.Delete(ctx, tx, category_)

	return category_.ToResponseCategory()
}
func (service *categoryService) Update(ctx context.Context, request request.RequestUpdateCategory) response.ResponseCategory {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var category_ = service.categoryRepository.FindById(ctx, tx, request.Id)
	category_.SetName(&request.Name)
	service.categoryRepository.Update(ctx, tx, category_)
	return category_.ToResponseCategory()
}
func (service *categoryService) FindById(ctx context.Context, id int) response.ResponseCategory {
	tx, err := service.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var category_ = service.categoryRepository.FindById(ctx, tx, id)
	return category_.ToResponseCategory()
}
