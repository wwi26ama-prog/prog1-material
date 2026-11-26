package moretypes

import "fmt"

type Time struct {
	Hours   int
	Minutes int
}

// String wird zur Ausgabe genutzt.
// Sie liefert einen String als Repräsentation
// von Time und wird z.B. von `fmt.Println()` benutzt.
func (t Time) String() string {
	return fmt.Sprintf("%v:%v", t.Hours, t.Minutes)

	//result := ""
	//
	//result += fmt.Sprint(t.Hours)
	//result += ":"
	//result += fmt.Sprint(t.Minutes)
	//
	//return result
}

// Normalize rechnet die Minuten und Stunden
// so um, dass die Stunden <= 23 und die Minuten <= 59 sind.
func (t *Time) Normalize() {
	t.Hours = (t.Hours + t.Minutes/60) % 24
	t.Minutes = t.Minutes % 60
}
