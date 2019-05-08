Blue Technical Test
===================

Recruitment test for Fairfax media.

Author: Tobin C. Harding <me@tobin.cc>

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
		{"Tag":"health","Count":2,"Articles":[0,1],"RelatedTags":["fitness","science"]}

