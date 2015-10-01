package main

func main() {
    server := NewServer()

    server.commands["echo"] = Echo{}
    server.commands["set"]  = Set{}
    server.commands["get"]  = Get{}
    server.commands["quit"] = Quit{}

    server.Start()
}