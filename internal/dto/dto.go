package dto

type ProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
