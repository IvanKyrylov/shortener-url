package store

type Service interface {
	Save(string) (string, error)
	Load(string) (string, error)
	LoadInfo(string) (*ItemStore, error)
	Close() error
}

type ItemStore struct {
	URL     string `json:"url"`
	Visited bool   `json:"visited"`
	Count   int    `json:"count"`
}
