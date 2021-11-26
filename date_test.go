package date

import (
	"testing"
	"time"
)

const message = "\033[31m`%v` was expected, but it is `%v`\033[0m"

func TestYear(t *testing.T) {
	subject := NewDate(2021, 12, 31).Year()

	expected := 2021

	if subject != expected {
		t.Errorf(message, subject, expected)
	}
}

func TestMonth(t *testing.T) {
	subject := NewDate(2021, 12, 31).Month()

	expected := 12

	if subject != expected {
		t.Errorf(message, subject, expected)
	}
}

func TestDay(t *testing.T) {
	subject := NewDate(2021, 12, 31).Day()

	expected := 31

	if subject != expected {
		t.Errorf(message, subject, expected)
	}
}

func TestIsEqual(t *testing.T) {
	var subject, expected Date

	subject = NewDate(2021, 12, 31)

	expected = NewDate(2021, 12, 31)

	if subject.IsEqual(expected) == false {
		t.Errorf(message, subject, expected)
	}

	subject = NewDate(2021, 12, 31)

	expected = NewDate(2022, 12, 31)

	if subject.IsEqual(expected) == true {
		t.Errorf(message, subject, expected)
	}
}

func TestToday(t *testing.T) {
	year, month, day := time.Now().Date()

	subject := Today()

	expected := NewDate(year, int(month), day)

	if subject.IsEqual(expected) == false {
		t.Errorf(message, subject, expected)
	}
}

func TestString(t *testing.T) {
	subject := NewDate(2021, 12, 31).String()

	expected := "2021-12-31"

	if subject != expected {
		t.Errorf(message, subject, expected)
	}
}

func TestMarshalJSON(t *testing.T) {
	subject := NewDate(2021, 12, 31).MarshalJSON()

	expected := "2021-12-31"

	if subject != expected {
		t.Errorf(message, subject, expected)
	}
}

func TestParse(t *testing.T) {
	subject, _ := Parse("2021-12-31")

	expected := NewDate(2021, 12, 31)

	if subject != expected {
		t.Errorf(message, subject, expected)
	}
}
