package food

type Food struct {
	ID          string `json:"food_id" db:"food_id"`
	FoodName    string `json:"food_name" db:"food_name"`
	ShortDesc   string `json:"short_desc" db:"short_desc"`
	LongDesc    string `json:"long_desc" db:"long_desc"`
	Price       int    `json:"price" db:"price"`
	IsAvailable int    `json:"is_availability" db:"is_availability"`
}
