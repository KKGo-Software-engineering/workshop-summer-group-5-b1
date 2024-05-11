package transaction

import (
	"database/sql"
	"net/http"

	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/KKGo-Software-engineering/workshop-summer/api/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type handlerTransaction struct {
	flag config.FeatureFlag
	db   *sql.DB
}

const (
	cStmt = `INSERT INTO transaction (date, amount, category, transaction_type, spender_id, note, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;`
	uStmt = `UPDATE transaction SET date = $1, amount = $2, category = $3, transaction_type = $4, spender_id = $5, note = $6, image_url = $7 WHERE id = $8 RETURNING *;`
)

func NewHandler(cfg config.FeatureFlag, db *sql.DB) *handlerTransaction {
	return &handlerTransaction{cfg, db}
}

func (h handlerTransaction) Create(c echo.Context) error {

	logger := mlog.L(c)
	ctx := c.Request().Context()
	var trBody TransactionReqBody
	err := c.Bind(&trBody)
	if err != nil {
		logger.Error("bad request body", zap.Error(err))
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var insertTransaction Transaction

	err = h.db.QueryRowContext(ctx, cStmt, trBody.Date, trBody.Amount, trBody.Category, trBody.TransactionType, trBody.SpenderID, trBody.Note, trBody.ImageURL).Scan(&insertTransaction.ID, &insertTransaction.Date, &insertTransaction.Amount, &insertTransaction.Category, &insertTransaction.TransactionType, &insertTransaction.SpenderID, &insertTransaction.Note, &insertTransaction.ImageURL)

	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	logger.Info("create successfully", zap.String("id", insertTransaction.ID))
	return c.JSON(http.StatusCreated, insertTransaction)
}

func (h handler) Update(c echo.Context) error {

	logger := mlog.L(c)
	ctx := c.Request().Context()
	var trBody TransactionReqBody
	err := c.Bind(&trBody)
	if err != nil {
		logger.Error("bad request body", zap.Error(err))
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")
	var updatedTransaction Transaction
	err = h.db.QueryRowContext(ctx, uStmt, trBody.Date, trBody.Amount, trBody.Category, trBody.TransactionType, trBody.SpenderID, trBody.Note, trBody.ImageURL, id).Scan(&updatedTransaction.ID, &updatedTransaction.Date, &updatedTransaction.Amount, &updatedTransaction.Category, &updatedTransaction.TransactionType, &updatedTransaction.SpenderID, &updatedTransaction.Note, &updatedTransaction.ImageURL)

	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	logger.Info("create successfully", zap.String("id", updatedTransaction.ID))
	return c.JSON(http.StatusOK, updatedTransaction)
}
