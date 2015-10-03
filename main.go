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
    s.Commands["add"]  = server.Add
    s.Commands["replace"]  = server.Replace
    s.Commands["append"] = server.Append
    s.Commands["prepend"] = server.Prepend
    s.Commands["get"]  = server.Get
    s.Commands["gets"]  = server.Get
    s.Commands["delete"]  = server.Delete
    s.Commands["flush_all"]  = server.FlushAll
    s.Commands["version"] = server.Version
    s.Commands["quit"] = server.Quit
    s.Commands["incr"] = server.Incr
    s.Commands["decr"] = server.Decr
    s.Commands["touch"] = server.Touch

    s.Start()
}

