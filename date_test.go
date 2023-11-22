package date

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestParse(t *testing.T) {
	res, err := Parse("2023-11-15")

	assert.NilError(t, err)

	exp := New(2023, 11, 15)

	assert.Equal(t, exp, res)
}

func TestNew(t *testing.T) {
	res := New(2023, 11, 17)

	assert.Equal(t, res.Year(), 2023)

	assert.Equal(t, int(res.Month()), 11)

	assert.Equal(t, res.Day(), 17)
}

func TestEqual(t *testing.T) {
	t.Run("when it is the same", func(t *testing.T) {
		d1 := New(2023, 11, 17)
		d2 := New(2023, 11, 17)

		assert.Assert(t, d1 == d2)
	})

	t.Run("when it is not the same", func(t *testing.T) {
		d1 := New(2023, 11, 17)
		d2 := New(2023, 11, 18)

		assert.Assert(t, d1 != d2)
	})
}

func TestToday(t *testing.T) {
	year, month, day := time.Now().Date()

	exp := New(year, int(month), day)

	res := Today()

	assert.Equal(t, res, exp)
}

func TestIsEmpty(t *testing.T) {
	date, err := Parse("")

	assert.Assert(t, err != nil)

	assert.Assert(t, date.IsEmpty())
}

func TestString(t *testing.T) {
	res := New(2023, 11, 17).String()

	assert.Equal(t, res, "2023-11-17")
}

func TestMarshalJSON(t *testing.T) {
	t.Run("when it is an object", func(t *testing.T) {
		res, _ := New(2023, 11, 17).MarshalJSON()

		assert.Equal(t, string(res), `"2023-11-17"`)
	})

	t.Run("when it is in a struct", func(t *testing.T) {
		item := struct {
			Date Date `json:"date"`
		}{New(2023, 11, 17)}

		fd, _ := os.CreateTemp("", "j.json")

		defer os.Remove(fd.Name())

		json.NewEncoder(fd).Encode(item)

		exp := `{"date":"2023-11-17"}`

		data, _ := os.ReadFile(fd.Name())

		res := strings.Trim(string(data), "\n")

		assert.Equal(t, res, exp)
	})
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("when it is an object", func(t *testing.T) {
		res := &Date{}

		data := []byte(`"2023-11-17"`)

		_ = res.UnmarshalJSON(data)

		assert.Equal(t, *res, New(2023, 11, 17))
	})

	t.Run("when it is in a struct", func(t *testing.T) {
		exp := struct {
			Date Date
		}{New(2023, 11, 17)}

		res := struct {
			Date Date
		}{}

		fd, _ := os.CreateTemp("", "j.json")

		defer os.Remove(fd.Name())

		fd.Write([]byte(`{"date":"2023-11-17"}`))

		fd.Seek(0, 0)

		json.NewDecoder(fd).Decode(&res)

		assert.Equal(t, res, exp)
	})
}
