package lucenequery

import "fmt"

type Not struct {
	Clause Clause
}

func (n Not) String() string {
	return fmt.Sprintf(`-%s`, n.Clause)
}

type Boost struct {
	Clause   Clause
	Modifier float64
}

func (b Boost) String() string {
	return fmt.Sprintf(`%s^%G`, b.Clause, b.Modifier)
}

type Proximity struct {
	Clause   Clause
	Modifier int
}

func (p Proximity) String() string {
	return fmt.Sprintf(`%s~%d`, p.Clause, p.Modifier)
}
