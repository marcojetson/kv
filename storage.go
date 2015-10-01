package main

type key string

type Storage interface {
  Get(k key) []byte
  GetAll(ks []key) map[key]([]byte)
  Delete(k key)
  Set(k key, v []byte)
  FlushAll()
}

type MapStorage map[key]([]byte)

func (m MapStorage) Get(k key) []byte{
  i, _ := m[k]
  return i
}

func (m MapStorage) GetAll(ks []key) map[key]([]byte) {
  r := map[key]([]byte){}
  for _, k := range ks {
    r[k] = m.Get(k)
  }
  return r
}

func (m MapStorage) Delete(k key) {
  delete(m, k)
}

func (m MapStorage) Set(k key, v []byte) {
  m[k] = v
}

func (m MapStorage) FlushAll() {
     return
}
