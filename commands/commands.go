package commands

import (
    "strconv"
    "github.com/kv/kv/core"
)

func Set(storage core.Storage, conn core.Conn, args []string) bool {
    argc := len(args)

    if argc != 4 && argc != 5 {
        return false
    }

    flags, _ := strconv.Atoi(args[1])
    expirationTime, _ := strconv.Atoi(args[2])

    storage.Set(args[0], flags, expirationTime, []byte("data"))

    if argc == 4 {
        conn.Write("STORED")
    }

    return true
}

func Get(storage core.Storage, conn core.Conn, args []string) bool {
    if len(args) < 1 {
        return false
    }

    for _, k := range args {
        r, ok := storage.Get(k)

        if !ok {
            continue
        }

        conn.Write("VALUE " + k)
        conn.Write(string(r.Data))
    }
    
    conn.Write("END")

    return true
}

func Delete(storage core.Storage, conn core.Conn, args []string) bool {
    argc := len(args)

    if argc != 1 && argc != 2 {
        return false
    }

    r := storage.Delete(args[0])

    if argc == 1 {
        if r {
            conn.Write("DELETED")
        } else {
            conn.Write("NOT_FOUND")
        }
    }

    return true
}

func FlushAll(storage core.Storage, conn core.Conn, args []string) bool {
    argc := len(args)

    if argc != 0 && argc != 1 {
        return false
    }

    storage.FlushAll()

    if argc == 0 {
        conn.Write("OK")
    }

    return true
}


func Version(storage core.Storage, conn core.Conn, args []string) bool {
    if len(args) != 0 {
         return false
     }

    conn.Write("VERSION x")
    return true
}

func Quit(storage core.Storage, conn core.Conn, args []string) bool {
    if len(args) != 0 {
         return false
     }

    conn.Close()
    return true
}