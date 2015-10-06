package storage

type MapStorage []interface{}

func (m *MapStorage) Add(object interface{}) {
	*m = append(*m, object)
}

func (m MapStorage) Count(criteria interface{}) int {
	return 0
}

func (m MapStorage) Get(criteria interface{}) []interface{} {
	return m
}

func (m MapStorage) Delete(criteria interface{}) int {
	return 0
}

func (m MapStorage) DeIndex(key string) {
}

func (m MapStorage) Index(key string) {
}

func (m MapStorage) Set(criteria interface{}, values interface{}) int {
	return 0
}
