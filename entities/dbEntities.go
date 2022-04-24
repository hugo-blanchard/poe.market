package entities

type Item struct {
	Id      string `db:"id"`
	StashId string `db:"stash_id"`
	Name    string `db:"itemname"`
	Price   string `db:"price"`
}

type ItemStore interface {
	Items(stashId string) ([]Item, error)
	CreateItem(item Item) error
	DeleteItem(id string) error
}

type SoldItem struct {
	Id    string `db:"id"`
	Name  string `db:"itemname"`
	Price string `db:"price"`
}

type SoldItemStore interface {
	Items() ([]string, error)
	Prices(typeName string) ([]string, error)
	CreateSoldItem(item SoldItem) error
	DeleteSoldItem(id string) error
}
