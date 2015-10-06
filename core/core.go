package core

type Storage interface {
	Add(object interface{})
	Count(criteria interface{}) int
	Get(criteria interface{}) []interface{}
	Delete(criteria interface{}) int
	DeIndex(key string)
	Index(key string)
	Set(criteria interface{}, values interface{}) int
}

type Config interface {
	GetString(key string, def string) string
	GetInt(key string, def int) int
}
