package errors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBadRequestf(t *testing.T) {
	const errMsg = "test"
	e := BadRequestf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestForbiddenf(t *testing.T) {
	const errMsg = "test"
	e := Forbiddenf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestMethodNotAllowedf(t *testing.T) {
	const errMsg = "test"
	e := MethodNotAllowedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestMethodNotFoundf(t *testing.T) {
	const errMsg = "test"
	e := NotFoundf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestNotImplementedf(t *testing.T) {
	const errMsg = "test"
	e := NotImplementedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestNotSupportedf(t *testing.T) {
	const errMsg = "test"
	e := NotHTTPVersionNotSupportedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestNotValidf(t *testing.T) {
	const errMsg = "test"
	e := NotAcceptablef(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestTimeoutf(t *testing.T) {
	const errMsg = "test"
	e := RequestTimeoutf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestUnauthorizedf(t *testing.T) {
	const errMsg = "test"
	e := Unauthorizedf(errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != nil {
		t.Errorf("Invalid %T parent error %T[%s], expected nil", e, e.c, e.c)
	}
}

func TestNewBadRequest(t *testing.T) {
	const errMsg = "test"
	e := NewBadRequest(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewForbidden(t *testing.T) {
	const errMsg = "test"
	e := NewForbidden(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewMethodNotAllowed(t *testing.T) {
	const errMsg = "test"
	e := NewMethodNotAllowed(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotFound(t *testing.T) {
	const errMsg = "test"
	e := NewNotFound(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotImplemented(t *testing.T) {
	const errMsg = "test"
	e := NewNotImplemented(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotSupported(t *testing.T) {
	const errMsg = "test"
	e := NewHTTPVersionNotSupported(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewNotValid(t *testing.T) {
	const errMsg = "test"
	e := NewNotAcceptable(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewTimeout(t *testing.T) {
	const errMsg = "test"
	e := NewRequestTimeout(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestNewUnauthorized(t *testing.T) {
	const errMsg = "test"
	e := NewUnauthorized(err, errMsg)
	if e.m != errMsg {
		t.Errorf("Invalid %T message %s, expected %s", e, e.m, errMsg)
	}
	if e.c != err {
		t.Errorf("Invalid %T parent error %T[%s], expected %T[%s]", e, e.c, e.c, err, err)
	}
}

func TestIsBadRequest(t *testing.T) {
	e := httpError{}
	e1 := &Err{}
	if IsBadRequest(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusBadRequest}
	if !IsBadRequest(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusBadRequest}
	if !IsBadRequest(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestBadRequest_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}, s: http.StatusBadRequest}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if !Is(e1.c, e.c) {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
	}
	if e2.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e2, e, e2.c, e2.c, e.c, e.c)
	}
	e3 := *e2
	if !e.As(&e3) {
		t.Errorf("%T should be assertable as %T", e, e3)
	}
	if e3.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e3, e, e3.m, e.m)
	}
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestBadRequest_Is(t *testing.T) {
	e := httpError{s: http.StatusBadRequest}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusBadRequest}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestBadRequest_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}, s: http.StatusBadRequest}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsForbidden(t *testing.T) {
	e := httpError{s: http.StatusForbidden}
	e1 := &Err{}
	if IsForbidden(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusForbidden}
	if !IsForbidden(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusForbidden}
	if !IsForbidden(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestForbidden_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}, s: http.StatusForbidden}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestForbidden_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestForbidden_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}, s: http.StatusForbidden}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsMethodNotAllowed(t *testing.T) {
	e := httpError{s: http.StatusMethodNotAllowed}
	e1 := &Err{}
	if IsMethodNotAllowed(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusMethodNotAllowed}
	if !IsMethodNotAllowed(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusMethodNotAllowed}
	if !IsMethodNotAllowed(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestMethodNotAllowed_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestMethodNotAllowed_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestMethodNotAllowed_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsNotFound(t *testing.T) {
	e := httpError{s: http.StatusNotFound}
	e1 := &Err{}
	if IsNotFound(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusNotFound}
	if !IsNotFound(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusNotFound}
	if !IsNotFound(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotFound_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotFound_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotFound_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsNotImplemented(t *testing.T) {
	e := httpError{s: http.StatusNotImplemented}
	e1 := &Err{}
	if IsNotImplemented(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusNotImplemented}
	if !IsNotImplemented(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusNotImplemented}
	if !IsNotImplemented(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotImplemented_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotImplemented_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotImplemented_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsHTTPVersionNotSupported(t *testing.T) {
	e := httpError{s: http.StatusHTTPVersionNotSupported}
	e1 := &Err{}
	if IsHTTPVersionNotSupported(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusHTTPVersionNotSupported}
	if !IsHTTPVersionNotSupported(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusHTTPVersionNotSupported}
	if !IsHTTPVersionNotSupported(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotSupported_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotSupported_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotSupported_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsStatusNotAcceptable(t *testing.T) {
	e := httpError{s: http.StatusNotAcceptable}
	e1 := &Err{}
	if IsStatusNotAcceptable(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusNotAcceptable}
	if !IsStatusNotAcceptable(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusNotAcceptable}
	if !IsStatusNotAcceptable(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotValid_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestNotValid_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestNotValid_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsRequestTimeout(t *testing.T) {
	e := httpError{s: http.StatusRequestTimeout}
	e1 := &Err{}
	if IsRequestTimeout(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusRequestTimeout}
	if !IsRequestTimeout(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusRequestTimeout}
	if !IsRequestTimeout(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestTimeout_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestTimeout_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestTimeout_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestIsUnauthorized(t *testing.T) {
	e := httpError{s: http.StatusUnauthorized}
	e1 := &Err{}
	if IsUnauthorized(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{s: http.StatusUnauthorized}
	if !IsUnauthorized(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{s: http.StatusUnauthorized}
	if !IsUnauthorized(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestUnauthorized_As(t *testing.T) {
	e := httpError{Err: Err{m: "test", c: fmt.Errorf("ttt")}}
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
	if !reflect.DeepEqual(e1.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e1, e, e1.t, e.t)
	}
	if e1.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e1, e, e1.c, e1.c, e.c, e.c)
	}
	e2 := &httpError{}
	if !e.As(e2) {
		t.Errorf("%T should be assertable as %T", e, e2)
	}
	if e2.m != e.m {
		t.Errorf("%T message should equal %T's, received %s, expected %s", e2, e, e2.m, e.m)
	}
	if !reflect.DeepEqual(e2.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e2, e, e2.t, e.t)
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
	if !reflect.DeepEqual(e3.t, e.t) {
		t.Errorf("%T trace should equal %T's, received %v, expected %v", e3, e, e3.t, e.t)
	}
	if e3.c != e.c {
		t.Errorf("%T parent error should equal %T's, received %T[%s], expected %T[%s]", e3, e, e3.c, e3.c, e.c, e.c)
	}
}

func TestUnauthorized_Is(t *testing.T) {
	e := httpError{}
	if e.Is(err) {
		t.Errorf("%T should not be a valid %T", err, e)
	}
	e1 := &Err{}
	if e.Is(e1) {
		t.Errorf("%T should not be a valid %T", e1, e)
	}
	e2 := &httpError{}
	if !e.Is(e2) {
		t.Errorf("%T should be a valid %T", e2, e)
	}
	e3 := httpError{}
	if !e.Is(e3) {
		t.Errorf("%T should be a valid %T", e3, e)
	}
}

func TestUnauthorized_Unwrap(t *testing.T) {
	e := httpError{Err: Err{c: fmt.Errorf("ttt")}}
	w := e.Unwrap()
	if w != e.c {
		t.Errorf("Unwrap() returned: %T[%s], expected: %T[%s]", w, w, e.c, e.c)
	}
}

func TestUnauthorized_Challenge(t *testing.T) {
	e := &httpError{Err: Err{c: fmt.Errorf("ttt")}}
	msgChallenge := "test challenge"
	e = e.Challenge(msgChallenge)

	if e.extra != msgChallenge {
		t.Errorf("Invalid challenge message for %T %s, expected %s", e, e.extra, msgChallenge)
	}
}

func TestChallenge(t *testing.T) {
	var errI error
	msgChallenge := "test challenge"
	e := httpError{
		s:     http.StatusUnauthorized,
		Err:   Err{c: fmt.Errorf("ttt")},
		extra: msgChallenge,
	}
	errI = &e
	if e.extra != msgChallenge {
		t.Errorf("Invalid challenge message for %T %s, expected %s", e, e.extra, msgChallenge)
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
