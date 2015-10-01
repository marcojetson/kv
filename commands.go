package main

import (
    "net"
)

type Command interface {
    Run(conn net.Conn, storage Storage, args []string) bool
}

type Quit struct {
}

func (self Quit) Run(conn net.Conn, storage Storage, args []string) bool {
    conn.Close()
    return true
}

type Echo struct {
}

func (self Echo) Run(conn net.Conn, storage Storage, args []string) bool {
    if len(args) != 1 {
        return false
    }

    conn.Write([]byte(args[0] + "\n"))

    return true
}

type Set struct {
}

func (self Set) Run(conn net.Conn, storage Storage, args []string) bool {
     if len(args) != 2 {
     	return false
     }     

     storage.Set(key(args[0]), []byte(args[1]))

     return true
}

type Get struct {
}

func (self Get) Run(conn net.Conn, storage Storage, args []string) bool {
     if len(args) != 1 {
     	return false
     }

     r := storage.Get(key(args[0]))

     conn.Write(r)
     conn.Write([]byte("\n"))

     return true
}
