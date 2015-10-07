package storage

import (
	"encoding/json"
	"strings"
)

type Storage struct {
	id      int
	items   map[int]Object
	indexes map[string]map[string][]int
}

func (m *Storage) Add(j Object) {
	m.id++
	m.items[m.id] = j

	for k, _ := range m.indexes {
		m.addToIndex(k, m.id)
	}
}

func (m Storage) Count(q Object) (int, bool) {
	is, ok := m.filter(q)

	if !ok {
		return 0, false
	}

	return len(is), true
}

func (m Storage) Get(q Object) ([]Object, bool) {
	r := []Object{}
	ids, ok := m.filter(q)

	if !ok {
		return r, false
	}

	for _, id := range ids {
		r = append(r, m.items[id])
	}

	return r, true
}

func (m *Storage) Delete(q Object) (int, bool) {
	ids, ok := m.filter(q)

	if !ok {
		return 0, false
	}

	for _, id := range ids {
		for k, _ := range m.indexes {
			m.removeFromIndex(k, id)
		}

		delete(m.items, id)
	}

	return len(ids), true
}

func (m Storage) DeIndex(key string) {
	delete(m.indexes, key)
}

func (m *Storage) Index(k string) {
	if _, ok := m.indexes[k]; ok {
		return
	}

	m.indexes[k] = map[string][]int{}

	for id, _ := range m.items {
		m.addToIndex(k, id)
	}
}

func (m Storage) Indexes() map[string]int {
	r := map[string]int{}
	for k, v := range m.indexes {
		r[k] = len(v)
	}

	return r
}

func (m Storage) Set(q Object, values Object) int {
	return 0
}

func (m *Storage) removeFromIndex(k string, id int) {
	v, ok := m.items[id].Get(k)
	if !ok {
		return
	}

	h, _ := json.Marshal(v)
	for i, t := range m.indexes[k][string(h)] {
		if id == t {
			m.indexes[k][string(h)] = append(m.indexes[k][string(h)][:i], m.indexes[k][string(h)][i+1:]...)
			return
		}
	}
}

func (m *Storage) addToIndex(k string, id int) {
	v, ok := m.items[id].Get(k)
	if !ok {
		return
	}

	h, _ := json.Marshal(v)
	m.indexes[k][string(h)] = append(m.indexes[k][string(h)], id)
}

func (m Storage) filter(q Object) ([]int, bool) {
	r := []int{}

	for k, _ := range q {
		if _, ok := m.indexes[k]; !ok {
			return r, false
		}
	}

	os := map[int]int{}
	for k, v := range q {
		h, _ := json.Marshal(v)
		ids, ok := m.indexes[k][string(h)]

		if !ok || len(ids) == 0 {
			return r, true
		}

		// count i ocurrences
		for _, id := range ids {
			if _, ok := os[id]; !ok {
				os[id] = 0
			}

			os[id]++
		}
	}

	e := len(q)
	for id, o := range os {
		if e == o {
			r = append(r, id)
		}
	}

	return r, true
}

func NewStorage() *Storage {
	return &Storage{
		id:      0,
		items:   map[int]Object{},
		indexes: map[string]map[string][]int{},
	}
}

type Object map[string]interface{}

func (o Object) Get(k string) (interface{}, bool) {
	var x interface{}

	ps := strings.Split(k, ".")

	x = o

	for _, p := range ps {
		c, valid := x.(Object)
		if !valid {
			return nil, false
		}

		v, ok := c[p]
		if !ok {
			return nil, false
		}

		x = v
	}

	return x, true
}

func NewObject(s string) (Object, bool) {
	var v Object
	err := json.Unmarshal([]byte(s), &v)
	return v, err == nil
}
