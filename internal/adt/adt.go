// Package adt provides abstract data types for the service.
package adt

import (
	"fmt"
)

// Article is used to store an article in the database.
type Article struct {
	ID    int
	Title string
	Date  string
	Body  string
	Tags  []string
}

// String satisfies the stringer interface.
func (a *Article) String() string {
	return fmt.Sprintf("Article (%d): %s", a.ID, a.Title)
}

// TagView holds data for the tags/ endpoint.
type TagView struct {
	Tag         string
	Count       int
	Articles    []int
	RelatedTags []string
	rtMap       map[string]bool
}

// NewTagView creates a new TagView.
func NewTagView(tag string) *TagView {
	tv := &TagView{
		Tag:   tag,
		rtMap: make(map[string]bool), // Map keyed by tag.
	}
	return tv
}

// AddRelatedTags adds tags to the TagView's related tags.
// All tags in RelatedTags are unique.
func (tv *TagView) AddRelatedTags(tags []string) {
	for _, tag := range tags {
		if tag == tv.Tag {
			continue
		}
		if _, ok := tv.rtMap[tag]; !ok {
			tv.rtMap[tag] = true
			tv.RelatedTags = append(tv.RelatedTags, tag)
		}
	}
}

// IsEmpty returns true if TagView references 0 articles.
func (tv *TagView) IsEmpty() bool {
	return len(tv.Articles) == 0
}
