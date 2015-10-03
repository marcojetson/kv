package server

import (
    "strconv"
)

type Command func (server Server, conn Conn, args []string) bool

func Set(server Server, conn Conn, args []string) bool {
    argc := len(args)

    if argc != 4 && argc != 5 {
        return false
    }

    flags, _ := strconv.Atoi(args[1])
    expirationTime, _ := strconv.Atoi(args[2])
    bytes, _ := strconv.Atoi(args[3])

    data, err := conn.Read()
    if err != nil {
        return false
    }

    if len(data) != bytes {
        conn.Write("CLIENT_ERROR bad data chunk")
        return false
    }

    server.Storage.Set(args[0], flags, expirationTime, []byte(data))

    if argc == 4 {
        conn.Write("STORED")
    }

    return true
}

func Get(server Server, conn Conn, args []string) bool {
    if len(args) < 1 {
        return false
    }

    for _, k := range args {
        r, ok := server.Storage.Get(k)

        if !ok {
            continue
        }

        flags := strconv.Itoa(r.Flags)
        bytes := strconv.Itoa(len(r.Data))

        conn.Write("VALUE " + k + " " + flags + " " + bytes)
        conn.Write(string(r.Data))
    }
    
    conn.Write("END")

    return true
}

func Delete(server Server, conn Conn, args []string) bool {
    argc := len(args)

    if argc != 1 && argc != 2 {
        return false
    }

    r := server.Storage.Delete(args[0])

    if argc == 1 {
        if r {
            conn.Write("DELETED")
        } else {
            conn.Write("NOT_FOUND")
        }
    }

    return true
}

func FlushAll(server Server, conn Conn, args []string) bool {
    argc := len(args)

    if argc != 0 && argc != 1 {
        return false
    }

    server.Storage.FlushAll()

    if argc == 0 {
        conn.Write("OK")
    }

    return true
}


func Version(server Server, conn Conn, args []string) bool {
    if len(args) != 0 {
         return false
     }

    conn.Write("VERSION " + server.Version)
    return true
}

func Quit(server Server, conn Conn, args []string) bool {
    if len(args) != 0 {
         return false
     }

    conn.Close()
    return true
}

func Incr(server Server, conn Conn, args []string) bool {
    if len(args) != 2 {
         return false
     }

    offset, _ := strconv.ParseUint(args[1], 10, 64)

    v, found, valid := server.Storage.Incr(args[0], offset)

    if !found {
        conn.Write("NOT FOUND")
        return true
    }

    if !valid {
        conn.Write("CLIENT_ERROR not a valid value")
        return true
    }

    conn.Write(strconv.FormatUint(v, 10))
    return true
}

func Decr(server Server, conn Conn, args []string) bool {
    if len(args) != 2 {
         return false
     }

    offset, _ := strconv.ParseUint(args[1], 10, 64)

    v, found, valid := server.Storage.Decr(args[0], offset)

    if !found {
        conn.Write("NOT FOUND")
        return true
    }

    if !valid {
        conn.Write("CLIENT_ERROR not a valid value")
        return true
    }

    conn.Write(strconv.FormatUint(v, 10))
    return true
}
