package openai

import (
	"errors"
	"io"
	"testing"
)

func TestParseEvents(t *testing.T) {
	var tcs = []struct {
		in  []byte
		out [][]byte
		err error
	}{
		{
			in:  []byte(`bad-prefix: {}`),
			err: ErrBadPrefix,
		},
		{
			in: []byte(`data: {}`),
			out: [][]byte{
				[]byte(`data: {}`),
			},
		},
		{
			in: []byte(`data: {}

data: {}

`),
			out: [][]byte{
				[]byte(`data: {}`),
				[]byte(`data: {}`),
			},
		},
		{
			in: []byte(`data: {}

data: [DONE]

`),
			out: [][]byte{
				[]byte(`data: {}`),
			},
			err: io.EOF,
		},
	}

	for _, tc := range tcs {
		var out, err = parseEvents(tc.in)
		if len(out) != len(tc.out) {
			t.Fatal("mismatched event counts")
		}
		if !errors.Is(err, tc.err) {
			t.Fatalf("expected err=%v, got err=%v", tc.err, err)
		}
	}
}
