package modifiers

// Modifiers is an interface that defines methods for working with a range of values.
// It provides methods for getting the minimum and maximum range values, the default value,
// and checking if a given value is within the range.
type Modifiers interface {
	Contain(s string) bool
}

type Month string
type Day string

const (
	All Month = "*"
	Jan Month = "JAN"
	Feb Month = "FEB"
	Mar Month = "MAR"
	Apr Month = "APR"
	May Month = "MAY"
	Jun Month = "JUN"
	Jul Month = "JUL"
	Aug Month = "AUG"
	Sep Month = "SEP"
	Oct Month = "OCT"
	Nov Month = "NOV"
	Dec Month = "DEC"
)

const (
	Mon Day = "MON"
	Tue Day = "TUE"
	Wed Day = "WED"
	Thu Day = "THU"
	Fri Day = "FRI"
	Sat Day = "SAT"
	Sun Day = "SUN"
)

var Months = []Month{All, Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec}
var Days = []Day{Mon, Tue, Wed, Thu, Fri, Sat, Sun}
