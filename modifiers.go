package lucenequery

import "fmt"

func Not(c Clause) NotModifier {
	return NotModifier{c}
}

type NotModifier struct {
	Clause Clause
}

func (n NotModifier) String() string {
	return fmt.Sprintf(`-%s`, n.Clause)
}

func Boost(c Clause, m float64) BoostModifier {
	return BoostModifier{c, m}
}

type BoostModifier struct {
	Clause   Clause
	Modifier float64
}

func (b BoostModifier) String() string {
	return fmt.Sprintf(`%s^%G`, b.Clause, b.Modifier)
}

func Proximity(c Clause, m int) ProximityModifier {
	return ProximityModifier{c, m}
}

type ProximityModifier struct {
	Clause   Clause
	Modifier int
}

func (p ProximityModifier) String() string {
	return fmt.Sprintf(`%s~%d`, p.Clause, p.Modifier)
}
