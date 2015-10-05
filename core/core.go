package core

type Storage interface {
	Add()
	Count()
	Get()
	Delete()
	DeIndex()
	Index()
	Set()
}

type Config interface {
	GetString(key string, def string) string
	GetInt(key string, def int) int
}
