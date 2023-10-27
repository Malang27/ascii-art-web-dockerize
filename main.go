package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Créez un gestionnaire de fichiers statiques pour le répertoire "static".
	staticHandler := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)

	// Démarrer le serveur HTTP sur le port 8080
	fmt.Println("Serveur en cours d'exécution sur le port 8080...")
	fmt.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
