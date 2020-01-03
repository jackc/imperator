package validate_test

import (
	"testing"

	"github.com/jackc/imperator/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorsMarshalJSON(t *testing.T) {
	var e validate.Errors

	// when nil
	jsonBytes, err := e.MarshalJSON()
	require.NoError(t, err)
	assert.Equal(t, []byte(`{}`), jsonBytes)

	// when empty
	e = make(validate.Errors)
	jsonBytes, err = e.MarshalJSON()
	require.NoError(t, err)
	assert.Equal(t, []byte(`{}`), jsonBytes)

	e.Add(validate.NewError("foo", "is barred"))
	jsonBytes, err = e.MarshalJSON()
	require.NoError(t, err)
	assert.Equal(t, []byte(`{"foo":["is barred"]}`), jsonBytes)
}

func TestPresence(t *testing.T) {
	tests := []struct {
		argument string
		valid    bool
	}{
		{argument: "a", valid: true},
		{argument: " a ", valid: true},
		{argument: "", valid: false},
		{argument: " ", valid: false},
		{argument: "\u200B", valid: false}, // zero-width space
	}

	for i, tt := range tests {
		verr := validate.Presence("fieldName", tt.argument)
		assert.Equal(t, tt.valid, verr == nil, "%d: %s", i, tt.argument)
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		argument string
		min      int
		max      int
		valid    bool
	}{
		{argument: "a", min: 1, max: 1, valid: true},
		{argument: "a", min: 1, max: 3, valid: true},
		{argument: "aa", min: 1, max: 3, valid: true},
		{argument: "aaa", min: 1, max: 3, valid: true},
		{argument: " a ", min: 1, max: 3, valid: true},
		{argument: "⌘⌘⌘", min: 1, max: 3, valid: true},
		{argument: " ", min: 1, max: 3, valid: true},
		{argument: "", min: 1, max: 3, valid: false},
	}

	for i, tt := range tests {
		verr := validate.Length("fieldName", tt.argument, tt.min, tt.max)
		assert.Equal(t, tt.valid, verr == nil, "%d: %s", i, tt.argument)
	}
}
