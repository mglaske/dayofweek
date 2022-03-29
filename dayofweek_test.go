package dayofweek

import (
	"encoding/json"
	"testing"
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
