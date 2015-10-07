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
		v, ok := j.Get(k)
		if !ok {
			continue
		}

		h, _ := json.Marshal(v)
		m.indexes[k][string(h)] = append(m.indexes[k][string(h)], i)
	}
}

func (m Storage) Count(criteria Object) int {
	return 0
}

func (m Storage) Get(criteria Object) ([]Object, bool) {
	for k, _ := range criteria {
		if _, ok := m.indexes[k]; !ok {
			// invalid index found
			return nil, false
		}
	}

	r := []Object{}

	os := map[int]int{}
	for k, v := range criteria {
		h, _ := json.Marshal(v)
		is, ok := m.indexes[k][string(h)]

		if !ok || len(is) == 0 {
			// empty index, no need to continue
			return r, true
		}

		// count i ocurrences
		for _, i := range is {
			if _, ok := os[i]; !ok {
				os[i] = 0
			}

			os[i]++
		}
	}

	// keep only the ones where ocurrences equals criteria count
	e := len(criteria)
	for i, o := range os {
		if e == o {
			r = append(r, m.items[i])
		}
	}

	return r, true
}

func (m Storage) Delete(criteria Object) int {
	return 0
}

func (m Storage) DeIndex(key string) {
	delete(m.indexes, key)
}

func (m *Storage) Index(k string) {
	if _, ok := m.indexes[k]; ok {
		return
	}

	m.indexes[k] = map[string][]int{}

	for i, j := range m.items {
		v, ok := j.Get(k)
		if !ok {
			continue
		}

		h, _ := json.Marshal(v)
		m.indexes[k][string(h)] = append(m.indexes[k][string(h)], i)
	}
}

func (m Storage) Indexes() map[string]int {
	r := map[string]int{}
	for k, v := range m.indexes {
		r[k] = len(v)
	}

	return r
}

func (m Storage) Set(criteria Object, values Object) int {
	return 0
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
