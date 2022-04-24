package streamreader

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func readRoutine(pageChan chan<- StashPage) {
	nextId := getNextChangeId()

	var page StashPage

	for {
		page = readStream(nextId)
		pageChan <- page

		if len(page.Stashes) == 0 {
			log.Println("Caught up to the head of the stream, sleeping for 10 seconds.")
			time.Sleep(10 * time.Second)
		} else {
			nextId = page.NextChangeID
		}
	}
}

func readStream(stashId string) StashPage {
	resp, err := http.Get(fmt.Sprintf("https://pathofexile.com/api/public-stash-tabs?id=%s", stashId))
	if err != nil {
		log.Fatalln(err)
	}

	var page StashPage

	err = json.NewDecoder(resp.Body).Decode(&page)
	if err != nil {
		log.Fatalln(err)
	}

	return page
}
