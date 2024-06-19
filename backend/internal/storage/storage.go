package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"backend/internal/models"
)

const MAXBIN = 99999999

type Storage struct {
	cards []string
}

func NewStorage(capacity uint) *Storage {
	s := &Storage{
		cards: make([]string, capacity+1),
	}
	return s
}

func (storage *Storage) GetIssuer(bin int) (issuer string) {
	issuer = storage.cards[bin]
	return
}

func InitStorage(pool *pgxpool.Pool) (storage *Storage, err error) {
	storage = NewStorage(MAXBIN)
	rows, err := pool.Query(context.Background(), "SELECT bin, issuer FROM beans")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c models.Cards
		err := rows.Scan(&c.Bin, &c.Issuer)
		if err != nil {
			return nil, err
		}
		storage.cards[c.Bin] = c.Issuer
	}
	rows.Close()
	return
}
