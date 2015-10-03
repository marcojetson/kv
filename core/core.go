package core

type Storage interface {
    Get(k string) (Value, bool)
    Delete(k string) bool
    Set(k string, flags int, expirationTime int, data []byte)
    FlushAll()
    Delta(k string, negative bool, offset uint64) (uint64, bool, bool)
    Touch(k string, expirationTime int) bool
    Append(k string, data []byte) bool
    Prepend(k string, data []byte) bool
    GarbageCollect()
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
