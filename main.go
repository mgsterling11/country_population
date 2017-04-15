package main

import "net/http"

func serverConnect(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server is connected to port 8080"))
}



func main() {
	// tests local server is running on port 8080
	http.HandleFunc("/test-server", serverConnect)
	http.ListenAndServe(":8080", nil)
}

