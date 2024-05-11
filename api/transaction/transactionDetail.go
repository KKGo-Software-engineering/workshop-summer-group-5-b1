package transaction

import (
	"net/http"

	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/KKGo-Software-engineering/workshop-summer/api/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type handler struct {
	flag config.FeatureFlag
	db   TxDetailStorer
}

type TxDetailStorer interface {
	GetTransactionDetailBySpenderId(id string) (TransactionWithDetail, error)
}

func New(cfg config.FeatureFlag, db TxDetailStorer) *handler {
	return &handler{cfg, db}
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

func (h handler) GetTransactionDetailBySpenderIdHandler(c echo.Context) error {

	logger := mlog.L(c)
	//ctx := c.Request().Context()

	id := c.Param("id")

	txDetail, err := h.db.GetTransactionDetailBySpenderId(id)
	if err != nil {
		logger.Error("query error", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, txDetail)

}

func GetTransactionDetailBySpenderId(id string) (TransactionWithDetail, error) {
	return TransactionWithDetail{
		Transactions: []Transaction{
			{
				ID:              "1",
				Date:            "2024-04-30T09:00:00.000Z",
				Amount:          1000,
				Category:        "Food",
				TransactionType: "expense",
				SpenderID:       1,
				Note:            "Lunch",
				ImageURL:        "https://example.com/image1.jpg",
			},
			{
				ID:              "2",
				Date:            "2024-04-29T19:00:00.000Z",
				Amount:          2000,
				Category:        "Transport",
				TransactionType: "income",
				SpenderID:       1,
				Note:            "Salary",
				ImageURL:        "https://example.com/image2.jpg",
			},
		},
		Summary: TransactionSummary{
			TotalIncome:    2000,
			TotalExpenses:  1000,
			CurrentBalance: 1000,
		},
		Pagination: PaginationInfo{
			CurrentPage: 1,
			TotalPages:  1,
			PerPage:     10,
		},
	}, nil
}
