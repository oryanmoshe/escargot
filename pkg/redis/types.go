package redis

import (
	"reflect"
)

//type Type interface {
//Name() string
//Value() string
//fmt.Stringer
//// string
//}

type redisType string

//type redisType struct {
//string
//}

type redisTypes struct {
	Object redisType
	String redisType
	Json   redisType
}

var Types = &redisTypes{
	// Object: redisType{"OBJECT"},
	// String: redisType{"STRING"},
	// Json:   redisType{"ReJSON-RL"},
	Object: "OBJECT",
	String: "STRING",
	Json:   "ReJSON-RL",
}

var parseMap = map[string]redisType{
	"Object":              Types.Object,
	"String":              Types.String,
	"Json":                Types.Json,
	Types.Object.String(): Types.Object,
	Types.String.String(): Types.String,
	Types.Json.String():   Types.Json,
}

var namesMap = map[redisType]string{
	Types.Object: "Object",
	Types.String: "String",
	Types.Json:   "Json",
}

func (t redisType) String() string {
	return namesMap[t]
}

func (t redisType) Name() string {
	return t.String()
}

func (t redisType) Value() string {
	return string(t)
}

func (rt redisTypes) List() []redisType {
	v := reflect.ValueOf(rt)

	values := make([]redisType, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface().(redisType)
	}

	return values
}

func (rt redisTypes) Parse(name string) redisType {
	return parseMap[name]
}
