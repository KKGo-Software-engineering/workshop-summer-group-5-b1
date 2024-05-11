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

/*
{
	"transactions": [
		{
			"id": 1,
			"date": "2024-04-30T09:00:00.000Z",
			"amount": 1000,
			"category": "Food",
			"transaction_type": "expense",
			"spender_id": 1,
			"note": "Lunch",
			"image_url": "https://example.com/image1.jpg"
		},
		{
			"id": 2,
			"date": "2024-04-29T19:00:00.000Z",
			"amount": 2000,
			"category": "Transport",
			"transaction_type": "income",
			"spender_id": 1,
			"note": "Salary",
			"image_url": "https://example.com/image2.jpg"
		}
	],
	"summary": {
		"total_income": 2000,
		"total_expenses": 1000,
		"current_balance": 1000
	},
	"pagination": {
		"current_page": 1,
		"total_pages": 1,
		"per_page": 10
	}
}
*/

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
