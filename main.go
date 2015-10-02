package main

import (
    "github.com/kv/kv/commands"
    "github.com/kv/kv/server"
    "github.com/kv/kv/storage"
    "github.com/kv/kv/config"
)

func main() {
    c := config.ReadConfig("config.kv")
    s := server.NewServer(storage.MapStorage{}, c)

    s.Commands["set"]  = commands.Set
    s.Commands["get"]  = commands.Get
    s.Commands["gets"]  = commands.Get
    s.Commands["delete"]  = commands.Delete
    s.Commands["flush_all"]  = commands.FlushAll
    s.Commands["version"] = commands.Version
    s.Commands["quit"] = commands.Quit

    s.Start()
}

