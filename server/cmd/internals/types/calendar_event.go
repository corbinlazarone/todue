package types

type EventTime struct {
	Date     string `json:"date"`
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type EventOverride struct {
	Method  string `json:"method"` // set default to email
	Minutes int    `json:"minutes"`
}

type EventReminder struct {
	UseDefault bool            `json:"useDefault"` // set default to false
	Overrides  []EventOverride `json:"overrides"`
}

type CalendarEvent struct {
	Summary     string        `json:"summary"`
	Description string        `json:"description"`
	Start       EventTime     `json:"start"`
	End         EventTime     `json:"end"`
	ColorId     string        `json:"colorId"`
	Reminders   EventReminder `json:"reminders"`
}

type TimeConversionError struct {
	AssignmentID   string   `json:"assignment_id"`
	AssignmentName string   `json:"assignment_name"`
	Error          []string `json:"errors"`
}

func NewCalendarEvent() *CalendarEvent {
	return &CalendarEvent{
		Reminders: EventReminder{
			Overrides: []EventOverride{
				{
					Method: "email",
				},
			},
		},
	}
}
