package dayofweek

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDays(t *testing.T) {
	var newdow Dow = Monday
	newdow |= Tuesday
	newdow |= Thursday

	if !newdow.IsSet(Tuesday) {
		t.Errorf("Error, tuesday should be set here")
	}
	if newdow.IsSet(Wednesday) {
		t.Errorf("Error, Wednesday should NOT be set")
	}

	if !newdow.IsSet(Monday) {
		t.Errorf("Error, Monday should be set here")
	}
	newdow |= Saturday
	if !newdow.IsSet(Saturday) {
		t.Errorf("Error, Saturday should be set here")
	}
	newdow |= Sunday
	if !newdow.IsSet(Sunday) {
		t.Errorf("Error, Sunday should be set here")
	}
}

func TestTimeModule(t *testing.T) {
	var newdow Dow
	if newdow.IsWeekday(time.Sunday) {
		t.Errorf("Error, time.Sunday should not be set")
	}
	if newdow.IsWeekday(time.Monday) {
		t.Errorf("Error, time.Monday should not be set")
	}
	if newdow.IsWeekday(time.Tuesday) {
		t.Errorf("Error, time.Tuesday should not be set")
	}
	if newdow.IsWeekday(time.Wednesday) {
		t.Errorf("Error, time.Wednesday should not be set")
	}
	if newdow.IsWeekday(time.Thursday) {
		t.Errorf("Error, time.Thursday should not be set")
	}
	if newdow.IsWeekday(time.Friday) {
		t.Errorf("Error, time.Friday should not be set")
	}
	if newdow.IsWeekday(time.Saturday) {
		t.Errorf("Error, time.Saturday should not be set")
	}
	newdow |= Sunday
	if !newdow.IsWeekday(time.Sunday) {
		t.Errorf("Error, time.Sunday should be set")
	}
	newdow |= Monday
	if !newdow.IsWeekday(time.Monday) {
		t.Errorf("Error, time.Monday should be set")
	}
	newdow |= Tuesday
	if !newdow.IsWeekday(time.Tuesday) {
		t.Errorf("Error, time.Tuesday should be set")
	}
	newdow |= Wednesday
	if !newdow.IsWeekday(time.Wednesday) {
		t.Errorf("Error, time.Wednesday should be set")
	}
	newdow |= Thursday
	if !newdow.IsWeekday(time.Thursday) {
		t.Errorf("Error, time.Thursday should be set")
	}
	newdow |= Friday
	if !newdow.IsWeekday(time.Friday) {
		t.Errorf("Error, time.Friday should be set")
	}
	if newdow.IsWeekday(time.Saturday) {
		t.Errorf("Error, time.Saturday should not be set YET")
	}
	newdow |= Saturday
	if !newdow.IsWeekday(time.Saturday) {
		t.Errorf("Error, time.Saturday should be set")
	}
}

func TestParse(t *testing.T) {
	var newdow Dow
	newdow.Parse("Monday, Tuesday, Thursday, Sunday")
	if !newdow.IsSet(Tuesday) {
		t.Errorf("Parsed string should contain Tuesday")
	}
	if newdow.IsSet(Wednesday) {
		t.Errorf("Parsed string should NOT contain Wednesday")
	}
	if !newdow.IsSet(Sunday) {
		t.Errorf("Parsed string should contain Sunday")
	}
	newdow.Clear()
	newdow.Parse("Mon, tue,thu")
	if !newdow.IsSet(Thursday) {
		t.Errorf("Parsed string should contain Thursday")
	}
	if newdow.IsSet(Sunday) {
		t.Errorf("Parsed string should not contain Sunday - was cleared")
	}
	if !newdow.IsSet(Tuesday) {
		t.Errorf("Parsed string should contain Tuesday")
	}
}

func TestJSONParse(t *testing.T) {
	teststring := []byte(`{"runon": "Mon, Tuesday,thu,sat"}`)
	var into struct {
		RunOn Dow
	}
	if err := json.Unmarshal(teststring, &into); err != nil {
		t.Errorf("json Unmarshal Error (external): %s\n", err)
	}
	if !into.RunOn.IsSet(Tuesday) {
		t.Errorf("JSON Unmarshal, Tuesday should be set!")
	}
	if into.RunOn.IsSet(Wednesday) {
		t.Errorf("JSON Unmarshal, Wednesday should NOT be set!")
	}
	if !into.RunOn.IsSaturday() {
		t.Errorf("JSON Unmarshal, Saturday should be set!")
	}
}

func TestEquality(t *testing.T) {
	var one, two, three Dow
	one.Parse("Mon, sat, tue,Thursday")
	two.Parse("Monday, Tuesday, Thu, Sat")
	three.Parse("Mon, Tue, Wed, Thu, Sat")
	if one != two {
		t.Errorf("Equality, One string and Two strings should match")
	}
	if one == three {
		t.Errorf("Equality, One and Three should NOT match")
	}
}

func TestRemove(t *testing.T) {
	var one, two, three Dow
	one.Parse("Mon, sat, tue,Thursday")
	two.Parse("Monday, Tuesday, Thu, Sat")
	three.Parse("Monday, Thu, Sat")
	if one != two {
		t.Errorf("Removal, initial setup failed, one does not equal two!")
	}
	two.RemoveDay(Tuesday)
	if two != three {
		t.Errorf("Removal Of Tuesday (Equality test) failed")
	}
	if two.IsTuesday() {
		t.Errorf("Removal Of Tuesday (Bintest) failed")
	}
}

func TestString(t *testing.T) {
	var one Dow
	one.Parse("Mon, sat, tue,Thursday")
	s := one.String()
	if s != "Monday, Tuesday, Thursday, Saturday" {
		t.Errorf("Stringify failed, result=%s\n", s)
	}
}

func TestMarshall(t *testing.T) {
	var one Dow
	one.Parse("Mon, sat, tue,Thursday")
	s, err := one.MarshalJSON()
	if err != nil {
		t.Errorf("Error marshalling (%v): %s\n", one, err)
	}
	test := "\"Monday, Tuesday, Thursday, Saturday\""
	if string(s) != test {
		t.Errorf("JSON Marshal failed, result=%s\n", s)
	}
}

func TestUnMarshall(t *testing.T) {
	datestring := "Mon, Tue, Wednesday, Sun, Friday"
	testjson := "{\"coverage\": \"" + datestring + "\"}"
	var into map[string]Dow
	err := json.Unmarshal([]byte(testjson), &into)
	if err != nil {
		t.Errorf("Unmarshal Test failed: %s\n", err)
	}
	var test Dow
	test.Parse(datestring)
	if into["coverage"] != test {
		t.Errorf("Error unmarshalling! unmarshalled=%v != test=%v\n", into["coverage"], test)
	}
}

func ExampleSet() {
	var x Dow
	x.Set(true, false, true, false, true, false, true)
	fmt.Printf("%t\n", x.IsSet(Monday))
	fmt.Printf("%t\n", x.IsSet(Tuesday))
	// Output:
	// true
	// false
}

func ExampleRemoveDay() {
	x := New(true, false, true, false, true, false, true)
	x.RemoveDay(Wednesday)

	fmt.Printf("%s", x.String())
	// Output: Monday, Friday, Sunday
}

func ExampleParse() {
	var x Dow
	s := "Mon, tuesday, sun, Fri, Saturday"
	err := x.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", x.String())
	// Output: Monday, Tuesday, Friday, Saturday, Sunday
}

func ExampleParseRemove() {
	var x Dow
	s := "Mon, tuesday, sun, Fri, Saturday"
	err := x.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed: %s\n", x.String())

	err = x.ParseRemove("Sun, Sat")
	if err != nil {
		panic(err)
	}

	fmt.Printf("After Remove: %s\n", x.String())
	// Output:
	// Parsed: Monday, Tuesday, Friday, Saturday, Sunday
	// After Remove: Monday, Tuesday, Friday
}
