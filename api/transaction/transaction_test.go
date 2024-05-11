package transaction

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	t.Run("create transaction", func(t *testing.T) {

		e := echo.New()
		defer e.Close()

		req := httptest.NewRequest(http.MethodPost, "/transactions", strings.NewReader(`{"date": "2021-08-01", "amount": 1000, "category": "food", "transaction_type": "expense", "spender_id": 1, "note": "lunch", "image_url": "http://image.com"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("error creating mock: %v", err)
		}
		defer db.Close()

		column := []string{"id", "date", "amount", "category", "transaction_type", "spender_id", "note", "image_url"}
		mock.ExpectQuery(cStmt).WithArgs("2021-08-01", 1000.0, "food", "expense", 1, "lunch", "http://image.com").WillReturnRows(sqlmock.NewRows(column).AddRow(1, "2021-08-01", 1000, "food", "expense", 1, "lunch", "http://image.com"))

		h := New(config.FeatureFlag{}, db)
		err = h.Create(c)
		if err != nil {
			t.Fatalf("error creating transaction: %v", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, `{"id": "1", "date": "2021-08-01", "amount": 1000, "category": "food", "transaction_type": "expense", "spender_id": 1, "note": "lunch", "image_url": "http://image.com"}`, rec.Body.String())
	})
}
