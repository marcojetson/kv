package main

import (
    "kv/commands"
    "kv/server"
    "kv/storage"
)

func main() {
    s := server.NewServer(storage.MapStorage{})

    s.Commands["set"]  = commands.Set
    s.Commands["get"]  = commands.Get
    s.Commands["gets"]  = commands.Get
    s.Commands["delete"]  = commands.Delete
    s.Commands["flush_all"]  = commands.FlushAll
    s.Commands["version"] = commands.Version
    s.Commands["quit"] = commands.Quit

    s.Start()
}
