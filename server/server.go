package server

import (
	"bufio"
	"encoding/json"
	"github.com/kv/kv/config"
	"github.com/kv/kv/storage"
	"io/ioutil"
	"net"
	"path"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	Protocol     string
	Port         int
	DumpInterval time.Duration
	Path         string
	Commands     map[string]Command
	Storage      *storage.Storage
	Version      string
}

func (s Server) Start() bool {
	server, err := net.Listen(s.Protocol, ":"+strconv.Itoa(s.Port))
	if err != nil {
		panic("Failed to start " + err.Error())
	}

	if s.DumpInterval > 0 {
		go func() {
			// @TODO implement
		}()
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}

		go s.serve(Conn{
			reader: bufio.NewReader(conn),
			conn:   conn,
		})
	}

	return true
}

func (s Server) RestoreIndexes() {
	f, err := ioutil.ReadFile(path.Join(s.Path, file_indexes))
	if err != nil {
		return
	}

	var ks []string
	json.Unmarshal(f, &ks)
	for _, k := range ks {
		s.Storage.Index(k)
	}
}

func (s Server) DumpIndexes() {
	ks := []string{}

	for k, _ := range s.Storage.Indexes() {
		ks = append(ks, k)
	}

	j, _ := json.Marshal(ks)
	ioutil.WriteFile(path.Join(s.Path, file_indexes), j, 0644)
}

func (s Server) serve(conn Conn) {
	for {
		line, err := conn.Read()
		if err != nil {
			return
		}

		parts := strings.Split(line, " ")

		command, ok := s.Commands[parts[0]]

		if !ok || !command(s, conn, parts[1:]) {
			conn.Write(response_error)
		}
	}
}

func NewServer(config config.Config) *Server {
	return &Server{
		Protocol:     config.GetString("protocol", "tcp"),
		Port:         config.GetInt("port", 11211),
		DumpInterval: time.Duration(config.GetInt("dump", 0)) * time.Second,
		Path:         config.GetString("path", "/var/data/kv"),
		Commands:     map[string]Command{},
		Storage:      storage.NewStorage(),
		Version:      "0.1b",
	}
}
