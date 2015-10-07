package storage

import (
	"encoding/json"
	"strings"
)

type Storage struct {
	items   []Object
	indexes map[string]map[string][]int
}

func (m *Storage) Add(j Object) {
	i := len(m.items)
	m.items = append(m.items, j)

	for k, _ := range m.indexes {
		m.addToIndex(k, i)
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
	is, ok := m.filter(q)

	if !ok {
		return r, false
	}

	for _, i := range is {
		r = append(r, m.items[i])
	}

	return r, true
}

func (m *Storage) Delete(q Object) (int, bool) {
	is, ok := m.filter(q)

	if !ok {
		return 0, false
	}

	// @TODO delete

	return len(is), true
}

func (m Storage) DeIndex(key string) {
	delete(m.indexes, key)
}

func (m *Storage) Index(k string) {
	if _, ok := m.indexes[k]; ok {
		return
	}

	m.indexes[k] = map[string][]int{}

	for i, _ := range m.items {
		m.addToIndex(k, i)
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

func addToIndex(k string, j Object) {

}

func (m *Storage) addToIndex(k string, i int) {
	v, ok := m.items[i].Get(k)
	if !ok {
		return
	}

	h, _ := json.Marshal(v)
	m.indexes[k][string(h)] = append(m.indexes[k][string(h)], i)
}

func (m Storage) filter(q Object) ([]int, bool) {
	is := []int{}

	for k, _ := range q {
		if _, ok := m.indexes[k]; !ok {
			return is, false
		}
	}

	os := map[int]int{}
	for k, v := range q {
		h, _ := json.Marshal(v)
		is, ok := m.indexes[k][string(h)]

		if !ok || len(is) == 0 {
			return is, true
		}

		// count i ocurrences
		for _, i := range is {
			if _, ok := os[i]; !ok {
				os[i] = 0
			}

			os[i]++
		}
	}

	e := len(q)
	for i, o := range os {
		if e == o {
			is = append(is, i)
		}
	}

	return is, true
}

func NewStorage() *Storage {
	return &Storage{
		items:   []Object{},
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
