package domain

import "prrestapi/model/response"

type Products struct {
	id         int
	categoryId int
	name       string
}

func (category *Products) ToResponseProducts() *response.ResponseProducts {
	return &response.ResponseProducts{
		Id:         *category.GetId(),
		CategoryId: *category.GetCategoryId(),
		Name:       *category.GetName(),
	}
}

func (category *Products) GetId() *int {
	return &category.id
}

func (category *Products) SetId(id *int) {
	category.id = *id
}

func (category *Products) GetCategoryId() *int {
	return &category.categoryId
}

func (category *Products) SetCategoryId(categoryId *int) {
	category.categoryId = *categoryId
}

func (category *Products) GetName() *string {
	return &category.name
}

func (category *Products) SetName(name *string) {
	category.name = *name
}

// func (category *Products) GetCategory() *Products {
// 	return &Products{}
// }

func (category *Products) SetCategory(name *string, categoryId *int, id *int) {
	category.name = *name
	category.categoryId = *categoryId
	category.id = *id
}
