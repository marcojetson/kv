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
    Version string
    Commands map[string]core.Command
    Storage core.Storage
}

func (s Server) Start() bool {
    server, err := net.Listen(s.Protocol, ":" + strconv.Itoa(s.Port))
    if err != nil {
        panic("Failed to start")
    }

    for {   
        conn, err := server.Accept()
        if err != nil {
            continue
        }

        go s.serve(Conn{
            reader: bufio.NewReader(conn),
            conn: conn,
        })
    }

    return true
}

func (s Server) serve(conn Conn) {
    for {
        line, err := conn.Read()
        if err != nil {
            return
        }

        parts := strings.Split(line, " ")

        command, ok := s.Commands[parts[0]]

        if !ok || !command(s.Storage, conn, parts[1:]) {
            conn.Write("ERROR")
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
