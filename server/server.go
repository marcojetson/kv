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
    xconn := Conn{
        buffer: bufio.NewReader(conn),
        conn: conn,
    }

    for {
        line, err := xconn.Read()
        if err != nil {
            return
        }

        line = strings.TrimSpace(line)
        parts := strings.Split(line, " ")

        command, ok := self.Commands[parts[0]]

        if !ok || !command(self.Storage, xconn, parts[1:]) {
            xconn.Write("ERROR")
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

type Conn struct {
    buffer *bufio.Reader
    conn net.Conn
}

func (c Conn) Read() (string, error) {
    return c.buffer.ReadString('\n')
}

func (c Conn) Write(s string) {
    c.conn.Write([]byte(s + "\r\n"))
}

func (c Conn) Close() {
    c.conn.Close()
}