Blue Technical Test
===================

Recruitment test for Fairfax Media.

Author: Tobin C. Harding <me@tobin.cc>


Assumptions
-----------

The following assumptions made while building the service.

- Service runs on localhost:8080

- The endpoint should return something useful for the index page.  Choose to
show available endpoints with example usage.

### Assumptions relating to tags

- Count is the total number of articles that are tagged with the tag or any of
its related tags.  See `internal/adt/database.go:Count()`.


Usage
-----

1. Clone this repo into `$GOPATH/src/github.com/tcharding/blue`

2. Build the project with `make`

3. Start the server with `bin/serve` <!-- listens on localhost:8080 -->

4. Test with web browser or curl, server starts up with a single article
   in the database (ID==0)

		## View article #0
		curl localhost:8080/articles/0
		{"ID":0,"Title":"latest science shows that potato ...

		## Add an article
		curl --header "Content-Type: application/json" \
		--request POST \
		--data '{"title":"test title","body":"body", "date":"2016-09-22","tags":["health"]}' \
		http://localhost:8080/articles

		## View newly added article #1
		curl localhost:8080/articles/1
		{"ID":1,"Title":"test title","Date":"2016-09-22","Body":"body","Tags":["health"]}

		## View tags for 2016-09-22
		curl http://localhost:8080/tags/health/2016-09-22
		{"Tag":"health","Count":6,"Articles":[0,1],"RelatedTags":["fitness","science"]}


Solution description
--------------------

Simple API with three endpoints.


### Choice of language

The Go programming language is used to implement the solution.  Go is a good
language to use for implementing microservices and also it is the advertised
language used by Fairfax media for the job role this project is targeting.

Some other things I like about Go; it is a statically typed language, has good
tooling support, and is a relatively simple language i.e. the standard library
is complete yet quite small.


### Project structure

The project is structured in a typical Golang manner

	$ tree
	.
	├── bin
	│   └── serve
	├── cmd
	│   └── serve
	│       └── main.go
	├── example.json
	├── internal
	│   └── adt
	│       ├── adt.go
	│       └── database.go
	├── LICENSE
	├── Makefile
	├── pkg
	└── README.md

- `cmd/serve/` - Source code for the executable.
- `bin/` - Project binary.
- `internal/` - Internal libraries i.e. not useful to other projects.
- `pkg/` - Libraries that may be useful to other projects (empty).


### Database

Implemented in `internal/adt/database.go`

A simple in memory database implementation is used.  This is trivial to code,
easy to read, and solves the problem.  It is not scalable, if this were a real
project a relational database running as a separate microservice would typically
be used.

The database includes two maps.  One map is keyed by ID, this ensures IDs are
unique and allows fast article look up for the `/articles/{id}` endpoint.  The
second map is keyed by date.  This map returns a slice of articles that all have
the same date.  This map is used by the `/tags/{tag}/{date}` endpoint.

In order to ensure unique IDs a glabol integer accumulator is used.  This value
is incremented for each article added to the database.  This is a very simple
solution to the problem of unique identifiers.  It does not take into
consideration integer overflow.  Also if removal of articles was to be
implemented then this method of identifier creation would need re-working.


### Abstract Data Types

The project uses two main abstract data types (ADT).  An 'Article' is the
structure used to represent articles and is what is stored in the database.  A
'TagView' is a dynamically created structure used to hold information returned
by the tags endpoint.  A TagView object is created dynamically because it
represents a snap shot of the database at the time of query.

Both structures use Golang's data member export feature, structure members are
capitialised so are accessible to code that imports the adt package.  We do not
need to implement setter and getter methods for this reason.  'Article'
implements a String method thereby satisfying the stringer interface.


### Libraries

This project is quite simple, the only libraries created were the database and
the ADTs.  Both of these are tightly coupled with the project and would not be
useful to another project.  These files are therefore located under `internal/`

If a more general library package was implemented it would of been placed under
`pkg/` so that other projects could import the library. The Golang compiler
enforces this behaviour disallowing import of internal packages from outside the
current projects path.

### Testing

Test Driven Development was not used while coding this project.  The service is
too simple to require unit tests.  The only testing done was using a web browser
and curl to check the endpoints were up and that the JSON was handled correctly.


Development time
----------------

This project was completed in 4 hours, a 3 hour session to do most of the coding
then an hour session the following day to lint, refactor, and document.

Additional features
-------------------

- A relational database as a separate microservice.
- Allow grepping articles.
- An endpoint to get all articles for a tag.
- An endpoint to get all articles for a date.


I enjoyed coding up this challenge, thanks.
