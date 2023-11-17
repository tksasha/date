package date

import (
	"time"
	"fmt"
)

type Date struct{
	time.Time
}

func Parse(input string) (Date, error) {
	t, err := time.Parse(time.DateOnly, input)
	if err != nil { return Date{}, err }

	return Date{t}, nil
}

func New(year, month, day int) Date {
	return Date{time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)}
}

func Today() Date {
	year, month, day := time.Now().Date()

	return New(year, int(month), day)
}

func (date Date) String() string {
	return date.Format(time.DateOnly)
}

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", date.String())), nil
}

////////

// func (date Date) Empty() bool {
// 	return date.Equal(Date{})
// }
