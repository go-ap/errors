package errors

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"
)

var err = fmt.Errorf("test error")

func TestAnnotatef(t *testing.T) {
	testStr := "Annotatef string"
	te := Annotatef(err, testStr)
	_, ff, ll, _ := runtime.Caller(0)
	if te.c != err {
		t.Errorf("Invalid parent error %T:%s, expected %T:%s", te.c, te.c, err, err)
	}
	if te.f != ff {
		t.Errorf("Invalid file %s, expected %s", te.f, ff)
	}
	if te.l != ll-1 {
		t.Errorf("Invalid line %d, expected %d", te.l, ll-1)
	}
	if te.m != testStr {
		t.Errorf("Invalid error message %s, expected %s", te.m, testStr)
	}
}

func TestNewf(t *testing.T) {
	testStr := "Newf string"
	te := Newf(testStr)
	_, ff, ll, _ := runtime.Caller(0)
	if te.c != nil {
		t.Errorf("Invalid parent error %T:%s, expected nil", te.c, te.c)
	}
	if te.f != ff {
		t.Errorf("Invalid file %s, expected %s", te.f, ff)
	}
	if te.l != ll-1 {
		t.Errorf("Invalid line %d, expected %d", te.l, ll-1)
	}
	if te.m != testStr {
		t.Errorf("Invalid error message %s, expected %s", te.m, testStr)
	}
}

func TestErrorf(t *testing.T) {
	testStr := "Errorf string"
	err := Errorf(testStr)
	_, ff, ll, _ := runtime.Caller(0)
	if te, ok := err.(*Err); ok {
		if te.c != nil {
			t.Errorf("Invalid parent error %T:%s, expected nil", te.c, te.c)
		}
		if te.f != ff {
			t.Errorf("Invalid file %s, expected %s", te.f, ff)
		}
		if te.l != ll-1 {
			t.Errorf("Invalid line %d, expected %d", te.l, ll-1)
		}
		if te.m != testStr {
			t.Errorf("Invalid error message %s, expected %s", te.m, testStr)
		}
	} else {
		t.Errorf("Invalid error type returned %T, expected type %T", err, &Err{})
	}
}

func TestErr_As(t *testing.T) {
	e := Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}
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
		t.Errorf("%T trace should equal %T's, received %s, expected %s", e1, e, e1.t, e.t)
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
	w :=  e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}
