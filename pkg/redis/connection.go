package redis

// type Key struct {
// Name string
// Type RedisType

//}

type Connection interface {
	ListKeys(string) []Key
}

type connection struct {
	Connection
}
