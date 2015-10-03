package main

import (
    "strconv"
    "testing"
    "time"
    "github.com/bradfitz/gomemcache/memcache"
)

var mc *memcache.Client

func TestConnect(t *testing.T) {
    go main()

    mc = memcache.New("127.0.0.1:11211")
}

func TestGetNotFound(t *testing.T) {
    _, err := mc.Get("non_existant_key")

    if err == nil {
        t.Fatalf("Expecting error retrieving non_existant_key")
    }
}

func TestSet(t *testing.T) {
    err := mc.Set(&memcache.Item{Key: "foo", Value: []byte("bar")})

    if err != nil {
        t.Fatalf("Set foo = bar returned error " + err.Error())
    }
}

func TestGet(t *testing.T) {
    i, err := mc.Get("foo")

    if err != nil {
        t.Fatalf("Get foo returned error " + err.Error())
    }

    v := string(i.Value)

    if v != "bar" {
        t.Fatalf("Expecting foo to be bar but " + v + " found")
    }
}

func TestDelete(t *testing.T) {
    mc.Delete("foo")
    _, err := mc.Get("foo")

    if err == nil {
        t.Fatalf("Expecting error retrieving foo after delete")
    }
}

func TestSetExpiration(t *testing.T) {
    err := mc.Set(&memcache.Item{Key: "foo", Value: []byte("bar"), Expiration: 2})

    if err != nil {
        t.Fatalf("Set expiration foo = bar returned error " + err.Error())
    }
}

func TestGetExpiration(t *testing.T) {
    i, err := mc.Get("foo")

    if err != nil {
        t.Fatalf("Get foo returned error " + err.Error())
    }

    v := string(i.Value)

    if v != "bar" {
        t.Fatalf("Expecting foo to be bar but " + v + " found")
    }
}

func TestGetExpired(t *testing.T) {
    time.Sleep(3 * time.Second)
    _, err := mc.Get("foo")

    if err == nil {
        t.Fatalf("Expecting error retrieving foo after expiration")
    }
}

func TestGetMulti(t *testing.T) {
    mc.Set(&memcache.Item{Key: "foo", Value: []byte("bar")})
    mc.Set(&memcache.Item{Key: "hello", Value: []byte("world")})

    r, err := mc.GetMulti([]string{"foo", "hello", "other"})

    if err != nil {
        t.Fatalf("Get multi foo, hello, other returned error " + err.Error())
    }

    s := len(r)

    if s != 2 {
        t.Fatalf("Expecting get multi to return 2 items but " + strconv.Itoa(s) + " returned")
    }
}

func TestFlush(t *testing.T) {
    err := mc.FlushAll()

    if err != nil {
        t.Fatalf("FlushAll returned error " + err.Error())
    }

    v, err := mc.Get("foo")

    if err == nil {
        t.Fatalf("Expecting error retrieving foo after flush_all but got " + string(v.Value))
    }
}