package errors

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ap/jsonld"
	"net/http"
	"strconv"
)

type Error interface {
	error
	json.Unmarshaler
}

type notFound struct {
	Err
}

type methodNotAllowed struct {
	Err
}

type notValid struct {
	Err
}

type forbidden struct {
	Err
}

type notImplemented struct {
	Err
}

type conflict struct {
	Err
}

type gone struct {
	Err
}

type badRequest struct {
	Err
}

type unauthorized struct {
	Err
	challenge string
}

type notSupported struct {
	Err
}

type timeout struct {
	Err
}

type badGateway struct {
	Err
}

func wrapErr(err error, s string, args ...interface{}) Err {
	e := Annotatef(err, s, args...)
	asErr := Err{}
	errors.As(e, &asErr)
	return asErr
}

func WrapWithStatus(status int, err error, s string, args ...interface{}) error {
	switch status {
	case http.StatusBadRequest:
		return NewBadRequest(err, s, args...)
	case http.StatusUnauthorized:
		return NewUnauthorized(err, s, args...)
	// http.StatusPaymentRequired
	case http.StatusForbidden:
		return NewForbidden(err, s, args...)
	case http.StatusNotFound:
		return NewNotFound(err, s, args...)
	case http.StatusMethodNotAllowed:
		return NewMethodNotAllowed(err, s, args...)
	case http.StatusNotAcceptable:
		return NewNotValid(err, s, args...)
	// case http.StatusProxyAuthRequired
	// case http.StatusRequestTimeout
	//  TODO(marius): http.StatusConflict
	case http.StatusConflict:
		return NewConflict(err, s, args...)
	case http.StatusGone:
		return NewGone(err, s, args...)
	// case http.StatusLengthRequres
	// case http.StatusPreconditionFailed
	// case http.StatusRequestEntityTooLarge
	// case http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	// case http.StatusRequestedRangeNotSatisfiable
	// case http.StatusExpectationFailed
	// case http.StatusTeapot
	// case http.StatusMisdirectedRequest
	// case http.StatusUnprocessableEntity
	// case http.StatusLocked
	// case http.StatusFailedDependency
	// case http.StatusTooEarly
	// case http.StatusTooManyRequests
	// case http.StatusRequestHeaderFieldsTooLarge
	// case http.StatusUnavailableForLegalReason
	// case http.StatusInternalServerError
	case http.StatusNotImplemented:
		return NewNotImplemented(err, s, args...)
	case http.StatusBadGateway:
		return NewBadGateway(err, s, args...)
	// case http.StatusServiceUnavailable
	// case http.StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return NewNotSupported(err, s, args...)
	case http.StatusGatewayTimeout:
		return NewTimeout(err, s, args...)
	}
	return wrapErr(err, s, args...)
}
func NotFoundf(s string, args ...interface{}) *notFound {
	return &notFound{wrapErr(nil, s, args...)}
}
func NewNotFound(e error, s string, args ...interface{}) *notFound {
	return &notFound{wrapErr(e, s, args...)}
}
func MethodNotAllowedf(s string, args ...interface{}) *methodNotAllowed {
	return &methodNotAllowed{wrapErr(nil, s, args...)}
}
func NewMethodNotAllowed(e error, s string, args ...interface{}) *methodNotAllowed {
	return &methodNotAllowed{wrapErr(e, s, args...)}
}
func NotValidf(s string, args ...interface{}) *notValid {
	return &notValid{wrapErr(nil, s, args...)}
}
func NewNotValid(e error, s string, args ...interface{}) *notValid {
	return &notValid{wrapErr(e, s, args...)}
}
func Conflictf(s string, args ...interface{}) *conflict {
	return &conflict{wrapErr(nil, s, args...)}
}
func NewConflict(e error, s string, args ...interface{}) *conflict {
	return &conflict{wrapErr(e, s, args...)}
}
func Gonef(s string, args ...interface{}) *gone {
	return &gone{wrapErr(nil, s, args...)}
}
func NewGone(e error, s string, args ...interface{}) *gone {
	return &gone{wrapErr(e, s, args...)}
}
func Forbiddenf(s string, args ...interface{}) *forbidden {
	return &forbidden{wrapErr(nil, s, args...)}
}
func NewForbidden(e error, s string, args ...interface{}) *forbidden {
	return &forbidden{wrapErr(e, s, args...)}
}
func NotImplementedf(s string, args ...interface{}) *notImplemented {
	return &notImplemented{wrapErr(nil, s, args...)}
}
func NewNotImplemented(e error, s string, args ...interface{}) *notImplemented {
	return &notImplemented{wrapErr(e, s, args...)}
}
func BadRequestf(s string, args ...interface{}) *badRequest {
	return &badRequest{wrapErr(nil, s, args...)}
}
func NewBadRequest(e error, s string, args ...interface{}) *badRequest {
	return &badRequest{wrapErr(e, s, args...)}
}
func Unauthorizedf(s string, args ...interface{}) *unauthorized {
	return &unauthorized{Err: wrapErr(nil, s, args...)}
}
func NewUnauthorized(e error, s string, args ...interface{}) *unauthorized {
	return &unauthorized{Err: wrapErr(e, s, args...)}
}
func NotSupportedf(s string, args ...interface{}) *notSupported {
	return &notSupported{wrapErr(nil, s, args...)}
}
func NewNotSupported(e error, s string, args ...interface{}) *notSupported {
	return &notSupported{wrapErr(e, s, args...)}
}
func Timeoutf(s string, args ...interface{}) *timeout {
	return &timeout{wrapErr(nil, s, args...)}
}
func NewTimeout(e error, s string, args ...interface{}) *timeout {
	return &timeout{wrapErr(e, s, args...)}
}
func BadGatewayf(s string, args ...interface{}) *badGateway {
	return &badGateway{wrapErr(nil, s, args...)}
}
func NewBadGateway(e error, s string, args ...interface{}) *badGateway {
	return &badGateway{wrapErr(e, s, args...)}
}
func IsBadRequest(e error) bool {
	_, okp := e.(*badRequest)
	_, oks := e.(badRequest)
	return okp || oks
}
func IsForbidden(e error) bool {
	_, okp := e.(*forbidden)
	_, oks := e.(forbidden)
	return okp || oks
}
func IsNotSupported(e error) bool {
	_, okp := e.(*notSupported)
	_, oks := e.(notSupported)
	return okp || oks
}
func IsConflict(e error) bool {
	_, okp := e.(*conflict)
	_, oks := e.(conflict)
	return okp || oks
}
func IsGone(e error) bool {
	_, okp := e.(*gone)
	_, oks := e.(gone)
	return okp || oks
}
func IsMethodNotAllowed(e error) bool {
	_, okp := e.(*methodNotAllowed)
	_, oks := e.(methodNotAllowed)
	return okp || oks
}
func IsNotFound(e error) bool {
	_, okp := e.(*notFound)
	_, oks := e.(notFound)
	return okp || oks
}
func IsNotImplemented(e error) bool {
	_, okp := e.(*notImplemented)
	_, oks := e.(notImplemented)
	return okp || oks
}
func IsUnauthorized(e error) bool {
	_, okp := e.(*unauthorized)
	_, oks := e.(unauthorized)
	return okp || oks
}
func IsTimeout(e error) bool {
	_, okp := e.(*timeout)
	_, oks := e.(timeout)
	return okp || oks
}
func IsNotValid(e error) bool {
	_, okp := e.(*notValid)
	_, oks := e.(notValid)
	return okp || oks
}

func IsBadGateway(e error) bool {
	_, okp := e.(*badGateway)
	_, oks := e.(badGateway)
	return okp || oks
}
func (n notFound) Is(e error) bool {
	return IsNotFound(e)
}
func (n notValid) Is(e error) bool {
	return IsNotValid(e)
}
func (n notImplemented) Is(e error) bool {
	return IsNotImplemented(e)
}
func (n notSupported) Is(e error) bool {
	return IsNotSupported(e)
}
func (b badRequest) Is(e error) bool {
	return IsBadRequest(e)
}
func (t timeout) Is(e error) bool {
	return IsTimeout(e)
}
func (u unauthorized) Is(e error) bool {
	return IsUnauthorized(e)
}
func (m methodNotAllowed) Is(e error) bool {
	return IsMethodNotAllowed(e)
}
func (f forbidden) Is(e error) bool {
	return IsForbidden(e)
}
func (b badGateway) Is(e error) bool {
	return IsBadGateway(e)
}
func (n notFound) Unwrap() error {
	return n.Err.Unwrap()
}
func (n notValid) Unwrap() error {
	return n.Err.Unwrap()
}
func (n notImplemented) Unwrap() error {
	return n.Err.Unwrap()
}
func (n notSupported) Unwrap() error {
	return n.Err.Unwrap()
}
func (b badRequest) Unwrap() error {
	return b.Err.Unwrap()
}
func (t timeout) Unwrap() error {
	return t.Err.Unwrap()
}
func (u unauthorized) Unwrap() error {
	return u.Err.Unwrap()
}
func (m methodNotAllowed) Unwrap() error {
	return m.Err.Unwrap()
}
func (f forbidden) Unwrap() error {
	return f.Err.Unwrap()
}
func (b badGateway) Unwrap() error {
	return b.Err.Unwrap()
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a notFound to its underlying type Err.
func (n *notFound) As(err interface{}) bool {
	switch x := err.(type) {
	case **notFound:
		*x = n
	case *notFound:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a notValid to its underlying type Err.
func (n *notValid) As(err interface{}) bool {
	switch x := err.(type) {
	case **notValid:
		*x = n
	case *notValid:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a notImplemented to its underlying type Err.
func (n *notImplemented) As(err interface{}) bool {
	switch x := err.(type) {
	case **notImplemented:
		*x = n
	case *notImplemented:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a notSupported to its underlying type Err.
func (n *notSupported) As(err interface{}) bool {
	switch x := err.(type) {
	case **notSupported:
		*x = n
	case *notSupported:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a badRequest to its underlying type Err.
func (b *badRequest) As(err interface{}) bool {
	switch x := err.(type) {
	case **badRequest:
		*x = b
	case *badRequest:
		*x = *b
	case *Err:
		return b.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a timeout to its underlying type Err.
func (t *timeout) As(err interface{}) bool {
	switch x := err.(type) {
	case **timeout:
		*x = t
	case *timeout:
		*x = *t
	case *Err:
		return t.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a unauthorized to its underlying type Err.
func (u *unauthorized) As(err interface{}) bool {
	switch x := err.(type) {
	case **unauthorized:
		*x = u
	case *unauthorized:
		*x = *u
	case *Err:
		return u.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a methodNotAllowed to its underlying type Err.
func (m *methodNotAllowed) As(err interface{}) bool {
	switch x := err.(type) {
	case **methodNotAllowed:
		*x = m
	case *methodNotAllowed:
		*x = *m
	case *Err:
		return m.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a forbidden to its underlying type Err.
func (f *forbidden) As(err interface{}) bool {
	switch x := err.(type) {
	case **forbidden:
		*x = f
	case *forbidden:
		*x = *f
	case *Err:
		return f.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//  if the underlying logic of the receiver's type can understand it.
// In this case we're converting a badGateway to its underlying type Err.
func (b *badGateway) As(err interface{}) bool {
	switch x := err.(type) {
	case **badGateway:
		*x = b
	case *badGateway:
		*x = *b
	case *Err:
		return b.Err.As(x)
	default:
		return false
	}
	return true
}

// Challenge adds a challenge token to be added to the HTTP response
func (u *unauthorized) Challenge(c string) *unauthorized {
	u.challenge = c
	return u
}

// Challenge returns the challenge of the err parameter if it's an unauthorized type error
func Challenge(err error) string {
	un := unauthorized{}
	if ok := errors.As(err, &un); ok {
		return un.challenge
	}
	return ""
}

// ErrorHandlerFn
type ErrorHandlerFn func(http.ResponseWriter, *http.Request) error

// ServeHTTP implements the http.Handler interface for the ItemHandlerFn type
func (h ErrorHandlerFn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dat []byte
	var status int

	if err := h(w, r); err != nil {
		status, dat = RenderErrors(r, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}

// HandleError is a generic method to return an HTTP handler that passes an error up the chain
func HandleError(e error) ErrorHandlerFn {
	return func(w http.ResponseWriter, r *http.Request) error {
		return e
	}
}

type Http struct {
	Code     int    `jsonld:"status"`
	Message  string `jsonld:"message"`
	Trace    *Stack `jsonld:"trace,omitempty"`
	Location string `jsonld:"location,omitempty"`
}

func HttpErrors(err error) (int, []Http) {
	https := make([]Http, 0)

	load := func(err error) Http {
		var loc string
		var trace *Stack
		var msg string
		switch e := err.(type) {
		case *Err:
			msg = fmt.Sprintf("%s", e.Error())
			if IncludeBacktrace {
				trace, _ = parseStack(e.StackTrace())
				f, l := e.Location()
				if len(f) > 0 {
					loc = fmt.Sprintf("%s:%d", f, l)
				}
			}
		default:
			local := new(Err)
			if ok := errors.As(err, local); ok {
				if IncludeBacktrace {
					trace, _ = parseStack(local.StackTrace())
					f, l := local.Location()
					if len(f) > 0 {
						loc = fmt.Sprintf("%s:%d", f, l)
					}
				}
			}
			msg = err.Error()
		}

		return Http{
			Message:  msg,
			Trace:    trace,
			Location: loc,
			Code:     httpErrorResponse(err),
		}
	}
	code := httpErrorResponse(err)
	https = append(https, load(err))
	for {
		if err = errors.Unwrap(err); err != nil {
			https = append(https, load(err))
		} else {
			break
		}
	}

	return code, https
}

func httpErrorResponse(e error) int {
	if IsBadRequest(e) {
		return http.StatusBadRequest
	}
	if IsUnauthorized(e) {
		return http.StatusUnauthorized
	}
	// http.StatusPaymentRequired
	if IsForbidden(e) {
		return http.StatusForbidden
	}
	if IsNotFound(e) {
		return http.StatusNotFound
	}
	if IsMethodNotAllowed(e) {
		return http.StatusMethodNotAllowed
	}
	if IsNotValid(e) {
		return http.StatusNotAcceptable
	}
	//  http.StatusProxyAuthRequired
	//  http.StatusRequestTimeout
	if IsConflict(e) {
		return http.StatusConflict
	}
	if IsGone(e) {
		return http.StatusGone
	}
	//  TODO(marius): http.StatusGone
	//  http.StatusLengthRequres
	//  http.StatusPreconditionFailed
	//  http.StatusRequestEntityTooLarge
	//  http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	//  http.StatusRequestedRangeNotSatisfiable
	//  http.StatusExpectationFailed
	//  http.StatusTeapot
	//  http.StatusMisdirectedRequest
	//  http.StatusUnprocessableEntity
	//  http.StatusLocked
	//  http.StatusFailedDependency
	//  http.StatusTooEarly
	//  http.StatusTooManyRequests
	//  http.StatusRequestHeaderFieldsTooLarge
	//  http.StatusUnavailableForLegalReason

	//  http.StatusInternalServerError
	if IsNotImplemented(e) {
		return http.StatusNotImplemented
	}
	if IsBadGateway(e) {
		return http.StatusBadGateway
	}
	//  http.StatusServiceUnavailable
	//  http.StatusGatewayTimeout
	if IsNotSupported(e) {
		return http.StatusHTTPVersionNotSupported
	}

	if IsTimeout(e) {
		return http.StatusGatewayTimeout
	}

	return http.StatusInternalServerError
}

func errorFromStatus(status int) Error {
	switch status {
	case http.StatusBadRequest:
		return new(badRequest)
	case http.StatusUnauthorized:
		return new(unauthorized)
	// case http.StatusPaymentRequired:
	case http.StatusForbidden:
		return new(forbidden)
	case http.StatusNotFound:
		return new(notFound)
	case http.StatusMethodNotAllowed:
		return new(methodNotAllowed)
	case http.StatusNotAcceptable:
		return new(notValid)
	// case http.StatusProxyAuthRequired:
	//  case http.StatusRequestTimeout:
	//  case shttp.StatusConflict: // TODO(marius):
	//  case http.StatusGone: // TODO(marius):
	//  case http.StatusLengthRequres:
	//  case http.StatusPreconditionFailed:
	//  case http.StatusRequestEntityTooLarge:
	//  case http.StatusRequestURITooLong:
	//  case http.StatusUnsupportedMediaType: // TODO(marius):
	//  case http.StatusRequestedRangeNotSatisfiable:
	//  case http.StatusExpectationFailed:
	//  case http.StatusTeapot:
	//  case http.StatusMisdirectedRequest:
	//  case http.StatusUnprocessableEntity:
	//  case http.StatusLocked:
	//  case http.StatusFailedDependency:
	//  case http.StatusTooEarly:
	//  case http.StatusTooManyRequests:
	//  case http.StatusRequestHeaderFieldsTooLarge:
	//  case http.StatusUnavailableForLegalReason:
	//  case http.StatusInternalServerError:
	case http.StatusNotImplemented:
		return new(notImplemented)
	case http.StatusBadGateway:
		return new(badGateway)
	// case http.StatusServiceUnavailable:
	case http.StatusHTTPVersionNotSupported:
		return new(notSupported)
	case http.StatusGatewayTimeout:
		return new(badGateway)
	case http.StatusInternalServerError:
		fallthrough
	default:
		return new(Err)
	}
}

// StackFunc is a function call in the backtrace
type StackFunc struct {
	Name    string  `jsonld:"name"`
	ArgPtrs []int64 `jsonld:"name,omitempty"`
}

// StackElement represents a stack call including file, line, and function call
type StackElement struct {
	File   string `jsonld:"file"`
	Line   int64  `jsonld:"line"`
	Callee string `jsonld:"calee,omitempty"`
	Addr   int64  `jsonld:"address,omitempty"`
}

// Stack is an array of stack elements representing the parsed relevant bits of a backtrace
// Relevant in this ctxt means, it strips the calls that are happening in the package
type Stack []StackElement

func parseStack(b []byte) (*Stack, error) {
	lvl := 2 // go up the stack call tree to hide the two internal calls
	lines := bytes.Split(b, []byte("\n"))

	if len(lines) <= lvl*2+1 {
		return nil, Newf("invalid stack trace")
	}

	skipLines := lvl * 2
	stackLen := (len(lines) - 1 - skipLines) / 2
	relLines := lines[1+skipLines:]

	stack := make(Stack, stackLen)
	for i, curLine := range relLines {
		cur := i / 2
		if cur >= len(stack) {
			break
		}
		if len(curLine) == 0 {
			continue
		}
		curStack := stack[cur]
		if i%2 == 0 {
			// function line
			curStack.Callee = string(curLine)
		} else {
			// file line
			curLine = bytes.Trim(curLine, "\t")
			elems := bytes.Split(curLine, []byte(":"))
			curStack.File = string(elems[0])

			if len(elems) > 1 {
				elems1 := bytes.Split(elems[1], []byte(" "))
				cnt := len(elems1)
				if cnt > 0 {
					curStack.Line, _ = strconv.ParseInt(string(elems1[0]), 10, 64)
				}
				if cnt > 1 {
					curStack.Addr, _ = strconv.ParseInt(string(elems1[1]), 16, 64)
				}
			}
		}
		stack[cur] = curStack
	}
	return &stack, nil
}

// TODO(marius): get a proper ctxt
func ctxt(r *http.Request) jsonld.Context {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return jsonld.Context{
		jsonld.ContextElement{
			Term: "errors",
			IRI:  jsonld.IRI(fmt.Sprintf("%s://%s/ns#errors", scheme, r.Host)),
		},
	}
}

// RenderErrors outputs the json encoded errors, with the JsonLD ctxt for current
func RenderErrors(r *http.Request, errs ...error) (int, []byte) {
	errMap := make([]Http, 0)
	var status int
	for _, err := range errs {
		code, more := HttpErrors(err)
		errMap = append(errMap, more...)
		status = code
	}
	var dat []byte
	var err error

	m := struct {
		Errors []Http `jsonld:"errors"`
	}{Errors: errMap}
	if dat, err = jsonld.WithContext(ctxt(r)).Marshal(m); err != nil {
		return http.StatusInternalServerError, dat
	}
	return status, dat
}
