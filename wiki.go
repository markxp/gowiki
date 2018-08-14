package main

import (
	"net/http"
)

func main() {
	// register to http default serve mux
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	
	// listen to :6060, use default serve mux
	http.ListenAndServe(":6060", nil)
}
