package api

import (
	"fmt"
	"net/http"
	"strings"
)

// BooksContent contains the raw JSON of /books.jsonld
const BooksContent = `{
	"@id": "/books.jsonld",
	"hydra:member": [
		"/books/1.jsonld",
		"/books/2.jsonld"
	],
	"foo": [
		{"bar": [{"a": "b"}, {"c": "d"}], "car": "caz"},
		{"bar": [{"a": "d"}, {"c": "e"}], "car": "caz2"}
	]
	}`

// Fixtures provides a dummy API
func Fixtures(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/ld+json")
	rw.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
	rw.Header().Add("Access-Control-Allow-Credentials", "true")

	var myCookieValue string
	if cookie, err := req.Cookie("myCookie"); err == nil {
		myCookieValue = cookie.Value
		rw.Header().Add("Passed-Cookie", myCookieValue)
	}

	if strings.HasPrefix(req.RequestURI, "/books.jsonld") {
		if myCookieValue == "" {
			http.SetCookie(rw, &http.Cookie{Name: "myCookie", Value: "foo"})
		}

		fmt.Fprint(rw, BooksContent)

		return
	}

	fmt.Fprint(rw, `{
"@id": "/books/1.jsonld",
"title": "Book 1",
"description": "A good book",
"author": "/authors/1.jsonld"
}`)
}
