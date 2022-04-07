package errors

import (
	"encoding/json"
	"fmt"
	"testing"
)

var err = fmt.Errorf("test error")

func TestMarshalJSON(t *testing.T) {
	IncludeBacktrace = true
	e := Newf("test")

	b, _ := json.Marshal(e)

	t.Logf("JSON: %s", b)
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
