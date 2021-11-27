package date

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Date struct {
	time time.Time
}

func NewDate(year, month, day int) Date {
	return Date{time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)}
}

func New(year, month, day int) Date {
	return NewDate(year, month, day)
}

func (date Date) Year() int {
	return date.time.Year()
}

func (date Date) Month() int {
	return int(date.time.Month())
}

func (date Date) Day() int {
	return date.time.Day()
}

func (this Date) Equal(another Date) bool {
	if this.Year() != another.Year() {
		return false
	}

	if this.Month() != another.Month() {
		return false
	}

	if this.Day() != another.Day() {
		return false
	}

	return true
}

func Today() Date {
	year, month, day := time.Now().Date()

	return NewDate(year, int(month), day)
}

func (this Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", this.Year(), this.Month(), this.Day())
}

func (this Date) MarshalJSON() string {
	return this.String()
}

func Parse(input string) (Date, error) {
	var year, month, day int

	var err error

	parsed := regexp.
		MustCompile(`\A(\d{4})-(\d{2})-(\d{2})\z`).
		FindStringSubmatch(input)

	if len(parsed) < 4 {
		return Date{}, errors.New("date is invalid")
	}

	year, err = strconv.Atoi(parsed[1])

	if err != nil {
		return Date{}, err
	}

	month, err = strconv.Atoi(parsed[2])

	if err != nil {
		return Date{}, err
	}

	day, err = strconv.Atoi(parsed[3])

	if err != nil {
		return Date{}, err
	}

	return NewDate(year, month, day), nil
}
