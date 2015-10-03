package storage

import (
    "github.com/kv/kv/core"
    "time"
)

type MapStorage map[string]core.Value

func (m MapStorage) Get(k string) (core.Value, bool) {
    v, ok := m[k]

    if v.ExpirationTime < int(time.Now().Unix()) {
       delete(m, k)
       return core.Value{}, false
    }

    return v, ok
}

func (m MapStorage) Set(k string, flags int, expirationTime int, data []byte) {
    if expirationTime <= 60 * 60 * 24 * 30 {
       expirationTime += int(time.Now().Unix())
    }

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