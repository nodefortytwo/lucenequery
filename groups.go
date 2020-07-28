package lucenequery

import "strings"

func And(c ...Clause) AndClause {
	return AndClause{Clauses: c}
}

type AndClause struct {
	Clauses []Clause
}

func (a *AndClause) Add(c Clause) {
	a.Clauses = append(a.Clauses, c)
}

func (a AndClause) String() string {
	var parts []string
	for _, i := range a.Clauses {
		parts = append(parts, i.String())
	}

	return `(` + strings.Join(parts, " AND ") + `)`
}

func Or(c ...Clause) OrClause {
	return OrClause{Clauses: c}
}

type OrClause struct {
	Clauses []Clause
}

func (a *OrClause) Add(c Clause) {
	a.Clauses = append(a.Clauses, c)
}

func (a OrClause) String() string {
	var parts []string
	for _, i := range a.Clauses {
		parts = append(parts, i.String())
	}

	return `(` + strings.Join(parts, " OR ") + `)`
}

func NOOP(c ...Clause) NOOPClause {
	return NOOPClause{Clauses: c}
}

type NOOPClause struct {
	Clauses []Clause
}

func (n *NOOPClause) Add(c Clause) {
	n.Clauses = append(n.Clauses, c)
}

func (n NOOPClause) String() string {
	var parts []string
	for _, i := range n.Clauses {
		parts = append(parts, i.String())
	}

	return `(` + strings.Join(parts, " ") + `)`
}
