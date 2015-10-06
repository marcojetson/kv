package storage

type Storage []interface{}

func (m *Storage) Add(object interface{}) {
	*m = append(*m, object)
}

func (m Storage) Count(criteria interface{}) int {
	return 0
}

func (m Storage) Get(criteria interface{}) []interface{} {
	return m
}

func (m Storage) Delete(criteria interface{}) int {
	return 0
}

func (m Storage) DeIndex(key string) {
}

func (m Storage) Index(key string) {
}

func (m Storage) Set(criteria interface{}, values interface{}) int {
	return 0
}
