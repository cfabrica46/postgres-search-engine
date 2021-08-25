package structure

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Brand string `json:"brand"`
	Date  string `json:"date"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
}

type Token struct {
	Content string `header:"Authorization"`
}
