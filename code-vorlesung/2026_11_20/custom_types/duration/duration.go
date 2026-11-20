package duration

// Dieser Code wurde nicht direkt in der Vorlesung
// gezeigt, sondern ist eine Übung gewesen, deren
// Lösung hier gezeigt wird.

type Duration int

func FromSeconds(s int) Duration {
	return Duration(s)
}

func FromMinutes(m int) Duration {
	return FromSeconds(m * 60)
}

func FromHours(h int) Duration {
	return FromMinutes(h * 60)
}

func (d Duration) Seconds() int {
	return int(d)
}

func (d Duration) Minutes() int {
	return d.Seconds() / 60
}

func (d Duration) Hours() int {
	return d.Minutes() / 60
}
