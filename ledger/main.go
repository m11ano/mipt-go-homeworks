package main

import (
	"log/slog"
	"time"

	"github.com/m11ano/mipt-go-homeworks/ledger/internal/ledger"
)

func main() {
	slog.Info("Ledger service started")

	store := ledger.NewStore()

	_ = store.AddTransaction(ledger.Transaction{
		Amount:      1299.90,
		Category:    "Electronics",
		Description: "Power bank",
		Date:        time.Now().AddDate(0, 0, -1),
	})
	_ = store.AddTransaction(ledger.Transaction{
		Amount:      450.00,
		Category:    "Groceries",
		Description: "Supermarket",
		Date:        time.Now(),
	})
	_ = store.AddTransaction(ledger.Transaction{
		Amount:      199.99,
		Category:    "Transport",
		Description: "Taxi ride",
	})

	txs := store.ListTransactions()
	for _, t := range txs {
		slog.Info(
			"Transaction",
			slog.Int("ID", t.ID),
			slog.Float64("Amount", t.Amount),
			slog.String("Category", t.Category),
			slog.String("Desc", t.Description),
			slog.String("Date", t.Date.Format(time.RFC3339)),
		)
	}
}
