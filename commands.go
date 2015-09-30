package main

import (
    "net"
)

type Command interface {
    Run(conn net.Conn, args []string) bool
}

type Quit struct {
}

func (self Quit) Run(conn net.Conn, args []string) bool {
    conn.Close()
    return true
}

type Echo struct {
}

func (self Echo) Run(conn net.Conn, args []string) bool {
    if len(args) != 1 {
        return false
    }

    conn.Write([]byte(args[0] + "\n"))

    return true
}