package main

func main() {
	server := NewServer(":8087")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
