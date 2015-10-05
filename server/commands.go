package server

type Command func(server Server, conn Conn, args []string) bool

func Add(server Server, conn Conn, args []string) bool {
	return false
}

func Count(server Server, conn Conn, args []string) bool {
	return false
}

func Get(server Server, conn Conn, args []string) bool {
	return false
}

func Delete(server Server, conn Conn, args []string) bool {
	return false
}

func DeIndex(server Server, conn Conn, args []string) bool {
	return false
}

func Index(server Server, conn Conn, args []string) bool {
	return false
}

func Quit(server Server, conn Conn, args []string) bool {
	return false
}

func Set(server Server, conn Conn, args []string) bool {
	return false
}
