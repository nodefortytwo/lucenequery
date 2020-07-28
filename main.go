package lucenequery

import (
	"fmt"
	"strings"
)

type Clause interface {
	String() string
}

func Clauses(c ...Clause) []Clause {
	return c
}

type Query struct {
	Clause Clause
}

func (q Query) String() string {
	return q.Clause.String()
}

func New(c Clause) Query {
	return Query{Clause: c}
}

func Term(field, value string) Clause {
	return TermSearch{Field: field, Value: value}
}

type TermSearch struct {
	Field string
	Value string
}

func (t TermSearch) String() string {
	pattern := `%s`
	if hasSpaces(t.Value) {
		pattern = `"%s"`
	}

	if t.Field == "" {
		return fmt.Sprintf(pattern, t.Value)
	}

	return fmt.Sprintf(`%s:`+pattern, t.Field, t.Value)
}

func Range(field, to, from string) Clause {
	return RangeSearch{Field: field, From: to, To: from}
}

type RangeSearch struct {
	Field string
	From  string
	To    string
}

func (r RangeSearch) String() string {
	pattern := `%s:[%s TO %s]`

	return fmt.Sprintf(pattern, r.Field, r.From, r.To)
}

func hasSpaces(s string) bool {
	return strings.Contains(s, " ")
}
