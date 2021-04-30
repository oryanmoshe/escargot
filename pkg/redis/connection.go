package redis

type Connection interface {
	Connect() error
	// SearchKeys() //With channel
	ListKeys(offset int) []Key
	Get(key string) (string, error)
	Set(key Key) error
	Delete(key string) error
	Update(key Key) (string, error)
}

type connection struct {
	Connection Connection `json:"connection"`
}
