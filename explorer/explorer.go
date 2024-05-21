package explorer

import (
	"coin/blockchain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

type homepage struct {
	PageTitle string
	Block     []*blockchain.Block
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	// telp := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homepage{"Home", blockchain.GetBlockchain().AllBlock()}
	// telp.Execute(w, data)
	templates.ExecuteTemplate(w, "home", data)

}
func handleAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

}
func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleFunc)
	http.HandleFunc("/add", handleAdd)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
