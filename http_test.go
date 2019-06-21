package errors

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewBadRequest(t *testing.T) {
	errMsg := "test"
	e := NewBadRequest(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewForbidden(t *testing.T) {
	errMsg := "test"
	e := NewForbidden(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewMethodNotAllowed(t *testing.T) {
	errMsg := "test"
	e := NewMethodNotAllowed(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotFound(t *testing.T) {
	errMsg := "test"
	e := NewNotFound(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotImplemented(t *testing.T) {
	errMsg := "test"
	e := NewNotImplemented(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotSupported(t *testing.T) {
	errMsg := "test"
	e := NewNotSupported(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotValid(t *testing.T) {
	errMsg := "test"
	e := NewNotValid(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewTimeout(t *testing.T) {
	errMsg := "test"
	e := NewTimeout(nil, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
}

func TestNewUnauthorized(t *testing.T) {
	errMsg := "test"
	e := NewUnauthorized(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestIsBadRequest(t *testing.T) {
	e := badRequest{}
	e1 := &Err{}
	if IsBadRequest(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &badRequest{}
	if !IsBadRequest(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := badRequest{}
	if !IsBadRequest(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestBadRequest_As(t *testing.T) {
	e := badRequest{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &badRequest{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestBadRequest_Is(t *testing.T) {
	e := badRequest{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &badRequest{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := badRequest{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestBadRequest_Unwrap(t *testing.T) {
	e := badRequest{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsForbidden(t *testing.T) {
	e := forbidden{}
	e1 := &Err{}
	if IsForbidden(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &forbidden{}
	if !IsForbidden(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := forbidden{}
	if !IsForbidden(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestForbidden_As(t *testing.T) {
	e := forbidden{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &forbidden{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestForbidden_Is(t *testing.T) {
	e := forbidden{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &forbidden{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := forbidden{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestForbidden_Unwrap(t *testing.T) {
	e := forbidden{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsMethodNotAllowed(t *testing.T) {
	e := methodNotAllowed{}
	e1 := &Err{}
	if IsMethodNotAllowed(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &methodNotAllowed{}
	if !IsMethodNotAllowed(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := methodNotAllowed{}
	if !IsMethodNotAllowed(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestMethodNotAllowed_As(t *testing.T) {
	e := methodNotAllowed{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &methodNotAllowed{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestMethodNotAllowed_Is(t *testing.T) {
	e := methodNotAllowed{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &methodNotAllowed{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := methodNotAllowed{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestMethodNotAllowed_Unwrap(t *testing.T) {
	e := methodNotAllowed{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsNotFound(t *testing.T) {
	e := notFound{}
	e1 := &Err{}
	if IsNotFound(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notFound{}
	if !IsNotFound(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notFound{}
	if !IsNotFound(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotFound_As(t *testing.T) {
	e := notFound{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &notFound{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotFound_Is(t *testing.T) {
	e := notFound{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notFound{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notFound{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotFound_Unwrap(t *testing.T) {
	e := notFound{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsNotImplemented(t *testing.T) {
	e := notImplemented{}
	e1 := &Err{}
	if IsNotImplemented(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notImplemented{}
	if !IsNotImplemented(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notImplemented{}
	if !IsNotImplemented(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotImplemented_As(t *testing.T) {
	e := notImplemented{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &notImplemented{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotImplemented_Is(t *testing.T) {
	e := notImplemented{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notImplemented{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notImplemented{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotImplemented_Unwrap(t *testing.T) {
	e := notImplemented{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsNotSupported(t *testing.T) {
	e := notSupported{}
	e1 := &Err{}
	if IsNotSupported(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notSupported{}
	if !IsNotSupported(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notSupported{}
	if !IsNotSupported(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotSupported_As(t *testing.T) {
	e := notSupported{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &notSupported{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotSupported_Is(t *testing.T) {
	e := notSupported{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notSupported{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notSupported{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotSupported_Unwrap(t *testing.T) {
	e := notSupported{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsNotValid(t *testing.T) {
	e := notValid{}
	e1 := &Err{}
	if IsNotValid(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notValid{}
	if !IsNotValid(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notValid{}
	if !IsNotValid(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotValid_As(t *testing.T) {
	e := notValid{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &notValid{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotValid_Is(t *testing.T) {
	e := notValid{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &notValid{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := notValid{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotValid_Unwrap(t *testing.T) {
	e := notValid{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsTimeout(t *testing.T) {
	e := timeout{}
	e1 := &Err{}
	if IsTimeout(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &timeout{}
	if !IsTimeout(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := timeout{}
	if !IsTimeout(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestTimeout_As(t *testing.T) {
	e := timeout{ Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &timeout{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestTimeout_Is(t *testing.T) {
	e := timeout{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &timeout{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := timeout{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestTimeout_Unwrap(t *testing.T) {
	e := timeout{Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsUnauthorized(t *testing.T) {
	e := unauthorized{}
	e1 := &Err{}
	if IsUnauthorized(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &unauthorized{}
	if !IsUnauthorized(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := unauthorized{}
	if !IsUnauthorized(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestUnauthorized_As(t *testing.T) {
	e := unauthorized{ Err: Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
	e0 := err
	if e.As(e0) {
		t.Errorf("%T should not be assertable as %T", e, e0)
	}
	e1 := Err{}
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
	e2 := &unauthorized{}
	if !e.As(e2) {
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
	e3 := e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if e3.l != e.l {
		t.Errorf("%T line should equal %T's, received %d, expected %d", e3, e, e3.l, e.l)
	}
	if e3.f != e.f {
		t.Errorf("%T file should equal %T's, received %s, expected %s", e3, e, e3.f, e.f)
	}
	if !bytes.Equal(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %2x, expected %2x", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestUnauthorized_Is(t *testing.T) {
	e := unauthorized{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &unauthorized{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := unauthorized{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestUnauthorized_Unwrap(t *testing.T) {
	e := unauthorized{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestUnauthorized_Challenge(t *testing.T) {
	e := &unauthorized{Err: Err{c: fmt.Errorf("ttt")}}
	msgChallenge := "test challenge"
	e = e.Challenge(msgChallenge)

	if e.challenge != msgChallenge {
		t.Errorf("Invalid challenge message for %T %s, expected %s", e, e.challenge, msgChallenge)
	}
}

func TestChallenge(t *testing.T) {
	var errI error
	msgChallenge := "test challenge"
	e := unauthorized{
		Err: Err{c: fmt.Errorf("ttt")},
		challenge: msgChallenge,
	}
	errI = &e
	if e.challenge != msgChallenge {
		t.Errorf("Invalid challenge message for %T %s, expected %s", e, e.challenge, msgChallenge)
	}
	ch := Challenge(errI)
	if ch != msgChallenge {
		t.Errorf("Invalid challenge message for %T %s, expected %s", errI, ch, msgChallenge)
	}
}

func TestHttpErrors(t *testing.T) {
	t.Skipf("TODO")
}

func TestErrorHandlerFn_ServeHTTP(t *testing.T) {
	t.Skipf("TODO")
}

func TestHandleError(t *testing.T) {
	t.Skipf("TODO")
}

func TestRenderErrors(t *testing.T) {
	t.Skipf("TODO")
}
