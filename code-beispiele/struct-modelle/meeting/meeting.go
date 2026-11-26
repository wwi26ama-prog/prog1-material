package meeting

type Meeting struct {
	Topic        string
	Date         Date
	Duration     Duration
	Participants []Person
}
