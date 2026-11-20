package length

type Length int

func FromMillimeters(m int) Length {
	return Length(m)
}

func FromCentimeters(c int) Length {
	//return FromMillimeters(c * 10)
	return Length(c * 10)
}

func FromKilometers(k int) Length {
	return FromMillimeters(k * 1000000)
}

func (l Length) Millimeters() int {
	return int(l)
}

func (l Length) Centimeters() int {
	return l.Millimeters() / 10
}

func (l Length) Meters() int {
	return l.Centimeters() / 100
}

func (l Length) Kilometers() int {
	return l.Meters() / 1000
}
