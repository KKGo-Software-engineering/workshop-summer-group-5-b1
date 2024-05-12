package transaction

type Transaction struct {
	ID              string  `json:"id"`
	Date            string  `json:"date"`
	Amount          float64 `json:"amount"`
	Category        string  `json:"category"`
	TransactionType string  `json:"transaction_type"`
	SpenderID       int     `json:"spender_id"`
	Note            string  `json:"note"`
	ImageURL        string  `json:"image_url"`
}

type TransactionSummary struct {
	TotalIncome    float64 `json:"total_income"`
	TotalExpenses  float64 `json:"total_expenses"`
	CurrentBalance float64 `json:"current_balance"`
}

type PaginationInfo struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
}

type TransactionWithDetail struct {
	Transactions []Transaction      `json:"transactions"`
	Summary      TransactionSummary `json:"summary"`
	Pagination   PaginationInfo     `json:"pagination"`
}

type TransactionReqBody struct {
	Date            string  `json:"date"`
	Amount          float64 `json:"amount"`
	Category        string  `json:"category"`
	TransactionType string  `json:"transaction_type"`
	SpenderID       int     `json:"spender_id"`
	Note            string  `json:"note"`
	ImageURL        string  `json:"image_url"`
}
