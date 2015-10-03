package core

type Command func (storage Storage, conn Conn, args []string) bool

type Conn interface {
    Read() (string, error)
    Write(s string)
    Close()
}

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