package entities

// Book represents a single book
type Book struct {
	ID          int     `json:"-"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	ISBN        string  `json:"isbn"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"gte=5"`
}
