[![PkgGoDev](https://pkg.go.dev/badge/github.com/mglaske/dayofweek)](https://pkg.go.dev/github.com/mglaske/dayofweek)
[![codecov](https://codecov.io/gh/mglaske/dayofweek/branch/master/graph/badge.svg?token=9L0V9FG54Q)](https://codecov.io/gh/mglaske/dayofweek)

Day of Week Module
==================

A Golang module that allows you to perform operations on a day of week set, and
determine if specific days are set.

See the godoc for more details.

# Usage

## Instantiation

By using the constructor:
```
// This is          mon,  tue,   wed,   thu,   fri,  sat,  sun
x := dayofweek.New(true, true, false, false, false, true, true)

x.IsSet(dayofweek.Monday)
true

x.IsThursday()
false

x.IsWeekday(time.Weekday variable)
true

x.OnDate(time.Time variable)
false | true
```

Using a time.Weekday:
```
var x dayofweek.Dow
x = x.WeekdayToDow(time.Weekday variable)
```

Generating it
```
var x dayofweek.Dow
x |= dayofweek.Monday
x |= dayofweek.Sunday

or
x.AddDay(dayofweek.Thursday)

or
x.AddFriday()
x.RemoveDay(dayofweek.Monday)
x.RemoveSunday()
```

## Testing
The object contains multiple methods for testing if a specific day is set

## Comparison
The object can be compared to another object to see if the same days are set.

## JSON Marshalling
The object contains Marshall and UnMarshall functions as input / output 
of a string representation.
