// Package main provides the service endpoint.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/tcharding/blue/internal/adt"
)

const (
	// Populate controls adding a single record to the database when server starts.
	Populate = true
)

// Service database.
var db adt.Database

// Index is the handler function for endpoint "/".
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Available endpoints:\n")
	fmt.Fprint(w, "POST:/articles\n")
	fmt.Fprint(w, "GET:/articles/{ID} e.g. /articles/1 \n")
	fmt.Fprint(w, "GET:/tags/{tag}/{date} e.g. /tags/health/20190931\n")
}

// Articles is the handler function for endpoint "/articles".
func Articles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var article adt.Article

	json.NewDecoder(r.Body).Decode(&article)
	db.AddArticle(&article)

	// Write content-type, statuscode, payload.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "Article successfully added to the database.\n\n")
}

// ArticleByID is the handler function for endpoint "/articles/{id}".
func ArticleByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	article, ok := db.ArticleByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	aj, err := json.Marshal(article)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Write content-type, statuscode, payload.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s\n", aj)
}

// Tags is the handler function for endpoint "/tags/{tag}/{date}".
func Tags(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")
	date := ps.ByName("date")

	tv := db.TagViewForDate(tag, date)
	if tv.IsEmpty() {
		http.NotFound(w, r)
		return
	}

	tj, err := json.Marshal(tv)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Write content-type, statuscode, payload.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s\n", tj)
}

func main() {
	db = adt.NewDatabase()

	if Populate {
		debugInitDB()
	}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/articles/:id", ArticleByID)
	router.POST("/articles", Articles)
	router.GET("/tags/:tag/:date", Tags)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// debugInitDB populates the database with some test data for debugging.
func debugInitDB() {
	a := &adt.Article{
		Title: "latest science shows that potato chips are better for you than sugar",
		Date:  "2016-09-22",
		Body:  "some text, potentially containing simple markup about how potato chips are great",
		Tags:  []string{"health", "fitness", "science"},
	}
	db.AddArticle(a)
}
