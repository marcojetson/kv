package main

import (
    "bufio"
    "net"
    "strconv"
    "strings"
)

type Server struct {
    protocol string
    port int
    commands map[string]Command
}

func (self Server) Start() bool {
    server, err := net.Listen(self.protocol, ":" + strconv.Itoa(self.port))
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

        command, ok := self.commands[parts[0]]
        if !ok || !command.Run(conn, parts[1:]) {
            conn.Write([]byte("ERROR\n"))
        }
    }
}

func NewServer() *Server {
    return &Server{
        protocol: "tcp",
        port: 11211,
        commands: map[string]Command{},
    }
}