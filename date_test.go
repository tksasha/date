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

func TestIsEmpty(t *testing.T) {
	date, _ := Parse("")

	exp := true

	res := date.IsEmpty()

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
	t.Run("when it is an object", func(t *testing.T) {
		exp := `"2023-11-17"`

		res, _ := New(2023, 11, 17).MarshalJSON()

		if exp != string(res) {
			t.Errorf(M, exp, string(res))
		}
	})

	t.Run("when it is in a struct", func(t *testing.T) {
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

func TestUnmarshalJSON(t *testing.T) {
	t.Run("when it is an object", func(t *testing.T) {
		exp := New(2023, 11, 17)

		res := &Date{}

		data := []byte(`"2023-11-17"`)

		_ = res.UnmarshalJSON(data)

		if exp != *res {
			t.Errorf(M, exp, res)
		}
	})

	t.Run("when it is in a struct", func(t *testing.T) {
		exp := struct {
			Date Date
		}{ New(2023, 11, 17) }

		res := struct {
			Date Date
		}{}

		fd, _ := os.CreateTemp("", "j.json")

		defer os.Remove(fd.Name())

		fd.Write([]byte(`{"date":"2023-11-17"}`))

		fd.Seek(0, 0)

		json.NewDecoder(fd).Decode(&res)

		if exp != res {
			t.Errorf(M, exp, res)
		}
	})
}
