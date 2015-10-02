package server

import (
    "bufio"
    "net"
    "strconv"
    "strings"
    "github.com/kv/kv/core"
)

type Server struct {
    Protocol string
    Port int
    Commands map[string]core.Command
    Storage core.Storage
}

func (self Server) Start() bool {
    server, err := net.Listen(self.Protocol, ":" + strconv.Itoa(self.Port))
    if err != nil {
        panic("Failed to start")
    }

    for {   
        conn, err := server.Accept()
        if err != nil {
            continue
        }

        go self.serve(conn)
    }

    return true
}

func (self Server) serve(conn net.Conn) {
    bufr := bufio.NewReader(conn)

    for {
        line, err := bufr.ReadString('\n')
        if err != nil {
            return
        }

        line = strings.TrimSpace(line)
        parts := strings.Split(line, " ")

        command, ok := self.Commands[parts[0]]

        if !ok || !command(conn, self.Storage, parts[1:]) {
            conn.Write([]byte("ERROR\r\n"))
        }
    }
}

func NewServer(storage core.Storage) *Server {
    return &Server{
        Protocol: "tcp",
        Port: 11211,
        Commands: map[string]core.Command{},
        Storage: storage,
    }
}