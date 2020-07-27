package lucenequery

import "strings"

type And struct {
	Clauses []Clause
}

func (a And) String() string {
	var parts []string
	for _, i := range a.Clauses {
		parts = append(parts, i.String())
	}

	return `(` + strings.Join(parts, " AND ") + `)`
}

type Or struct {
	Clauses []Clause
}

func (a Or) String() string {
	var parts []string
	for _, i := range a.Clauses {
		parts = append(parts, i.String())
	}

	return `(` + strings.Join(parts, " OR ") + `)`
}

type NOOP struct {
	Clauses []Clause
}

func (n NOOP) String() string {
	var parts []string
	for _, i := range n.Clauses {
		parts = append(parts, i.String())
	}

	return `(` + strings.Join(parts, " ") + `)`
}
