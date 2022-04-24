package streamreader

import (
	"log"

	"poe.market/config"
	"poe.market/database"
	"poe.market/entities"
)

func writeToDbRoutine(pageChan <-chan StashPage, config config.ConfigFile) {
	store := database.NewStore(config.PgConfig)

	for {
		page := <-pageChan

		writePageToDb(store, listToMap(config.Items), page)
	}
}

func listToMap(list []string) map[string]bool {
	result := make(map[string]bool, len(list))
	for _, entry := range list {
		result[entry] = true
	}
	return result
}

func writePageToDb(store *database.Store, itemsToTrack map[string]bool, page StashPage) {
	for _, stash := range page.Stashes {
		oldItems, err := store.ItemStore.Items(stash.ID)
		if err != nil {
			log.Fatal(err)
		}

		newItems := filterImportantItems(itemsToTrack, stash.Items)

		oldItemsNotInNewStash, newItemsNotInOldStash := interpolateStashes(getMapOfDbStash(oldItems), getMapOfJsonStash(newItems))

		for _, item := range oldItemsNotInNewStash {
			log.Printf("Registered a %s to the sold db", item.Name)
			store.SoldItemStore.CreateSoldItem(entities.SoldItem{
				Id:    item.Id,
				Name:  item.Name,
				Price: item.Price,
			})
			store.ItemStore.DeleteItem(item.Id)
		}

		for _, item := range newItemsNotInOldStash {
			log.Printf("Registered a %s to the standard db", item.Name)
			store.ItemStore.CreateItem(entities.Item{
				Id:      item.Id,
				StashId: stash.ID,
				Name:    item.Name,
				Price:   item.Price,
			})
		}
	}
}

func filterImportantItems(itemsToTrack map[string]bool, items []Item) []Item {
	var importantItems []Item

	for _, item := range items {
		if _, contains := itemsToTrack[item.Name]; !contains {
			continue
		}

		if item.Price == "" {
			continue
		}

		importantItems = append(importantItems, item)
	}

	return importantItems
}

func getMapOfDbStash(stash []entities.Item) map[string]entities.Item {
	mapStash := make(map[string]entities.Item)

	for _, item := range stash {
		mapStash[item.Id] = item
	}

	return mapStash
}

func getMapOfJsonStash(stash []Item) map[string]Item {
	mapStash := make(map[string]Item)

	for _, item := range stash {
		mapStash[item.Id] = item
	}

	return mapStash
}

func interpolateStashes(oldStash map[string]entities.Item, newStash map[string]Item) ([]entities.Item, []Item) {
	var oldItems []entities.Item
	var newItems []Item

	for itemId, item := range oldStash {
		if _, present := newStash[itemId]; !present {
			oldItems = append(oldItems, item)
		}
	}

	for itemId, item := range newStash {
		if _, present := oldStash[itemId]; !present {
			newItems = append(newItems, item)
		}
	}

	return oldItems, newItems
}
