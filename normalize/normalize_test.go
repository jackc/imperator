package normalize_test

import (
	"testing"

	"github.com/jackc/imperator/normalize"
	"github.com/stretchr/testify/assert"
)

func TestTextField(t *testing.T) {
	tests := []struct {
		in  string
		out string
		msg string
	}{
		{in: "a", out: "a"},
		{in: " a", out: "a"},
		{in: "a ", out: "a"},
		{in: " a ", out: "a"},
		{in: "a\xfe\xffa", out: "aa", msg: "invalid UTF-8"},
		{in: "a\u200Ba", out: "a a", msg: "replace non-normal spaces"},
		{in: "a\ta", out: "a a", msg: "replace control character"},
		{in: "a\r\n", out: "a", msg: "trim happens after replaced control character"},
	}

	for i, tt := range tests {
		assert.Equal(t, tt.out, normalize.TextField(tt.in), "%d: %s %s", i, tt.in, tt.msg)
	}
}
