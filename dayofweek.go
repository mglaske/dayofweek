package dayofweek

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const (
	Monday Dow = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Dow int

func New(mon, tue, wed, thu, fri, sat, sun bool) *Dow {
	var d Dow
	d.Set(mon, tue, wed, thu, fri, sat, sun)
	return &d
}

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

func (me Dow) IsSet(p Dow) bool {
	return me&p == p
}

func (me Dow) IsMonday() bool {
	return me.IsSet(Monday)
}
func (me Dow) IsTuesday() bool {
	return me.IsSet(Tuesday)
}
func (me Dow) IsWednesday() bool {
	return me.IsSet(Wednesday)
}
func (me Dow) IsThursday() bool {
	return me.IsSet(Thursday)
}
func (me Dow) IsFriday() bool {
	return me.IsSet(Friday)
}
func (me Dow) IsSaturday() bool {
	return me.IsSet(Saturday)
}
func (me Dow) IsSunday() bool {
	return me.IsSet(Sunday)
}
func (me Dow) OnDate(t time.Time) bool {
	return me.IsSet(Dow(t.Weekday()))
}
func (me Dow) Today() bool {
	return me.OnDate(time.Now())
}

func (me *Dow) AddDay(p Dow) {
	if !me.IsSet(p) {
		*me |= p
	}
}
func (me Dow) AddMonday() {
	me.AddDay(Monday)
}
func (me Dow) AddTuesday() {
	me.AddDay(Tuesday)
}
func (me Dow) AddWednesday() {
	me.AddDay(Wednesday)
}
func (me Dow) AddThursday() {
	me.AddDay(Thursday)
}
func (me Dow) AddFriday() {
	me.AddDay(Friday)
}
func (me Dow) AddSaturday() {
	me.AddDay(Saturday)
}
func (me Dow) AddSunday() {
	me.AddDay(Sunday)
}

func (me *Dow) RemoveDay(p Dow) {
	if me.IsSet(p) {
		*me ^= p
	}
}
func (me Dow) RemoveMonday() {
	me.RemoveDay(Monday)
}
func (me Dow) RemoveTuesday() {
	me.RemoveDay(Tuesday)
}
func (me Dow) RemoveWednesday() {
	me.RemoveDay(Wednesday)
}
func (me Dow) RemoveThursday() {
	me.RemoveDay(Thursday)
}
func (me Dow) RemoveFriday() {
	me.RemoveDay(Friday)
}
func (me Dow) RemoveSaturday() {
	me.RemoveDay(Saturday)
}
func (me Dow) RemoveSunday() {
	me.RemoveDay(Sunday)
}

func (me *Dow) Clear() {
	*me = 0
}

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
	*me &= other
	return nil
}

func (me Dow) Equal(o Dow) bool {
	return me == o
}

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

func (me Dow) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.String())
}

func (me *Dow) UnmarshalJSON(d []byte) error {
	var s string
	if err := json.Unmarshal(d, &s); err != nil {
		return err
	}
	return me.Parse(s)
}
