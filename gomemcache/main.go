package main

func main() {
	server := Server{"127.0.0.1", "1234"}
	server.Run(4 * 1024 * 1024)
}
