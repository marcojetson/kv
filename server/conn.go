package server

import (
    "bufio"
    "net"
    "strings"
)

type Conn struct {
    reader *bufio.Reader
    conn net.Conn
}

func (c Conn) Read() (string, error) {
    line, err := c.reader.ReadString('\n')
    line = strings.TrimSpace(line)

    return line, err
}

func (c Conn) Write(s string) {
    c.conn.Write([]byte(s + "\r\n"))
}

func (c Conn) Close() {
    c.conn.Close()
}
