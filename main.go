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

type Term struct {
	Field string
	Value string
}

func (t Term) String() string {
	pattern := `%s`
	if hasSpaces(t.Value) {
		pattern = `"%s"`
	}

	if t.Field == "" {
		return fmt.Sprintf(pattern, t.Value)
	}

	return fmt.Sprintf(`%s:`+pattern, t.Field, t.Value)
}

type Range struct {
	Field string
	From  string
	To    string
}

func (r Range) String() string {
	pattern := `%s:[%s TO %s]`

	return fmt.Sprintf(pattern, r.Field, r.From, r.To)
}

func hasSpaces(s string) bool {
	return strings.Contains(s, " ")
}
