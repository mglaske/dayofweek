package dayofweek

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const (
	Monday    Dow = 1 << iota // 1
	Tuesday                   // 2
	Wednesday                 // 4
	Thursday                  // 8
	Friday                    // 16
	Saturday                  // 32
	Sunday                    // 64
)

type Dow int

// Return a new Dow with the given days set.
func New(mon, tue, wed, thu, fri, sat, sun bool) *Dow {
	var d Dow
	d.Set(mon, tue, wed, thu, fri, sat, sun)
	return &d
}

// Set Dow to the given days, mon, tue, wed, thu, fri, sat, sun.
func (me *Dow) Set(mon, tue, wed, thu, fri, sat, sun bool) {
	if mon {
		*me |= Monday
	}
	if tue {
		*me |= Tuesday
	}
	if wed {
		*me |= Wednesday
	}
	if thu {
		*me |= Thursday
	}
	if fri {
		*me |= Friday
	}
	if sat {
		*me |= Saturday
	}
	if sun {
		*me |= Sunday
	}
}

// Is a given day set (Eg. x.IsSet(dayofweek.Wednesday) )
func (me Dow) IsSet(p Dow) bool {
	return me&p == p
}

// Is Monday Set
func (me Dow) IsMonday() bool {
	return me.IsSet(Monday)
}

// Is Tuesday Set
func (me Dow) IsTuesday() bool {
	return me.IsSet(Tuesday)
}

// Is Wednesday Set
func (me Dow) IsWednesday() bool {
	return me.IsSet(Wednesday)
}

// Is Thursday Set
func (me Dow) IsThursday() bool {
	return me.IsSet(Thursday)
}

// Is Friday Set
func (me Dow) IsFriday() bool {
	return me.IsSet(Friday)
}

// Is Saturday Set
func (me Dow) IsSaturday() bool {
	return me.IsSet(Saturday)
}

// Is Sunday Set
func (me Dow) IsSunday() bool {
	return me.IsSet(Sunday)
}

// Convert time.Weekday to dayofweek.Dow
func (me Dow) WeekdayToDow(wd time.Weekday) Dow {
	// Because time.Weekday starts Sunday=0 through Saturday=6 (no 7)
	// so, weekday Monday=1, is 1<<1-1 = 1
	// Tuesday=2 (1 << 2-1) = 2 ; Saturday=6 (1 << 6-1) = 32
	if wd == 0 {
		return Sunday // 64
	}
	return Dow(1 << (wd - 1))
}

// See if time.Weekday is set
func (me Dow) IsWeekday(wd time.Weekday) bool {
	return me.IsSet(me.WeekdayToDow(wd))
}

// Check if time.Time's day is set
func (me Dow) OnDate(t time.Time) bool {
	return me.IsSet(me.WeekdayToDow(t.Weekday()))
}

// Check if today is set.
func (me Dow) Today() bool {
	return me.OnDate(time.Now())
}

// Add the given dow day (eg. x.AddDay(dayofweek.Monday) )
func (me *Dow) AddDay(p Dow) {
	if !me.IsSet(p) {
		*me |= p
	}
}

// Add Monday
func (me Dow) AddMonday() {
	me.AddDay(Monday)
}

// Add Tuesday
func (me Dow) AddTuesday() {
	me.AddDay(Tuesday)
}

// Add Wednesday
func (me Dow) AddWednesday() {
	me.AddDay(Wednesday)
}

// Add Thursday
func (me Dow) AddThursday() {
	me.AddDay(Thursday)
}

// Add Friday
func (me Dow) AddFriday() {
	me.AddDay(Friday)
}

// Add Saturday
func (me Dow) AddSaturday() {
	me.AddDay(Saturday)
}

// Add Sunday
func (me Dow) AddSunday() {
	me.AddDay(Sunday)
}

// Remove a day from us (eg. dayofweek.Wednesday
func (me *Dow) RemoveDay(p Dow) {
	if me.IsSet(p) {
		*me ^= p
	}
}

// Remove Monday
func (me Dow) RemoveMonday() {
	me.RemoveDay(Monday)
}

// Remove Tuesday
func (me Dow) RemoveTuesday() {
	me.RemoveDay(Tuesday)
}

// Remove Wednesday
func (me Dow) RemoveWednesday() {
	me.RemoveDay(Wednesday)
}

// Remove Thursday
func (me Dow) RemoveThursday() {
	me.RemoveDay(Thursday)
}

// Remove Friday
func (me Dow) RemoveFriday() {
	me.RemoveDay(Friday)
}

// Remove Saturday
func (me Dow) RemoveSaturday() {
	me.RemoveDay(Saturday)
}

// Remove Sunday
func (me Dow) RemoveSunday() {
	me.RemoveDay(Sunday)
}

// Clear out all days (unset)
func (me *Dow) Clear() {
	*me = 0
}

// Parse updates your Dow object from a given string such as (Mon, Tuesday, wed)
// Must be csv.  Returns error if unable to parse a value.
func (me *Dow) Parse(s string) error {
	s = strings.ToLower(s)
	if strings.Contains(s, "mon") {
		*me |= Monday
	}
	if strings.Contains(s, "tue") {
		*me |= Tuesday
	}
	if strings.Contains(s, "wed") {
		*me |= Wednesday
	}
	if strings.Contains(s, "thu") {
		*me |= Thursday
	}
	if strings.Contains(s, "fri") {
		*me |= Friday
	}
	if strings.Contains(s, "sat") {
		*me |= Saturday
	}
	if strings.Contains(s, "sun") {
		*me |= Sunday
	}
	if len(s) > 0 && *me == 0 {
		errors.New("Parse seems to have failed")
	}
	return nil
}

// Parse a string of dow, and remove those days.
func (me *Dow) ParseRemove(s string) error {
	var other Dow
	err := other.Parse(s)
	if err != nil {
		return err
	}
	*me = *me - other
	return nil
}

// Is the supplied Dow equal to us.
func (me Dow) Equal(o Dow) bool {
	return me == o
}

// Pretty Print the day of week (eg. "Monday, Wednesday, Friday")
func (me Dow) String() string {
	var days []string
	if me.IsMonday() {
		days = append(days, "Monday")
	}
	if me.IsTuesday() {
		days = append(days, "Tuesday")
	}
	if me.IsWednesday() {
		days = append(days, "Wednesday")
	}
	if me.IsThursday() {
		days = append(days, "Thursday")
	}
	if me.IsFriday() {
		days = append(days, "Friday")
	}
	if me.IsSaturday() {
		days = append(days, "Saturday")
	}
	if me.IsSunday() {
		days = append(days, "Sunday")
	}
	return strings.Join(days, ", ")
}

// Marshal function to dump output as a string eg {"Monday, Tuesday"}
func (me Dow) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.String())
}

// Unmarshal your JSON input into a Dow object.
// eg: {"run_on_days": "Mon, Tue, Thursday"}
func (me *Dow) UnmarshalJSON(d []byte) error {
	var s string
	if err := json.Unmarshal(d, &s); err != nil {
		return err
	}
	return me.Parse(s)
}
