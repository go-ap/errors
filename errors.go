package errors

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

// IncludeBacktrace is a static variable that decides if when creating an error we store the backtrace with it.
var IncludeBacktrace = true
var goPath = os.Getenv("GOPATH")
var hPath = os.Getenv("HOME")

// Err is our custom error type that can store backtrace, file and line number
type Err struct {
	c error
	m string
	t []byte
	l int
	f string
}

// Error implements the error interface
func (e Err) Error() string {
	return e.m
}

// Unwrap implements the errors.Wrapper interface
func (e Err) Unwrap() error {
	return e.c
}

// Location returns the file and line number pair of the instantiation of the error
func (e Err) Location() (string, int) {
	return e.f, e.l
}

// StackTrace returns the stack trace as returned by the debug.Stack function
func (e Err) StackTrace() []byte {
	return e.t
}

// Annotatef wraps an error with new message
func Annotatef(e error, s string, args ...interface{}) *Err {
	err := wrap(e, s, args...)
	return &err
}

// Newf creaates a new error
func Newf(s string, args ...interface{}) *Err {
	err := wrap(nil, s, args...)
	return &err
}

// Errorf is an alias for Newf
func Errorf(s string, args ...interface{}) error {
	err := wrap(nil, s, args...)
	return &err
}

// As implements support for errors.As
func (e *Err) As(err interface{}) bool {
	switch x := err.(type) {
	case **Err:
		*x = e
	case *Err:
		*x = *e
	default:
		return false
	}
	return true
}

const (
	homeVal = "$HOME"
	goPathVal = "$GOPATH"
)

func wrap(e error, s string, args ...interface{}) Err {
	err := Err{
		c: e,
		m: fmt.Sprintf(s, args...),
	}
	if IncludeBacktrace {
		skip := 2
		_, err.f, err.l, _ = runtime.Caller(skip)
		sStack := bytes.Replace(debug.Stack(), []byte(goPath), []byte(goPathVal), -1)
		sStack = bytes.Replace(sStack, []byte(hPath), []byte(homeVal), -1)
		err.t = sStack
		err.f = strings.Replace(err.f, hPath, homeVal, -1)
	}
	return err
}
