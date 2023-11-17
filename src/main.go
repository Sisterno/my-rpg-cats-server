package main

import (
	http_adapter "my-rpg-cats-server/src/adapters/http"
	"net/http"
)

func main() {
	http.HandleFunc("/", http_adapter.Handler)
	http.ListenAndServe(":8080", nil)
}
