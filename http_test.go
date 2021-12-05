package errors

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestBadRequestf(t *testing.T) {
	errMsg := "test"
	e := BadRequestf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestForbiddenf(t *testing.T) {
	errMsg := "test"
	e := Forbiddenf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestMethodNotAllowedf(t *testing.T) {
	errMsg := "test"
	e := MethodNotAllowedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestMethodNotFoundf(t *testing.T) {
	errMsg := "test"
	e := NotFoundf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestNotImplementedf(t *testing.T) {
	errMsg := "test"
	e := NotImplementedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestNotSupportedf(t *testing.T) {
	errMsg := "test"
	e := NotSupportedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestNotValidf(t *testing.T) {
	errMsg := "test"
	e := NotValidf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestTimeoutf(t *testing.T) {
	errMsg := "test"
	e := Timeoutf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestUnauthorizedf(t *testing.T) {
	errMsg := "test"
	e := Unauthorizedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

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
	e := NewTimeout(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
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
	e := badRequest{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := forbidden{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := methodNotAllowed{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := notFound{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := notImplemented{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := notSupported{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := notValid{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := timeout{Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
	e := unauthorized{Err: Err{m: "test", l: 11, f: "random", t: []byte{0x6, 0x6, 0x6}, c: fmt.Errorf("ttt")}}
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
		Err:       Err{c: fmt.Errorf("ttt")},
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

func TestWrapWithStatus(t *testing.T) {
	t.Skipf("TODO")
}

func TestStackParse(t *testing.T) {
	stack := `goroutine 25 [running]:
runtime/debug.Stack(0x2, 0x7ab7e1, 0x1025003)
	/usr/lib/go/src/runtime/debug/stack.go:24 +0x9d
github.com/go-ap/errors.wrap(0x0, 0x0, 0xb5c3c1, 0x31, 0xc0003d4270, 0x1, 0x1, 0x0, 0x0, 0x0, ...)
	/home/build/go/pkg/mod/github.com/go-ap/errors@v0.0.0-20191123201507-86232ca294a2/errors.go:90 +0x16a
github.com/go-ap/errors.Annotatef(...)
	/home/build/go/pkg/mod/github.com/go-ap/errors@v0.0.0-20191123201507-86232ca294a2/errors.go:48
github.com/go-ap/errors.wrapErr(0x0, 0x0, 0xb5c3c1, 0x31, 0xc0003d4270, 0x1, 0x1, 0x0, 0x0, 0x0, ...)
	/home/build/go/pkg/mod/github.com/go-ap/errors@v0.0.0-20191123201507-86232ca294a2/http.go:54 +0xce
github.com/go-ap/errors.Unauthorizedf(...)
	/home/build/go/pkg/mod/github.com/go-ap/errors@v0.0.0-20191123201507-86232ca294a2/http.go:145
github.com/go-ap/fedbox/validation.genericValidator.ValidateClientActivity(0xc000120cc0, 0x15, 0xc000045c00, 0xc3d3a0, 0xc0000a6158, 0x7f51da510288, 0xc0000c43c0, 0xc4efc0, 0xc0000d2000, 0xc00020f220, ...)
	/home/build/fedbox/validation/validation.go:216 +0x846
github.com/go-ap/fedbox/app.HandleRequest(0xb407de, 0x6, 0xc00018ed00, 0xc570c0, 0xc0000c43c0, 0x0, 0xaea360, 0xc0003d4ae8, 0x40e0e6, 0xc00042ae00)
	/home/build/fedbox/app/handlers.go:102 +0x533
github.com/go-ap/handlers.ActivityHandlerFn.ServeHTTP(0xb6cdf0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ed00)
	/home/build/go/pkg/mod/github.com/go-ap/handlers@v0.0.0-20191124120223-6d767f8fa46e/handlers.go:79 +0x117
github.com/go-ap/fedbox/app.Validator.func1.1(0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/fedbox/app/middleware.go:34 +0x1d2
net/http.HandlerFunc.ServeHTTP(0xc0002149c0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*ChainHandler).ServeHTTP(0xc0000cdb80, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/chain.go:31 +0x52
github.com/go-chi/chi.(*Mux).routeHTTP(0xc0000a92c0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:425 +0x278
net/http.HandlerFunc.ServeHTTP(0xc00008df90, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).ServeHTTP(0xc0000a92c0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:70 +0x513
github.com/go-chi/chi.(*Mux).Mount.func1(0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:292 +0x124
net/http.HandlerFunc.ServeHTTP(0xc00009bb20, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).routeHTTP(0xc0000a9260, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:425 +0x278
net/http.HandlerFunc.ServeHTTP(0xc00008df80, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).ServeHTTP(0xc0000a9260, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:70 +0x513
github.com/go-chi/chi.(*Mux).Mount.func1(0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:292 +0x124
net/http.HandlerFunc.ServeHTTP(0xc00009bbc0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*ChainHandler).ServeHTTP(0xc000232a40, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/chain.go:31 +0x52
github.com/go-chi/chi.(*Mux).routeHTTP(0xc0000a90e0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:425 +0x278
net/http.HandlerFunc.ServeHTTP(0xc00008df50, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).ServeHTTP(0xc0000a90e0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:70 +0x513
github.com/go-chi/chi.(*Mux).Mount.func1(0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:292 +0x124
net/http.HandlerFunc.ServeHTTP(0xc00009bc60, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).routeHTTP(0xc0000a9080, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:425 +0x278
net/http.HandlerFunc.ServeHTTP(0xc00008df40, 0x7f51da3712c8, 0xc0002f0440, 0xc00018ea00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-ap/fedbox/app.ActorFromAuthHeader.func1.1(0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/home/build/fedbox/app/middleware.go:60 +0x40d
net/http.HandlerFunc.ServeHTTP(0xc0000cd980, 0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi/middleware.GetHead.func1(0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/middleware/get_head.go:37 +0x171
net/http.HandlerFunc.ServeHTTP(0xc00009b9c0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).ServeHTTP(0xc0000a9080, 0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:70 +0x513
github.com/go-chi/chi.(*Mux).Mount.func1(0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:292 +0x124
net/http.HandlerFunc.ServeHTTP(0xc00009bea0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).routeHTTP(0xc0000a9020, 0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:425 +0x278
net/http.HandlerFunc.ServeHTTP(0xc0002341b0, 0x7f51da3712c8, 0xc0002f0440, 0xc00018e900)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi/middleware.RequestLogger.func1.1(0xc4a400, 0xc000251340, 0xc00018e000)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/middleware/logger.go:46 +0x2be
net/http.HandlerFunc.ServeHTTP(0xc0002158f0, 0xc4a400, 0xc000251340, 0xc00018e000)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-ap/fedbox/app.Repo.func1.1(0xc4a400, 0xc000251340, 0xc000309f00)
	/home/build/fedbox/app/middleware.go:21 +0x1d2
net/http.HandlerFunc.ServeHTTP(0xc000215920, 0xc4a400, 0xc000251340, 0xc000309f00)
	/usr/lib/go/src/net/http/server.go:2007 +0x44
github.com/go-chi/chi.(*Mux).ServeHTTP(0xc0000a9020, 0xc4a400, 0xc000251340, 0xc000309c00)
	/home/build/go/pkg/mod/github.com/go-chi/chi@v4.0.2+incompatible/mux.go:82 +0x2b2
net/http.serverHandler.ServeHTTP(0xc000250000, 0xc4a400, 0xc000251340, 0xc000309c00)
	/usr/lib/go/src/net/http/server.go:2802 +0xa4
net/http.(*conn).serve(0xc00022bae0, 0xc4dac0, 0xc000236440)
	/usr/lib/go/src/net/http/server.go:1890 +0x875
created by net/http.(*Server).Serve
	/usr/lib/go/src/net/http/server.go:2927 +0x38e`

	b := []byte(stack)
	st, err := parseStack(b)
	if err != nil {
		t.Errorf("Received error when parsing the stack: %s", err)
	}
	if st == nil {
		t.Errorf("Received nil Stack object when parsing the stack: %v", st)
	}

	err = os.Setenv("GOPATH", "")
	if err != nil {
		t.Errorf("Received error when setting emtpty GOPATH: %s", err)
	}
	if p := os.Getenv("GOPATH"); p != "" {
		t.Errorf("GOPATH is not empty loaded: %s", p)
	}

	st, err = parseStack(b)
	if err != nil {
		t.Errorf("Received error when parsing the stack: %s", err)
	}
	if st == nil {
		t.Errorf("Received nil Stack object when parsing the stack: %v", st)
	}
	// 41 is the number of stack elements not containing the errors package, or the runtime debug package
	validLines := 41
	if len(st) != validLines {
		t.Errorf("Stack length is incorrect %d, expected %d %q", len(st), validLines, st)
	}
	for i, se := range st {
		if strings.Contains(se.Callee, errorsPackageName) {
			t.Errorf("Stack element callee at pos %d contains error namespace %q", i, se.Callee)
		}
		if strings.Contains(se.Callee, runtimeDebugPackageName) {
			t.Errorf("Stack element callee at pos %d contains runtime debug namespace %q", i, se.Callee)
		}
		if strings.Contains(se.File, errorsPackageName) {
			t.Errorf("Stack element file name at pos %d contains error namespace %q", i, se.File)
		}
		if strings.Contains(se.File, runtimeDebugPackageName) {
			t.Errorf("Stack element file name at pos %d contains runtime debug namespace %q", i, se.File)
		}
	}
}
