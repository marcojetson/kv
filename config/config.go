package config

import (
    "bufio"
    "os"
    "strings"
    "strconv"
)

func ReadConfig(name string) Config {
    r := map[string](string){}

    file, err := os.Open(name)
    if err != nil {
        return r
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        vals := strings.SplitN(line, "=", 2)
        key := strings.TrimSpace(vals[0])
        value := strings.TrimSpace(vals[1])
        r[key] = value
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return r
}

type Config map[string]string

func (c Config) GetString(k string, def string) string {
    v, ok := c[k]
    
    if !ok {
        return def
    }

    return v
}

func (c Config) GetInt(k string, def int) int {
    v, ok := c[k]

    if !ok {
       return def
    }
    
    if i, err := strconv.Atoi(v); err == nil {
       return i
    }

    return def
}