package commands

import (
    "net"
    "strconv"
    "github.com/kv/kv/core"
)

func Set(conn net.Conn, storage core.Storage, args []string) bool {
    argc := len(args)

    if argc != 4 && argc != 5 {
        return false
    }

    flags, _ := strconv.Atoi(args[1])
    expirationTime, _ := strconv.Atoi(args[2])

    storage.Set(args[0], flags, expirationTime, []byte("data"))

    if argc == 4 {
        conn.Write([]byte("STORED\r\n"))
    }

    return true
}

func Get(conn net.Conn, storage core.Storage, args []string) bool {
    if len(args) < 1 {
        return false
    }

    for _, k := range args {
        r, ok := storage.Get(k)

        if !ok {
            continue
        }

        conn.Write([]byte("VALUE " + k + "\r\n"))
        conn.Write(r.Data)
        conn.Write([]byte("\r\n"))
    }
    
    conn.Write([]byte("END\r\n"))

    return true
}

func Delete(conn net.Conn, storage core.Storage, args []string) bool {
    argc := len(args)

    if argc != 1 && argc != 2 {
        return false
    }

    r := storage.Delete(args[0])

    if argc == 1 {
        if r {
            conn.Write([]byte("DELETED\r\n"))
        } else {
            conn.Write([]byte("NOT_FOUND\r\n"))
        }
    }

    return true
}

func FlushAll(conn net.Conn, storage core.Storage, args []string) bool {
    argc := len(args)

    if argc != 0 && argc != 1 {
        return false
    }

    storage.FlushAll()

    if argc == 0 {
        conn.Write([]byte("OK\r\n"))
    }

    return true
}


func Version(conn net.Conn, storage core.Storage, args []string) bool {
    if len(args) != 0 {
         return false
     }

    conn.Write([]byte("VERSION x\r\n"))
    return true
}

func Quit(conn net.Conn, storage core.Storage, args []string) bool {
    if len(args) != 0 {
         return false
     }

    conn.Close()
    return true
}