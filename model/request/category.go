package request

type RequestCreateCategory struct {
	Name       string `json:"name" validate:"required,min=1,max=200" message:"name is required"`
	CategoryId int    `json:"category_id" validate:"required,numeric" message:"id is required"`
	// Name       string `json:"name"`
	// CategoryId int    `json:"category_id"`
}

type RequestDeleteCategory struct {
	Id int `json:"id" validate:"required,numeric" message:"id is required"`
}

type RequestDeleteByCategoryId struct {
	CategoryId int `json:"category_id" validate:"required,numeric" message:"id is required"`
}

type RequestUpdateCategory struct {
	Id         int    `json:"id" validate:"required,numeric"`
	CategoryId int    `json:"category_id" validate:"required,numeric"`
	Name       string `json:"name" validate:"required,min=1,max=200"`
}
