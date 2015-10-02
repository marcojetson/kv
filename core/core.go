package core

import (
    "net"
)

type Command func (conn net.Conn, storage Storage, args []string) bool

type Storage interface {
    Get(k string) ([]byte, bool)
    Delete(k string) bool
    Set(k string, v []byte)
    FlushAll()
}
