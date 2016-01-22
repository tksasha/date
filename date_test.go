package date

import (
  "time"
  "testing"
  "fmt"
)

const message = "\033[32m`%v`\033[0m was expected, but it is \033[31m`%v`\033[0m"

func TestNewWithoutAnyParams(t *testing.T) {
  sbj := New()

  exp := time.Now().Format("2006-01-02")

  if sbj.String() != exp {
    t.Errorf(message, exp, sbj)
  }
}

func TestNewWithOnlyYear(t *testing.T) {
  sbj := New("1982")

  exp := fmt.Sprintf("1982-%02d-%02d", time.Now().Month(), time.Now().Day())

  if sbj.String() != exp {
    t.Errorf(message, exp, sbj)
  }
}

func TestNewWithYearAndMonth(t *testing.T) {
  sbj := New("1982", "05")

  exp := fmt.Sprintf("1982-05-%02d", time.Now().Day())

  if sbj.String() != exp {
    t.Errorf(message, exp, sbj)
  }
}

func TestNewWithYearMonthAndDay(t *testing.T) {
  sbj := New("1982", "05", "17")

  exp := "1982-05-17"

  if sbj.String() != exp {
    t.Errorf(message, exp, sbj)
  }
}

func TestNewWithString(t *testing.T) {
  sbj := New("1982-05-17")

  exp := "1982-05-17"

  if sbj.String() != exp {
    t.Errorf(message, exp, sbj)
  }
}

func TestBeginningOfMonth(t *testing.T) {
  sbj := New("1982-05-17")

  exp := "1982-05-01"

  if sbj.BeginningOfMonth().String() != exp {
    t.Errorf(message, exp, sbj)
  }
}

func TestEndOfMonth(t *testing.T) {
  sbj := New("1982-05-17")

  exp := "1982-05-31"

  if sbj.EndOfMonth().String() != exp {
    t.Errorf(message, exp, sbj)
  }
}
