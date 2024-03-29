package errors

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var err = fmt.Errorf("test error")

func TestFormat(t *testing.T) {
	IncludeBacktrace = false
	// %+s check for unwrapped error
	e1 := Newf("test")
	str := fmt.Sprintf("%+s", e1)
	if str != e1.m {
		t.Errorf("Error message invalid %s, expected %s", str, e1.m)
	}

	// %+s check for wrapped error
	e2 := Annotatef(e1, "another")
	str = fmt.Sprintf("%+s", e2)
	val := e2.m + ": " + e2.c.Error()
	if str != val {
		t.Errorf("Error message invalid %s, expected %s", str, val)
	}

	// %v check for unwrapped error with trace
	IncludeBacktrace = true
	e3 := Newf("test1")
	str = fmt.Sprintf("%+v", e3)
	if !strings.Contains(str, e3.m) {
		t.Errorf("Error message %s\n should contain %s", str, e3.m)
	}
	if !strings.Contains(str, fmt.Sprintf("%+v", e3.t.StackTrace())) {
		t.Errorf("Error message %s\n should contain %+v", str, e3.t.StackTrace())
	}
}

func TestMarshalJSON(t *testing.T) {
	IncludeBacktrace = true
	e := Newf("test")

	b, err := e.t.StackTrace().MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON failed with error: %+s", err)
	}
	stack := make([]json.RawMessage, 0)
	err = json.Unmarshal(b, &stack)
	if err != nil {
		t.Errorf("JSON message could not be unmarshaled: %+s", err)
	}
	if len(stack) != len(e.t) {
		t.Errorf("Count of stack frames different after marshaling, expected %d got %d", len(e.t), len(stack))
	}
}

func TestAnnotatef(t *testing.T) {
	testStr := "Annotatef string"
	te := Annotatef(err, testStr)
	if te.c != err {
		t.Errorf("Invalid parent error %T:%s, expected %T:%s", te.c, te.c, err, err)
	}
	if te.m != testStr {
		t.Errorf("Invalid error message %s, expected %s", te.m, testStr)
	}
}

var homeVal = "$HOME"

func TestNewf(t *testing.T) {
	testStr := "Newf string"
	te := Newf(testStr)
	if te.c != nil {
		t.Errorf("Invalid parent error %T:%s, expected nil", te.c, te.c)
	}
	if te.m != testStr {
		t.Errorf("Invalid error message %s, expected %s", te.m, testStr)
	}
}

func TestErrorf(t *testing.T) {
	testStr := "Errorf string"
	err := Errorf(testStr)
	if te, ok := err.(*Err); ok {
		if te.c != nil {
			t.Errorf("Invalid parent error %T:%s, expected nil", te.c, te.c)
		}
		if te.m != testStr {
			t.Errorf("Invalid error message %s, expected %s", te.m, testStr)
		}
	} else {
		t.Errorf("Invalid error type returned %T, expected type %T", err, &Err{})
	}
}

/*
func TestErr_As(t *testing.T) {
	e := Err{m: "test", l: 11, f: "random", t: []uintptr{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}
	if e.As(&err) {
		t.Errorf("%T should not be assertable as %T", err, e)
	}
	type clone = Err
	e1 := clone{}
	if !e.As(&e1) {
		t.Errorf("%T should be assertable as %T", e, e1)
	}
	if e1.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e1, e, e1.m, e.m)
	}
	if e1.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e1, e, e1.l, e.l)
	}
	if e1.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e1, e, e1.f, e.f)
	}
	if !bytes.Equal(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &e
	if !e.As(&e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if e2.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e2, e, e2.l, e.l)
	}
	if e2.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e2, e, e2.f, e.f)
	}
	if !bytes.Equal(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e2, e, e2.t, e.t)
	}
	if e2.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e2, e, e2.c, e2.c, e.c, e.c)
	}
}

func TestErr_Error(t *testing.T) {
	e := Err{m: "test"}
	if e.Error() != e.m {
		t.Errorf("Error() returned %s, expected %s", e.Error(), e.m)
	}
}

func TestErr_Location(t *testing.T) {
	e := Err{l: 11, f: "random"}
	if f, l := e.Location(); l != e.l || f != e.f {
		t.Errorf("Location() returned: %s:%d, expected: %s:%d", f, l, e.f, e.l)
	}
}

func TestErr_StackTrace(t *testing.T) {
	e := Err{t: []byte{0x6, 0x6, 0x6}}
	if !bytes.Equal(e.StackTrace(), e.t) {
		t.Errorf("StackTrace() returned: %2x, expected: %2x", e.StackTrace(), e.t)
	}
}

func TestErr_Unwrap(t *testing.T) {
	e := Err{c: fmt.Errorf("ttt")}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}
*/

func TestErr_Error(t *testing.T) {
	type fields struct {
		m string
		c error
		t stack
	}
	tests := []struct {
		quiet  bool
		name   string
		fields fields
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   "",
		},
		{
			name:   "just text",
			fields: fields{m: "test"},
			want:   "test",
		},
		{
			name:   "text with single wrapped error",
			fields: fields{m: "test", c: fmt.Errorf("error")},
			want:   "test: error",
		},
		{
			name:   "text with two wrapped errors",
			fields: fields{m: "test", c: fmt.Errorf("error: %w", Newf("some error"))},
			want:   "test: error: some error",
		},
		{
			name:   "text with two wrapped errors, but no unwrapping",
			quiet:  true,
			fields: fields{m: "test", c: fmt.Errorf("error: %w", Newf("some error"))},
			want:   "test",
		},
		{
			name:   "text with two wrapped Err errors",
			fields: fields{m: "test", c: &Err{m: "error", c: &Err{m: "another error"}}},
			want:   "test: error: another error",
		},
		{
			name:   "text with two wrapped Err errors, without unwrapping",
			quiet:  true,
			fields: fields{m: "test", c: &Err{m: "error", c: &Err{m: "another error"}}},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IncludeBacktrace = tt.quiet
			e := Err{
				m: tt.fields.m,
				c: tt.fields.c,
				t: tt.fields.t,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
