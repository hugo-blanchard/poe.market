package streamreader

type StashPage struct {
	NextChangeID string  `json:"next_change_id"`
	Stashes      []Stash `json:"stashes"`
}

type Stash struct {
	ID          string `json:"id"`
	AccountName string `json:"accountName"`
	League      string `json:"league"`
	Items       []Item `json:"items"`
}

type Item struct {
	Price      string `json:"note"`
	Verified   bool   `json:"verified"`
	Identified bool   `json:"identified"`
	Ilvl       int    `json:"ilvl"`
	Name       string `json:"name"`
	Id         string `json:"id"`
}
