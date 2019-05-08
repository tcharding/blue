// Package main provides the service endpoint.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tcharding/blue/internal/adt"
)

func main() {
	db := database{}
	http.HandleFunc("/articles", db.articles)
	http.HandleFunc("/tags", db.tags)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

type database struct {
	Articles map[int]adt.Article
	Tags     map[string]adt.Tag
}

func (db database) articles(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "articles endpoint in working\n")
}

func (db database) tags(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "tags endpoint in working\n")
}
