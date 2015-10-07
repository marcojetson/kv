package main

import (
	"flag"
	"github.com/kv/kv/config"
	"github.com/kv/kv/server"
)

func main() {
	f := flag.String("c", "kv.conf", "Configuration file")
	flag.Parse()

	c := config.ReadConfig(*f)
	s := server.NewServer(c)

	s.RestoreIndexes()

	s.Commands["add"] = server.Add
	s.Commands["count"] = server.Count
	s.Commands["get"] = server.Get
	s.Commands["delete"] = server.Delete
	s.Commands["deindex"] = server.DeIndex
	s.Commands["index"] = server.Index
	s.Commands["indexes"] = server.Indexes
	s.Commands["ping"] = server.Ping
	s.Commands["set"] = server.Set
	s.Commands["quit"] = server.Quit
	s.Commands["version"] = server.Version

	s.Start()
}
