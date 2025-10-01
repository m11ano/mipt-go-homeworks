package ledger

import (
	"errors"
	"sync"
	"time"
)

type Transaction struct {
	ID          int
	Amount      float64
	Category    string
	Description string
	Date        time.Time
}

type Store struct {
	mu           sync.Mutex
	transactions []Transaction
}

func NewStore() *Store {
	return &Store{
		transactions: make([]Transaction, 0, 16),
	}
}

func (s *Store) AddTransaction(tx Transaction) error {
	if tx.Amount == 0 {
		return errors.New("amount must be non-zero")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	tx.ID = len(s.transactions) + 1
	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}
	s.transactions = append(s.transactions, tx)
	return nil
}

func (s *Store) ListTransactions() []Transaction {
	s.mu.Lock()
	defer s.mu.Unlock()

	cp := make([]Transaction, len(s.transactions))
	copy(cp, s.transactions)
	return cp
}
