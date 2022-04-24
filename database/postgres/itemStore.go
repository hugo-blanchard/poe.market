package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"poe.market/entities"
)

type ItemStore struct {
	*sqlx.DB
}

func (store *ItemStore) Items(stashId string) ([]entities.Item, error) {
	var items []entities.Item
	if error := store.Select(&items, `SELECT * FROM item WHERE stash_id = $1`, stashId); error != nil {
		return []entities.Item{}, fmt.Errorf("error getting items: %w", error)
	}
	return items, nil
}

func (store *ItemStore) CreateItem(item entities.Item) error {
	if _, error := store.Exec(`INSERT INTO item VALUES ($1, $2, $3, $4)`, item.Id, item.StashId, item.Name, item.Price); error != nil {
		return fmt.Errorf("error creating item: %w", error)
	}
	return nil
}

func (store *ItemStore) DeleteItem(id string) error {
	if _, error := store.Exec(`DELETE FROM item WHERE id = $1`, id); error != nil {
		return fmt.Errorf("error deleting item: %w", error)
	}
	return nil
}
