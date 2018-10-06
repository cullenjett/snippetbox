package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/justinas/nosurf"

	"snippetbox.org/pkg/models"
)

type HTMLData struct {
	CSRFToken string
	Form      interface{}
	Flash     string
	LoggedIn  bool
	Path      string
	Snippet   *models.Snippet
	Snippets  []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func (app *App) RenderHTML(w http.ResponseWriter, r *http.Request, page string, data *HTMLData) {
	if data == nil {
		data = &HTMLData{}
	}

	data.Path = r.URL.Path

	data.CSRFToken = nosurf.Token(r)

	var err error
	data.LoggedIn, err = app.LoggedIn(r)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	files := []string{
		filepath.Join(app.HTMLDir, "base.html"),
		filepath.Join(app.HTMLDir, page),
	}

	funcMap := template.FuncMap{
		"humanDate": humanDate,
	}

	ts, err := template.New("").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	buf.WriteTo(w)
}
