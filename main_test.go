package lucenequery_test

import (
	lq "github.com/nodefortytwo/lucenequery"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLuceneQuery(t *testing.T) {

	q := lq.New(lq.Term("title", "foo")).String()
	assert.Equal(t, `title:foo`, q)

	q = lq.New(lq.Term("title", "foo bar")).String()
	assert.Equal(t, `title:"foo bar"`, q)

	q = lq.New(lq.And(
		lq.Term("title", "foo bar"),
		lq.Term("body", "quick fox"),
	)).String()
	assert.Equal(t, `(title:"foo bar" AND body:"quick fox")`, q)

	q = lq.New(lq.Or(
		lq.And(
			lq.Term("title", "foo bar"),
			lq.Term("body", "quick fox"),
		),
		lq.Term("title", "fox"),
	)).String()
	assert.Equal(t, `((title:"foo bar" AND body:"quick fox") OR title:fox)`, q)

	q = lq.New(lq.NOOP(
		lq.Term("title", "foo"),
		lq.Not(lq.Term("title", "bar")),
	)).String()
	assert.Equal(t, `(title:foo -title:bar)`, q)

	q = lq.New(lq.Proximity(lq.Term("", "foo bar"), 4)).String()
	assert.Equal(t, `"foo bar"~4`, q)

	q = lq.New(lq.Range("mod_date", "20020101", "20030101")).String()
	assert.Equal(t, `mod_date:[20020101 TO 20030101]`, q)

	q = lq.New(lq.NOOP(
		lq.Boost(lq.Or(lq.Term("title", "foo"), lq.Term("title", "bar")), 1.5),
		lq.Or(lq.Term("body", "foo"), lq.Term("body", "bar")),
	)).String()
	assert.Equal(t, `((title:foo OR title:bar)^1.5 (body:foo OR body:bar))`, q)

}
