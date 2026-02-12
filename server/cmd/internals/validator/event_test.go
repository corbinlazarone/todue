package validator

import (
	"testing"
)

type testEvent struct {
	name    string
	wantErr bool
	Event
}

func TestValidateDueDate(t *testing.T) {
	tests := []testEvent{
		{
			name:    "valid due date",
			wantErr: false,
			Event: Event{
				DueDate: "2025-12-11",
				AllDay:  true,
			},
		},
		{
			name:    "invalid due date",
			wantErr: true,
			Event: Event{
				DueDate: "2025-12-11T00:00:00Z",
				AllDay:  true,
			},
		},
		{
			name:    "invalid due date",
			wantErr: true,
			Event: Event{
				DueDate: "2025-12-11T00:00:00",
				AllDay:  true,
			},
		},
		{
			name:    "valid due date, all day marked false",
			wantErr: false,
			Event: Event{
				DueDate: "2025-12-10",
				AllDay:  false,
			},
		},
		{
			name:    "invalid due date, all day marked false",
			wantErr: true,
			Event: Event{
				DueDate: "",
				AllDay:  false,
			},
		},
		{
			name:    "invalid due date, all day marked true",
			wantErr: true,
			Event: Event{
				DueDate: "",
				AllDay:  true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.Event.validateDueDate()
			assertError(t, test.wantErr, err)
		})
	}
}

func TestValidateDateTimes(t *testing.T) {
	tests := []testEvent{
		{
			name:    "invalid start time format",
			wantErr: true,
			Event: Event{
				StartTime: "25:00",
				EndTime:   "01:00",
			},
		},
		{
			name:    "valid start time HH:MM",
			wantErr: false,
			Event: Event{
				StartTime: "00:00",
				EndTime:   "01:00",
			},
		},
		{
			name:    "valid start time HH:MM:SS",
			wantErr: false,
			Event: Event{
				StartTime: "00:00:00",
				EndTime:   "01:00:00",
			},
		},
		{
			name:    "invalid end time format",
			wantErr: true,
			Event: Event{
				StartTime: "00:00",
				EndTime:   "25:00",
			},
		},
		{
			name:    "valid end time HH:MM",
			wantErr: false,
			Event: Event{
				StartTime: "00:00",
				EndTime:   "00:00",
			},
		},
		{
			name:    "valid end time HH:MM:SS",
			wantErr: false,
			Event: Event{
				StartTime: "00:00:00",
				EndTime:   "00:00:00",
			},
		},
		{
			name:    "invalid blank end time",
			wantErr: true,
			Event: Event{
				StartTime: "00:00",
				EndTime:   "",
			},
		},
		{
			name:    "invalid blank start time",
			wantErr: true,
			Event: Event{
				StartTime: "",
				EndTime:   "00:00",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.Event.validateDateTimes()
			assertError(t, test.wantErr, err)
		})
	}
}

func TestValidateReminder(t *testing.T) {
	tests := []testEvent{
		{
			name:    "Invalid reminder",
			wantErr: true,
			Event: Event{
				Reminder: 1,
			},
		},
		{
			name:    "Valid reminder",
			wantErr: false,
			Event: Event{
				Reminder: 0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.Event.validateReminder()
			assertError(t, test.wantErr, err)
		})
	}
}

func TestValidateColor(t *testing.T) {
	tests := []testEvent{
		{
			name:    "Invalid color",
			wantErr: true,
			Event: Event{
				Color: "#invalidhex",
			},
		},
		{
			name:    "Valid color",
			wantErr: false,
			Event: Event{
				Color: "#7986cb",
			},
		},
		{
			name:    "Empty color",
			wantErr: true,
			Event: Event{
				Color: "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.Event.validateColor()
			assertError(t, test.wantErr, err)
		})
	}
}

func assertError(t *testing.T, wantErr bool, err error) {
	if wantErr {
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	} else {
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	}
}
