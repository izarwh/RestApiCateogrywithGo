package services

import (
	"context"
	"database/sql"
	"fmt"
	"prrestapi/helper"
	"prrestapi/model/domain"
	"prrestapi/model/request"
	"prrestapi/model/response"
	"prrestapi/repository"

	"github.com/go-playground/validator/v10"
)

type categoryService struct {
	db                 *sql.DB
	categoryRepository repository.CategoryRepositoryInterface
	validator          *validator.Validate
}

func NewCategoryService(db *sql.DB, catRepo repository.CategoryRepositoryInterface, valid validator.Validate) *categoryService {
	return &categoryService{db, catRepo, &valid}
}

func (cs *categoryService) FindAll(ctx context.Context) []response.ResponseProducts {
	tx, err := cs.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		panic(err)
	}
	products := cs.categoryRepository.FindAll(ctx, tx)
	var responseCategories []response.ResponseProducts
	for _, prod := range products {
		responseCategories = append(responseCategories, *prod.ToResponseProducts())
	}
	return responseCategories
}
func (cs *categoryService) FindById(ctx context.Context, id int) response.ResponseProducts {
	tx, err := cs.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// if err != nil {
	// 	panic(err)
	// }
	categoryById := cs.categoryRepository.FindById(ctx, tx, id)
	return *categoryById.ToResponseProducts()
}
func (cs *categoryService) FindByCategoryId(ctx context.Context, catId int) []response.ResponseProducts {
	tx, err := cs.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	// defer helper.CommitOrRollback(tx)
	// if err != nil {
	// 	panic(err)
	// }
	categories := cs.categoryRepository.FindByCategoryId(ctx, tx, catId)
	var responseCategories []response.ResponseProducts
	for _, cat := range categories {
		responseCategories = append(responseCategories, *cat.ToResponseProducts())
	}
	return responseCategories
}
func (cs *categoryService) Delete(ctx context.Context, req request.RequestDeleteProducts) response.ResponseProducts {
	err := cs.validator.Struct(req)
	if err != nil {
		panic(err)
	}
	tx, err := cs.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// if err != nil {
	// 	panic(err)
	// }
	catToDelete := cs.categoryRepository.FindById(ctx, tx, req.Id)
	fmt.Println(catToDelete)
	cs.categoryRepository.Delete(ctx, tx, catToDelete)
	return *catToDelete.ToResponseProducts()
}
func (cs *categoryService) Update(ctx context.Context, req request.RequestUpdateProducts) response.ResponseProducts {
	err := cs.validator.Struct(req)
	if err != nil {
		panic(err)
	}
	tx, err := cs.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// if err != nil {
	// 	panic(err)
	// }
	catToUpdate := cs.categoryRepository.FindById(ctx, tx, req.Id)
	catToUpdate.SetCategoryId(&req.CategoryId)
	catToUpdate.SetName(&req.Name)
	cs.categoryRepository.Update(ctx, tx, catToUpdate)
	return *catToUpdate.ToResponseProducts()
}
func (cs *categoryService) Create(ctx context.Context, req request.RequestCreateProducts) response.ResponseProducts {
	err := cs.validator.Struct(req)
	if err != nil {
		panic(err)
	}
	// fmt.Println("test")
	tx, err := cs.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(req)
	var reqToDomain domain.Products
	reqToDomain.SetCategoryId(&req.CategoryId)
	reqToDomain.SetName(&req.Name)
	resDomain := cs.categoryRepository.Create(ctx, tx, reqToDomain)
	return *resDomain.ToResponseProducts() //Data yang diinput harus berupa json
}
