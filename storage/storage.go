package storage

import (
    "github.com/kv/kv/core"
)

type MapStorage map[string]core.Value

func (m MapStorage) Get(k string) (core.Value, bool) {
    v, ok := m[k]
    return v, ok
}

func (m MapStorage) Set(k string, flags int, expirationTime int, data []byte) {
    m[k] = core.Value{
    	Flags: flags,
    	ExpirationTime: expirationTime,
    	Data: data,
    }
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