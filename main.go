package main

import (
	"log"
	"net/http"
	"fmt"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello from snippetbox"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request){

	id,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		http.NotFound(w , r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d ")
}
func SnippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific snippet"))
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) { 
	if r.Method != "POST"{
		w.Header().Set("Allow","POST")
		http.Error(w,"Method Not Allowed",405)
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", SnippetView)
	mux.HandleFunc("/snippet/create", SnippetCreate)
	mux.HandleFunc("/snippet/show", ShowSnippet)

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	log.Print("starting server on :4000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
