package models

// Schema for a UFC event
// Ex: http://www.ufcstats.com/event-details/f3743d8ef5dde970, UFC 303: Pereira vs. Prochazka 2, 29-Jun-24, Las Vegas, Nevada, USA

type UFCEvent struct {
	EventURL      string
	EventTitle    string
	EventDate     string
	EventLocation string
}
