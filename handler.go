package main

import (
	"ascii-art-web/asciifs"
	"html/template"
	"net/http"
)

// Data is a struct that will be sent as a respond
type Data struct {
	Output    string
	ErrorNum  int
	ErrorText string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	d := Data{}

	if r.URL.Path != "/" {
		d.ErrorNum = 404
		d.ErrorText = "Page Not Found"
		w.WriteHeader(404)
		errorHandler(w, r, d)
		return
	}

	// 	// Données pour le modèle
	data := map[string]interface{}{
		"Result": "",
	}
	AffichePage(w, r, "index.html", data)
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer les données du texte et de la bannière à partir du formulaire POST

	d := Data{}
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		if text == "" {
			AffichePage(w, r, "index.html", "template obligatoire")
			return
		}

		// Utiliser la fonction Read pour lire le fichier de la bannière correspondante
		dict, err := asciifs.Read(banner + ".txt")
		if err != nil {
			// Gérer l'erreur ici (par exemple, renvoyer un message d'erreur à l'utilisateur)
			d.ErrorNum = 500
			d.ErrorText = "internal server error"
			w.WriteHeader(500)
			errorHandler(w, r, d)
			//http.Error(w, "Erreur lors de la lecture de la bannière", http.StatusInternalServerError)
			return
		}

		// Utiliser la fonction PrintArt pour générer l'ASCII art en utilisant le dictionnaire (dict)
		var result string
		asciifs.PrintArt(text, dict, &result)

		// Afficher le résultat sur la page d'accueil en utilisant le modèle HTML
		data := map[string]interface{}{
			"Result": result,
		}

		if result == "" {
			d.ErrorNum = 400
			d.ErrorText = "bad request"
			w.WriteHeader(400)
			errorHandler(w, r, d)
			return
		}
		AffichePage(w, r, "index.html", data)
	} else {

		d.ErrorNum = 405
		d.ErrorText = "Method not allowed"
		w.WriteHeader(405)
		errorHandler(w, r, d)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, d Data) {
	// Compiler le modèle HTML
	t, err := template.ParseFiles("./htmltemplate/error.html")
	if err != nil {
		http.Error(w, "Erreur lors de la génération de", http.StatusInternalServerError)
		return
	}

	// Exécuter le modèle HTML en utilisant les données et écrire la réponse HTTP
	if err := t.Execute(w, d); err != nil {
		http.Error(w, "Erreur lors de l'écriture de la réponse", http.StatusInternalServerError)
		return
	}
}

func AffichePage(w http.ResponseWriter, r *http.Request, page string, d interface{}) {
	path := "./htmltemplate/" + page
	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Erreur lors de la génération de", http.StatusInternalServerError)
		return
	}

	// Exécuter le modèle HTML en utilisant les données et écrire la réponse HTTP
	if err := t.Execute(w, d); err != nil {
		http.Error(w, "--------Erreur champs de texte obligatoire-------", http.StatusInternalServerError)
		return
	}

}
