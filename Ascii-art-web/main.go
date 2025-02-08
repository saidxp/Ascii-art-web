package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	// "web/assci"
)

type data struct {
	Message     string
	Title       string
	Hope        string
	ASCIIOutput [][]string
}

func handleroot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/HELLO.html"))
	d := data{
		Message: "My page web !!",
		Title:   "Page web ",
		Hope:    "[Choice and type !!]",
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, d)
}

func handlesub(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/submit" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	var Detecttype string
	name := r.FormValue("name")
	///detect := checker(name)

	// name = strings.ReplaceAll(name, "\n", "\\n")
	Detecttype = r.FormValue("asciid")
	if name == "" {
		http.Error(w, "Missing 'name' field in form data", http.StatusBadRequest)
		return
	}
	// handleroot(w, r)
	slice := ascci.Makeandreturn(name, Detecttype)

	tmpl := template.Must(template.ParseFiles("template/HELLO.html"))
	tmpl.Execute(w, data{
		Message:     "My page web !!",
		Title:       "Page web ",
		Hope:        ">>> THE IS THE RESULT OF WANT U CHOICE -- BE HAPPY 😉 <<< ",
		ASCIIOutput: slice,
	})
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, slice)
}

func customHandler(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/filecss/") {
			mux.ServeHTTP(w, r) // Allow static files
			return
		}
		switch {
		case r.URL.Path == "/" || r.URL.Path == "/submit":
			mux.ServeHTTP(w, r)
		case strings.HasPrefix(r.URL.Path, "/submit/") ||
			strings.HasPrefix(r.URL.Path, "/") && r.URL.Path != "/":
			http.Error(w, "404 page not found", http.StatusNotFound)
		default:
			http.NotFound(w, r)
		}
	})
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("filecss"))
	mux.Handle("/filecss/", http.StripPrefix("/filecss/", fs))

	mux.HandleFunc("/", handleroot)
	mux.HandleFunc("/submit", handlesub)

	server := &http.Server{
		Addr:    ":7090",
		Handler: customHandler(mux),
	}
	log.Println("Server starting on: http://localhost:7090")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
