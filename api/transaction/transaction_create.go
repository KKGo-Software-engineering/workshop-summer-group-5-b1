package transaction

import (
	"database/sql"
	"net/http"

	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/KKGo-Software-engineering/workshop-summer/api/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type handler struct {
	flag config.FeatureFlag
	db   *sql.DB
}

const (
	cStmt = `INSERT INTO transaction (date, amount, category, transaction_type, spender_id, note, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;`
)

func New(cfg config.FeatureFlag, db *sql.DB) *handler {
	return &handler{cfg, db}
}

func (h handler) Create(c echo.Context) error {

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
