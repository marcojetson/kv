package main

import (
    "github.com/kv/kv/config"
    "github.com/kv/kv/server"
    "github.com/kv/kv/storage"
)

func main() {
    c := config.ReadConfig("kv.conf")
    s := server.NewServer(storage.MapStorage{}, c)

    s.Commands["set"]  = server.Set
    s.Commands["get"]  = server.Get
    s.Commands["gets"]  = server.Get
    s.Commands["delete"]  = server.Delete
    s.Commands["flush_all"]  = server.FlushAll
    s.Commands["version"] = server.Version
    s.Commands["quit"] = server.Quit
    s.Commands["incr"] = server.Incr
    s.Commands["decr"] = server.Decr

    s.Start()
}

