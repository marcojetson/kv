package main

import (
	"github.com/kv/kv/config"
	"github.com/kv/kv/server"
	"github.com/kv/kv/storage"
)

func main() {
	c := config.ReadConfig("kv.conf")
	s := server.NewServer(storage.MapStorage{}, c)

	s.Commands["add"] = server.Add
	s.Commands["count"] = server.Count
	s.Commands["get"] = server.Get
	s.Commands["delete"] = server.Delete
	s.Commands["deindex"] = server.DeIndex
	s.Commands["index"] = server.Index
	s.Commands["set"] = server.Set
	s.Commands["quit"] = server.Quit

	s.Start()
}
