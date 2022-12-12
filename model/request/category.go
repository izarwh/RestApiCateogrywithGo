package request

type RequestCreateCategory struct {
	Name string `json:"name" validate:"required,min=1,max=200" message:"name is required"`
}

type RequestDeleteCategory struct {
	Id int `json:"id" validate:"required,numeric" message:"id is required"`
}

type RequestUpdateCategory struct {
	Id   int    `json:"id" validate:"required,numeric"`
	Name string `json:"name" validate:"required,min=1,max=200"`
}
