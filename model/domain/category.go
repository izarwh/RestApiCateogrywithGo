package domain

import "prrestapi/model/response"

type Category struct {
	id   int
	name string
}

func (Category *Category) ToResponseCategory() response.ResponseCategory {
	return response.ResponseCategory{
		Id:   Category.id,
		Name: Category.name,
	}
}

func (c *Category) GetId() *int {
	return &c.id
}

func (c *Category) GetName() *string {
	return &c.name
}

func (c *Category) GetCategory() *Category {
	return &Category{c.id, c.name}
}

func (c *Category) SetId(id *int) {
	c.id = *id
}

func (c *Category) SetName(name *string) {
	c.name = *name
}

func (c *Category) SetCategory(id *int, name *string) {
	c.id = *id
	c.name = *name
}
