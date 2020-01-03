package validate

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Error interface {
	Attr() string
	Message() string
	Error() string
}

type Errors map[string][]Error

func (e Errors) Add(err Error) {
	e[err.Attr()] = append(e[err.Attr()], err)
}

func (e Errors) Error() string {
	if e.Len() == 0 {
		return "No errors"
	}

	sb := &strings.Builder{}

	join := false
	for attr, errs := range e {
		for _, err := range errs {
			if join {
				sb.WriteString(" and ")
			}
			fmt.Fprintf(sb, "%s %v", attr, err)
			join = true
		}
	}

	return sb.String()
}

func (e Errors) Len() int {
	if e == nil {
		return 0
	}

	count := 0
	for _, v := range e {
		count += len(v)
	}

	return count
}

func (e Errors) MarshalJSON() ([]byte, error) {
	if len(e) == 0 {
		return []byte(`{}`), nil
	}

	m := make(map[string][]interface{}, len(e))
	for attr, errs := range e {
		errSlice := make([]interface{}, len(errs))
		for i, err := range errs {
			if jm, ok := err.(json.Marshaler); ok {
				errSlice[i] = jm
			} else {
				errSlice[i] = err.Message()
			}
		}
		m[attr] = errSlice
	}

	return json.Marshal(m)
}

func errorString(err Error) string {
	return fmt.Sprintf("%s %s", err.Attr(), err.Message())
}

type baseError struct {
	attr    string
	message string
}

func (e *baseError) Attr() string {
	return e.attr
}

func (e *baseError) Message() string {
	return e.message
}

func (e *baseError) Error() string {
	return errorString(e)
}

func NewError(attr, message string) Error {
	return &baseError{attr: attr, message: message}
}

type PresenceError struct {
	attr string
}

func (e *PresenceError) Attr() string {
	return e.attr
}

func (e *PresenceError) Message() string {
	return "cannot be blank"
}

func (e *PresenceError) Error() string {
	return errorString(e)
}

// Presence validates that value contains at least one printable, non-space rune.
func Presence(attr string, value string) *PresenceError {
	for _, r := range value {
		if unicode.IsPrint(r) && r != ' ' {
			return nil
		}
	}

	return &PresenceError{attr: attr}
}

type LengthError struct {
	attr string
	min  int
	max  int
	len  int
}

func (e *LengthError) Attr() string {
	return e.attr
}

func (e *LengthError) Message() string {
	return fmt.Sprintf("must have a length between %d and %d", e.min, e.max)
}

func (e *LengthError) Error() string {
	return errorString(e)
}

// Length validates that value contains between min and max runes.
func Length(attr string, value string, min, max int) *LengthError {
	runeCount := utf8.RuneCountInString(value)
	if runeCount < min {
		return &LengthError{attr: attr, min: min, max: max, len: runeCount}
	}

	if runeCount > max {
		return &LengthError{attr: attr, min: min, max: max, len: runeCount}
	}

	return nil
}
