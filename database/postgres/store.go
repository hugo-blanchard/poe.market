package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, error := sqlx.Open("postgres", dataSourceName)
	if error != nil {
		return nil, fmt.Errorf("error opening database %w", error)
	}

	if error := db.Ping(); error != nil {
		return nil, fmt.Errorf("error connecting to database %w", error)
	}

	return &Store{
		&ItemStore{DB: db},
		&SoldItemStore{DB: db},
	}, nil
}

type Store struct {
	ItemStore     *ItemStore
	SoldItemStore *SoldItemStore
}
