package main

func main() {
    server := NewServer()

    server.commands["echo"] = Echo{}
    server.commands["quit"] = Quit{}

    server.Start()
}