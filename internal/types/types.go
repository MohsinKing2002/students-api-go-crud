package types

type Student struct{
	ID    int    `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"required,min=1"`
}