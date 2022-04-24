package streamreader

import (
	"encoding/json"
	"log"
	"net/http"
)

func getNextChangeId() string {
	resp, err := http.Get("https://poe.ninja/api/data/GetStats?")
	if err != nil {
		log.Fatalln(err)
	}

	var stats poeNinjaStats

	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		log.Fatalln(err)
	}

	return stats.NextChangeID
}

type poeNinjaStats struct {
	NextChangeID string `json:"next_change_id"`
	// ID                       int    `json:"id"`
	// APIBytesDownloaded       int64  `json:"api_bytes_downloaded"`
	// StashTabsProcessed       int64  `json:"stash_tabs_processed"`
	// APICalls                 int    `json:"api_calls"`
	// CharacterBytesDownloaded int64  `json:"character_bytes_downloaded"`
	// CharacterAPICalls        int    `json:"character_api_calls"`
	// LadderBytesDownloaded    int64  `json:"ladder_bytes_downloaded"`
	// LadderAPICalls           int    `json:"ladder_api_calls"`
	// PobCharactersCalculated  int    `json:"pob_characters_calculated"`
}
