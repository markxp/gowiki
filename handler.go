package main

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
)

var (
	pageStore PageInterface
	templates *template.Template
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)

func init() {
	pageStore = defaultPageStore
}

func getTitle(urlPath string) (string, error) {
	m := validPath.FindStringSubmatch(urlPath)
	if m == nil {
		return "", errors.New("invalid page title")
	}
	return m[2], nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := getTitle(r.URL.Path)
		if err != nil || title == "" {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pageStore.Load(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pageStore.Load(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{
		Title: title,
		Body:  []byte(body),
	}
	err := pageStore.Save(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	base := "template"
	suf := ".html"
	glob := base + "/*" + suf
	if templates == nil {
		templates = template.Must(template.ParseGlob(glob))
	}
	err := templates.ExecuteTemplate(w, tmpl+suf, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
