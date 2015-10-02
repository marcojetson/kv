package core

import (
    "net"
)

type Command func (conn net.Conn, storage Storage, args []string) bool

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