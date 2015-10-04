package storage

import (
	"github.com/kv/kv/core"
	"strconv"
	"time"
)

type MapStorage map[string]core.Value

func (m MapStorage) Get(k string) (core.Value, bool) {
	v, ok := m[k]

	if v.ExpirationTime != 0 && v.ExpirationTime < int(time.Now().Unix()) {
		delete(m, k)
		return core.Value{}, false
	}

	return v, ok
}

func (m MapStorage) Set(k string, flags int, expirationTime int, data []byte) {
	if expirationTime > 0 && expirationTime <= 60*60*24*30 {
		expirationTime += int(time.Now().Unix())
	}

	m[k] = core.Value{
		Flags:          flags,
		ExpirationTime: expirationTime,
		Data:           data,
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

func (m MapStorage) Delta(k string, negative bool, offset uint64) (uint64, bool, bool) {
	v, ok := m.Get(k)

	if !ok {
		return 0, false, false
	}

	data := v.Data

	i, err := strconv.ParseUint(string(data), 10, 64)

	if err != nil {
		return 0, true, false
	}

	if negative {
		i -= offset
	} else {
		i += offset
	}

	v.Data = []byte(strconv.FormatUint(i, 10))

	m[k] = v

	return i, true, true
}

func (m MapStorage) Touch(k string, expirationTime int) bool {
	v, ok := m.Get(k)

	if !ok {
		return false
	}

	if expirationTime > 0 && expirationTime <= 60*60*24*30 {
		expirationTime += int(time.Now().Unix())
	}

	v.ExpirationTime = expirationTime

	m[k] = v

	return true
}

func (m MapStorage) Append(k string, data []byte) bool {
	v, ok := m.Get(k)

	if !ok {
		return false
	}

	v.Data = append(v.Data, data...)

	m[k] = v

	return true
}

func (m MapStorage) Prepend(k string, data []byte) bool {
	v, ok := m.Get(k)

	if !ok {
		return false
	}

	v.Data = append(data, v.Data...)

	m[k] = v

	return true
}

func (m MapStorage) GarbageCollect() {
	for k := range m {
		m.Get(k)
	}
}
