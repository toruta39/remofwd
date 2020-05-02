package remofwd

import "time"

type Event struct {
	RawName      string
	RawValue     float64 `json:"val"`
	RawTimestamp string  `json:"created_at"`
}

func (e *Event) Name() string {
	return e.RawName
}

func (e *Event) Value() float64 {
	return e.RawValue
}

func (e *Event) Timestamp() time.Time {
	ts, err := time.Parse(time.RFC3339, e.RawTimestamp)
	if err != nil {
		return time.Now()
	}
	return ts
}

type NewestEvents struct {
	Humidity     Event `json:"hu"`
	Illumination Event `json:"il"`
	Motion       Event `json:"mo"`
	Temperature  Event `json:"te"`
}

type Device struct {
	Name         string `json:"name"`
	NewestEvents `json:"newest_events"`
}

func (d *Device) Humidity() *Event {
	e := &d.NewestEvents.Humidity

	if e.RawName == "" {
		e.RawName = "Humidity"
	}

	return e
}

func (d *Device) Illumination() *Event {
	e := &d.NewestEvents.Illumination

	if e.RawName == "" {
		e.RawName = "Illumination"
	}

	return e
}

func (d *Device) Motion() *Event {
	e := &d.NewestEvents.Motion

	if e.RawName == "" {
		e.RawName = "Motion"
	}

	return e
}

func (d *Device) Temperature() *Event {
	e := &d.NewestEvents.Temperature

	if e.RawName == "" {
		e.RawName = "Temperature"
	}

	return e
}
