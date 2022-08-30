package page

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	// Création d'une page
	var s interface{}
	fmt.Printf("s: %v\n", s)
	// Création d'une nouvelle instance de template
	t := template.New("Page")

	// Déclaration des fichiers à parser
	t = template.Must(t.ParseFiles("template/index.html"))

	// Exécution de la fusion et injection dans le flux de sortie
	// La variable p sera réprésentée par le "." dans le layout
	// Exemple {{.}} == p
	err := t.ExecuteTemplate(w, "index", s)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}
