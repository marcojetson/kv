package core

type Storage interface {
    Get(k string) (Value, bool)
    Delete(k string) bool
    Set(k string, flags int, expirationTime int, data []byte)
    FlushAll()
}

type Value struct {
    Flags int
    ExpirationTime int
    Data []byte
}

type Config interface {
    GetString(key string, def string) string
    GetInt(key string, def int) int
}