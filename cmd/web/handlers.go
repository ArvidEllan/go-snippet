
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"errors"

	"snippetbox.alexedwards.net/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost{
		w.Header().Set("Allow",http.MethodPost)
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly,slowly!\n\n- Kobayashi issa"
	expires := 7

	id,err := app.snippets.Insert(title,content,expires)
	if err != nil{
		app.serverError(w,err)
		return
	}

	http.Redirect(w,r,fmt.Sprintf("/snippet/view?id=%d",id),http.StatusSeeOther)

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _ , snippet := range snippets{
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a new snippet"))
}
 func (app *application) snippetCreate(w http.ResponseWriter,r *http.Request){

	if r.Method != http.MethodPost {
		w.Header().Set("Allow",http.MethodPost)
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut Slowly,slowly!\n\n- KObayashi Issa"
	expires := 7

	id,err := app.snippets.Insert(title,content,expires)
	if err!=nil{
		app.serverError(w,err)
	}
	http.Redirect(w,r,fmt.Sprintf("/snippet/view?id=%d",id),http.StatusSeeOther)
 }

 func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}
	snippet,err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err,models.ErrNoRecord){
			app.NotFound(w)
		} else {
			app.serverError(w,err)
		}
		return
	}
	fmt.Fprintf(w, "%+V",snippet)
	
 }