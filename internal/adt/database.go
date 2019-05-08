package adt

// Following the oft quoted Knuth 'optimisation is the root of all evil'
// we chose to use a simple in memory database implementation.  In order
// to solve the problem to specification this is sufficient.  Clearly this
// would not scale so a relational database would typically be used.

// Database holds the data for our service.
type Database struct {
	// articlesByID is keyed by ID.
	articlesByID map[int]*Article
	// articlesByDate is keyed by Date.
	articlesByDate map[string][]*Article
}

// NewDatabase creates a new database instance.
func NewDatabase() Database {
	return Database{
		articlesByID:   make(map[int]*Article),
		articlesByDate: make(map[string][]*Article),
	}
}

// AddArticle adds an article to the database.
func (db Database) AddArticle(a *Article) {
	a.ID = nextID()
	db.articlesByID[a.ID] = a
	if _, ok := db.articlesByDate[a.Date]; !ok {
		db.articlesByDate[a.Date] = []*Article{a}
	} else {
		db.articlesByDate[a.Date] = append(db.articlesByDate[a.Date], a)
	}
}

// idAccumulator stores the last used Article ID number.
var idAccumulator int

// nextID provides the next unique ID number, this is a rudimentary
// method of keeping ID numbers unique.  Does not scale because it does
// not take into consideration integer overflow.
func nextID() int {
	id := idAccumulator
	idAccumulator++
	return id
}

// ArticleByID gets the article with ID number id.
func (db Database) ArticleByID(id int) (*Article, bool) {
	a, ok := db.articlesByID[id]
	if !ok {
		return a, false
	}
	return a, true
}

// TagViewForDate constructs a TagView structure for the given date.
func (db Database) TagViewForDate(tag, date string) *TagView {
	tv := NewTagView(tag)
	articles, ok := db.articlesByDate[date]
	if ok {
		for _, a := range articles {
			if contains(a.Tags, tag) {
				tv.Articles = append(tv.Articles, a.ID)
				tv.AddRelatedTags(a.Tags)
				tv.Count += db.Count(a, date)
			}
		}
	}
	return tv
}

// contains returns true if tags contains tag.
// This is a naive implementation, Big-O N.
func contains(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

// Count returns the count for an article.
//
// It is not clear from the specification what this should be.
// Here we define count as the total of all articles in the
// database which have any one of the articles tags.
//
// WARNING: This is Big-O N squared.
func (db Database) Count(a *Article, date string) int {
	count := 0
	tags := a.Tags

	articles, ok := db.articlesByDate[date]
	if ok {
		for _, a := range articles {
			for _, tag := range tags {
				if contains(a.Tags, tag) {
					count++
				}
			}
		}
	}
	return count
}
