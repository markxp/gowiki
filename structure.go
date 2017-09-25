package main

type Page struct {
	Title string
	Body  []byte
}

type PageInterface interface {
	Save(p *Page) error
	Load(title string) (*Page, error)
}

type pageFileStore struct {
	Base string
	Suf  string
}


