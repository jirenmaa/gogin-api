package entity

type Person struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"gte=15,lte=200"`
	Email string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=3,max=50" validate:"test"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
