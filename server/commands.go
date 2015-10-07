package server

import (
	"encoding/json"
	"fmt"
	"github.com/kv/kv/storage"
	"strconv"
)

type Command func(server Server, conn Conn, args []string) bool

func Add(server Server, conn Conn, args []string) bool {
	if len(args) != 1 {
		return false
	}

	object, ok := storage.NewObject(args[0])

	if !ok {
		return false
	}

	server.Storage.Add(object)
	conn.Write(response_success)
	return true
}

func Count(server Server, conn Conn, args []string) bool {
	if len(args) != 1 {
		return false
	}

	q, ok := storage.NewObject(args[0])

	if !ok {
		return false
	}

	c, valid := server.Storage.Count(q)

	if !valid {
		conn.Write(response_error_no_index)
		return true
	}

	conn.Write(fmt.Sprintf(response_success_arg, strconv.Itoa(c)))
	return true
}

func Get(server Server, conn Conn, args []string) bool {
	if len(args) != 1 {
		return false
	}

	q, ok := storage.NewObject(args[0])

	if !ok {
		return false
	}

	items, valid := server.Storage.Get(q)

	if !valid {
		conn.Write(response_error_no_index)
		return true
	}

	for _, item := range items {
		b, err := json.Marshal(item)

		if err != nil {
			continue
		}

		conn.Write(string(b))
	}

	conn.Write(response_end)
	return true
}

func Delete(server Server, conn Conn, args []string) bool {
	if len(args) != 1 {
		return false
	}

	q, ok := storage.NewObject(args[0])

	if !ok {
		return false
	}

	c, valid := server.Storage.Delete(q)

	if !valid {
		conn.Write(response_error_no_index)
		return true
	}

	conn.Write(fmt.Sprintf(response_success_arg, strconv.Itoa(c)))
	return true
}

func DeIndex(server Server, conn Conn, args []string) bool {
	argc := len(args)

	if argc != 1 && argc != 2 {
		return false
	}

	if argc == 2 {
		go server.Storage.DeIndex(args[0])
		conn.Write(response_queued)
	} else {
		server.Storage.DeIndex(args[0])
		conn.Write(response_success)
	}

	server.DumpIndexes()

	return true
}

func Index(server Server, conn Conn, args []string) bool {
	argc := len(args)

	if argc != 1 && argc != 2 {
		return false
	}

	if argc == 2 {
		go server.Storage.Index(args[0])
		conn.Write(response_queued)
	} else {
		server.Storage.Index(args[0])
		conn.Write(response_success)
	}

	server.DumpIndexes()

	return true
}

func Indexes(server Server, conn Conn, args []string) bool {
	if len(args) != 0 {
		return false
	}

	for index, size := range server.Storage.Indexes() {
		conn.Write("KEY " + index + " " + strconv.Itoa(size))
	}

	conn.Write(response_end)
	return true
}

func Quit(server Server, conn Conn, args []string) bool {
	if len(args) != 0 {
		return false
	}

	conn.Close()
	return true
}

func Ping(server Server, conn Conn, args []string) bool {
	if len(args) != 0 {
		return false
	}

	conn.Write("PONG")
	return true
}

func Set(server Server, conn Conn, args []string) bool {
	if len(args) != 2 {
		return false
	}

	q, ok := storage.NewObject(args[0])

	if !ok {
		return false
	}

	values, ok2 := storage.NewObject(args[1])

	if !ok2 {
		return false
	}

	server.Storage.Set(q, values)

	return true
}

func Version(server Server, conn Conn, args []string) bool {
	if len(args) != 0 {
		return false
	}

	conn.Write("VERSION " + server.Version)
	return true
}
