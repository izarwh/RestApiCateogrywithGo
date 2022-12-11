package domain

import "prrestapi/model/response"

type Category struct {
	id         int
	categoryId int
	name       string
}

func (category *Category) ToResponseCategory() *response.ResponseCategory {
	return &response.ResponseCategory{
		Id:         *category.GetId(),
		CategoryId: *category.GetCategoryId(),
		Name:       *category.GetName(),
	}
}

func (category *Category) GetId() *int {
	return &category.id
}

func (category *Category) SetId(id *int) {
	category.id = *id
}

func (category *Category) GetCategoryId() *int {
	return &category.categoryId
}

func (category *Category) SetCategoryId(categoryId *int) {
	category.categoryId = *categoryId
}

func (category *Category) GetName() *string {
	return &category.name
}

func (category *Category) SetName(name *string) {
	category.name = *name
}

// func (category *Category) GetCategory() *Category {
// 	return &Category{}
// }

func (category *Category) SetCategory(name *string, categoryId *int, id *int) {
	category.name = *name
	category.categoryId = *categoryId
	category.id = *id
}
