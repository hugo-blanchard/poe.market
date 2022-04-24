package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"poe.market/entities"
)

type SoldItemStore struct {
	*sqlx.DB
}

func (store *SoldItemStore) Items() ([]string, error) {
	var items []string
	if error := store.Select(&items, `SELECT DISTINCT itemname FROM solditem`); error != nil {
		return []string{}, fmt.Errorf("Error getting items: %w", error)
	}
	return items, nil
}

func (store *SoldItemStore) Prices(typeName string) ([]string, error) {
	var prices []string
	if error := store.Select(&prices, `SELECT price FROM solditem WHERE itemname = $1`, typeName); error != nil {
		return []string{}, fmt.Errorf("Error getting prices: %w", error)
	}
	return prices, nil
}

func (store *SoldItemStore) CreateSoldItem(item entities.SoldItem) error {
	if _, error := store.Exec(`INSERT INTO solditem VALUES ($1, $2, $3)`, item.Id, item.Name, item.Price); error != nil {
		return fmt.Errorf("error creating solditem: %w", error)
	}
	return nil
}

func (store *SoldItemStore) DeleteSoldItem(id string) error {
	if _, error := store.Exec(`DELETE FROM solditem WHERE id = $1`, id); error != nil {
		return fmt.Errorf("error deleting solditem: %w", error)
	}
	return nil
}
