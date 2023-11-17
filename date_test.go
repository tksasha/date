package date

import (
	"testing"
	"time"

	"os"
	"encoding/json"
	"strings"
)

const M = "\033[31m`%v` was expected, but it is `%v`\033[0m"

func TestParse(t *testing.T) {
	res, err := Parse("2023-11-15")
	if err != nil {
		t.Fatal(err)
	}

	exp := New(2023, 11, 15)

	if exp != res {
		t.Errorf(M, exp, res)
	}
}

func TestNew(t *testing.T) {
	res := New(2023, 11, 17)

	if 2023 != res.Year() {
		t.Errorf(M, 2023, res.Year())
	}

	if 11 != res.Month() {
		t.Errorf(M, 11, res.Month())
	}

	if 17 != res.Day() {
		t.Errorf(M, 17, res.Day())
	}
}

func TestEqual(t *testing.T) {
	t.Run("when it is the same", func(t *testing.T) {
		d1 := New(2023, 11, 17)
		d2 := New(2023, 11, 17)

		if d1 != d2 {
			t.Errorf(M, true, false)
		}
	})

	t.Run("when it is not the same", func(t *testing.T) {
		d1 := New(2023, 11, 17)
		d2 := New(2023, 11, 18)

		if d1 == d2 {
			t.Errorf(M, false, true)
		}
	})
}

func TestToday(t *testing.T) {
	year, month, day := time.Now().Date()

	exp := New(year, int(month), day)

	res := Today()

	if exp != res {
		t.Errorf(M, exp, res)
	}
}

func TestString(t *testing.T) {
	exp := "2023-11-17"

	res := New(2023, 11, 17).String()

	if exp != res {
		t.Errorf(M, exp, res)
	}
}

func TestMarshalJSON(t *testing.T) {
	exp := `"2023-11-17"`

	res, _ := New(2023, 11, 17).MarshalJSON()

	if exp != string(res) {
		t.Errorf(M, exp, string(res))
	}

	t.Run("when it is in struct", func(t *testing.T) {
		item := struct {
			Date Date `json:"date"`
		}{ New(2023, 11, 17) }

		fd, _ := os.CreateTemp("", "j.json")

		defer os.Remove(fd.Name())

		json.NewEncoder(fd).Encode(item)

		exp := `{"date":"2023-11-17"}`

		data, _ := os.ReadFile(fd.Name())

		res := strings.Trim(string(data), "\n")

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})
}

// func TestEmpty(t *testing.T) {
// 	var subject, expected bool

// 	var date Date

// 	date, _ = Parse("")

// 	subject = date.Empty()

// 	expected = true

// 	if subject != expected {
// 		t.Errorf(message, subject, expected)
// 	}

// 	date, _ = Parse("2021-12-31")

// 	subject = date.Empty()

// 	expected = false

// 	if subject != expected {
// 		t.Errorf(message, subject, expected)
// 	}
// }
