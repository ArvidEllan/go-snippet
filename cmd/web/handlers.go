package main

import(
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r);
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal Server Error",500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil{
		log.Println(err.Error())
		http.Error(w,"Internal server error ",500 )
	}
	w.Write([]byte("hello from snippetbox"))
}

func showSnippet(w http.ResponseWriter , r *http.Request){
	id,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	fmt.Printf("%s",w, "Display a specific snippet with ID %d" , id)
}

func createSnippet(w http.ResponseWriter , r *http.Request){
	if r.Method != "POST"{
		w.Header().Set("Allow", "POST")
		http.Error(w, "method Not Allowed",405)
		return
	}
	w.Write([]byte("create a new snipper"))
}