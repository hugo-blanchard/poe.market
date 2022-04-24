package database

import (
	"fmt"
	"log"

	"poe.market/config"
	"poe.market/database/postgres"
	"poe.market/entities"
)

// This is where you would swap the store instantiation to an other database of your choice
func NewStore(config config.PgConfig) *Store {
	connectionString := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		config.DbName,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.SslMode,
	)
	log.Print(connectionString)
	store, err := postgres.NewStore(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		ItemStore:     store.ItemStore,
		SoldItemStore: store.SoldItemStore,
	}
}

type Store struct {
	ItemStore     entities.ItemStore
	SoldItemStore entities.SoldItemStore
}
