package customer

type Customer struct {
	ID        string `json:"customer_id" db:"customer_id"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
	AddressID string `json:"address_id" db:"address_id"`
	Mobile    string `json:"mobile" db:"mobile"`
	Email     string `json:"email" db:"email"`
}

type Address struct {
	ID      string `json:"address_id" db:"address_id"`
	NO      string `json:"no" db:"no"`
	Line_1  string `json:"line_1" db:"line_1"`
	Line_2  string `json:"line_2" db:"line_2"`
	City    string `json:"city" db:"city"`
	State   string `json:"state" db:"state"`
	Country string `json:"country" db:"country"`
	ZIP     string `json:"zip" db:"zip"`
	Type    string `json:"type" db:"type"`
}
