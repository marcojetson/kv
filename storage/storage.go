package storage

type MapStorage map[string]([]byte)

func (m MapStorage) Get(k string) ([]byte, bool) {
    v, ok := m[k]
    return v, ok
}

func (m MapStorage) Set(k string, v []byte) {
    m[k] = v
}

func (m MapStorage) Delete(k string) bool {
    _, ok := m[k]
    delete(m, k)
    return ok
}

func (m MapStorage) FlushAll() {
    for k := range m {
        delete(m, k)
    }
}