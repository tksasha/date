package date_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/tksasha/date"
	"gotest.tools/v3/assert"
)

func TestParse(t *testing.T) {
	res, err := date.Parse("2023-11-15")

	exp := date.New(2023, 11, 15)

	assert.NilError(t, err)
	assert.Equal(t, exp, res)
}

func TestNew(t *testing.T) {
	sbj := date.New(2023, 11, 17)

	year, month, day := sbj.Year(), int(sbj.Month()), sbj.Day()

	assert.Equal(t, year, 2023)

	assert.Equal(t, month, 11)

	assert.Equal(t, day, 17)
}

func TestEqual(t *testing.T) {
	t.Run("when it is the same", func(t *testing.T) {
		d1 := date.New(2023, 11, 17)
		d2 := date.New(2023, 11, 17)

		assert.Assert(t, d1 == d2)
	})

	t.Run("when it is not the same", func(t *testing.T) {
		d1 := date.New(2023, 11, 17)
		d2 := date.New(2023, 11, 18)

		assert.Assert(t, d1 != d2)
	})
}

func TestToday(t *testing.T) {
	year, month, day := time.Now().Date()

	exp := date.New(year, int(month), day)

	res := date.Today()

	assert.Equal(t, res, exp)
}

func TestIsEmpty(t *testing.T) {
	sbj, err := date.Parse("")

	assert.Assert(t, err != nil)
	assert.Assert(t, sbj.IsEmpty())
}

func TestString(t *testing.T) {
	res := date.New(2023, 11, 17).String()

	assert.Equal(t, res, "2023-11-17")
}

func TestMarshalJSON(t *testing.T) {
	t.Run("when it is an object", func(t *testing.T) {
		sbj := date.New(2023, 11, 17)

		res, err := json.Marshal(sbj)

		assert.NilError(t, err)
		assert.Equal(t, string(res), `"2023-11-17"`)
	})

	t.Run("when it is in a struct", func(t *testing.T) {
		item := struct {
			Date date.Date `json:"date"`
		}{date.New(2023, 11, 17)}

		res, err := json.Marshal(item)

		assert.NilError(t, err)
		assert.Equal(t, string(res), `{"date":"2023-11-17"}`)
	})
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("when it is an object", func(t *testing.T) {
		sbj := date.Date{}

		data := []byte(`"2023-11-17"`)

		err := json.Unmarshal(data, &sbj)

		exp := date.New(2023, 11, 17)

		assert.NilError(t, err)
		assert.Equal(t, sbj, exp)
	})

	t.Run("when it is in a struct", func(t *testing.T) {
		sbj := struct {
			Date date.Date `json:"date"`
		}{}

		exp := struct {
			Date date.Date `json:"date"`
		}{date.New(2023, 11, 17)}

		data := []byte(`{"date":"2023-11-17"}`)

		err := json.Unmarshal(data, &sbj)

		assert.NilError(t, err)
		assert.Equal(t, sbj, exp)
	})
}
