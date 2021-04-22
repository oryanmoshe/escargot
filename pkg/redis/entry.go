package redis

import "fmt"

type Key struct {
	Name string
	Type redisType
}

type Value interface {
	fmt.Stringer
}

type Entry interface {
	Value() Value
}

type stringValue string

type entry struct {
	Entry
	Key   Key
	value Value
}

func (e entry) Value() Value {
	return e.value.String()
}

//type Connection interface {
////ListKeys(string) []Key
//}

// type connection struct{}
