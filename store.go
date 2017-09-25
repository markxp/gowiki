package main

import (
	"io/ioutil"
	"path/filepath"
)

var defaultPageStore = &pageFileStore{
	Base: "storage/",
	Suf:  ".txt",
}

func (store pageFileStore) Save(p *Page) error {
	filename := p.Title + store.Suf
	return ioutil.WriteFile(filepath.Join(store.Base, filename), p.Body, 0600)
}

func (store pageFileStore) Load(title string) (*Page, error) {
	filename := title + store.Suf
	bs, err := ioutil.ReadFile(filepath.Join(store.Base, filename))
	if err != nil {
		return nil, err
	}
	return &Page{
		Title: title,
		Body:  bs,
	}, nil
}