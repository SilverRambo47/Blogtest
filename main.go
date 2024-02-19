package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("index.html", http.FileServer(http.Dir(".")))

	fmt.Println("Serveur en cours d'écoute sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}
