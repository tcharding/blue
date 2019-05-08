// Package adt provides abstract data types for the service.
package adt

type Article struct {
	ID    int
	Title string
	Date  string
	Body  string
	Tags  []string
}
