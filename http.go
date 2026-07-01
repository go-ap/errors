package errors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-ap/jsonld"
)

const errorsPackageName = "github.com/go-ap/errors"
const runtimeDebugPackageName = "runtime/debug"

type Error interface {
	error
	json.Unmarshaler
}

type httpError struct {
	Err
	s     int
	extra string
}

func (h *httpError) Error() string {
	return h.Err.Error()
}

func (h *httpError) Is(err error) bool {
	_, ok := err.(*httpError)
	return ok || h.Err.Is(err)
}

func (h *httpError) As(err any) bool {
	switch x := err.(type) {
	case httpError:
		x = *h
	case *httpError:
		x.Err = h.Err
		x.s = h.s
		x.extra = h.extra
	case **httpError:
		*x = h
	case **Err:
		*x = &h.Err
	case *Err:
		*x = h.Err
	case Err:
		x = h.Err
	default:
		return false
	}
	return true
}

func wrapErr(err error, s string, args ...any) *Err {
	var e error
	if err != nil {
		e = Annotatef(err, s, args...)
	} else {
		e = Newf(s, args...)
	}
	asErr := new(Err)
	As(e, &asErr)
	return asErr
}

func FromResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusBadRequest {
		return nil
	}
	body := make([]byte, 0)
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)

	errors, err := UnmarshalJSON(body)
	if err != nil {
		return err
	}
	if len(errors) == 0 {
		return nil
	}

	return Join(errors...)
}

func AnnotateFromStatus(err error, status int, s string, args ...any) error {
	switch status {
	case http.StatusNotModified:
		return NewNotModified(err, fmt.Sprintf(s, args...))
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
		return NewNotAcceptable(err, s, args...)
	//case http.StatusProxyAuthRequired
	//case http.StatusRequestTimeout
	case http.StatusConflict:
		return NewConflict(err, s, args...)
	case http.StatusGone:
		return NewGone(err, s, args...)
	//case http.StatusLengthRequres
	//case http.StatusPreconditionFailed
	//case http.StatusRequestEntityTooLarge
	//case http.StatusRequestURITooLong
	case http.StatusUnsupportedMediaType:
		return NewUnsupportedMedia(err, s, args...)
	//case http.StatusRequestedRangeNotSatisfiable
	//case http.StatusExpectationFailed
	//case http.StatusTeapot
	//case http.StatusMisdirectedRequest
	//case http.StatusUnprocessableEntity
	//case http.StatusLocked
	//case http.StatusFailedDependency
	//case http.StatusTooEarly
	//case http.StatusTooManyRequests
	//case http.StatusRequestHeaderFieldsTooLarge
	//case http.StatusUnavailableForLegalReason
	//case http.StatusInternalServerError
	case http.StatusNotImplemented:
		return NewNotImplemented(err, s, args...)
	case http.StatusBadGateway:
		return NewBadGateway(err, s, args...)
	//case http.StatusServiceUnavailable
	//case http.StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return NewHTTPVersionNotSupported(err, s, args...)
	case http.StatusGatewayTimeout:
		return NewRequestTimeout(err, s, args...)
	}
	return Annotatef(err, s, args...)
}

func NewFromStatus(status int, s string, args ...any) error {
	switch status {
	case http.StatusBadRequest:
		return BadRequestf(s, args...)
	case http.StatusUnauthorized:
		return Unauthorizedf(s, args...)
	// http.StatusPaymentRequired
	case http.StatusForbidden:
		return Forbiddenf(s, args...)
	case http.StatusNotFound:
		return NotFoundf(s, args...)
	case http.StatusMethodNotAllowed:
		return MethodNotAllowedf(s, args...)
	case http.StatusNotAcceptable:
		return NotAcceptablef(s, args...)
	//case http.StatusProxyAuthRequired
	//case http.StatusRequestTimeout
	case http.StatusConflict:
		return Conflictf(s, args...)
	case http.StatusGone:
		return Gonef(s, args...)
	//case http.StatusLengthRequres
	//case http.StatusPreconditionFailed
	//case http.StatusRequestEntityTooLarge
	//case http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	//case http.StatusRequestedRangeNotSatisfiable
	//case http.StatusExpectationFailed
	//case http.StatusTeapot
	//case http.StatusMisdirectedRequest
	//case http.StatusUnprocessableEntity
	//case http.StatusLocked
	//case http.StatusFailedDependency
	//case http.StatusTooEarly
	//case http.StatusTooManyRequests
	//case http.StatusRequestHeaderFieldsTooLarge
	//case http.StatusUnavailableForLegalReason
	//case http.StatusInternalServerError
	case http.StatusNotImplemented:
		return NotImplementedf(s, args...)
	case http.StatusBadGateway:
		return BadGatewayf(s, args...)
	//case http.StatusServiceUnavailable
	//case http.StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return NotHTTPVersionNotSupportedf(s, args...)
	case http.StatusGatewayTimeout:
		return RequestTimeoutf(s, args...)
	}
	return Newf(s, args...)
}

func WrapWithStatus(status int, err error, s string, args ...any) error {
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
		return NewNotAcceptable(err, s, args...)
	//case http.StatusProxyAuthRequired
	//case http.StatusRequestTimeout
	case http.StatusConflict:
		return NewConflict(err, s, args...)
	case http.StatusGone:
		return NewGone(err, s, args...)
	//case http.StatusLengthRequres
	//case http.StatusPreconditionFailed
	//case http.StatusRequestEntityTooLarge
	//case http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	//case http.StatusRequestedRangeNotSatisfiable
	//case http.StatusExpectationFailed
	//case http.StatusTeapot
	//case http.StatusMisdirectedRequest
	//case http.StatusUnprocessableEntity
	//case http.StatusLocked
	//case http.StatusFailedDependency
	//case http.StatusTooEarly
	//case http.StatusTooManyRequests
	//case http.StatusRequestHeaderFieldsTooLarge
	//case http.StatusUnavailableForLegalReason
	//case http.StatusInternalServerError
	case http.StatusNotImplemented:
		return NewNotImplemented(err, s, args...)
	case http.StatusBadGateway:
		return NewBadGateway(err, s, args...)
	case http.StatusServiceUnavailable:
		return NewServiceUnavailable(err, s, args...)
	//case http.StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return NewHTTPVersionNotSupported(err, s, args...)
	case http.StatusGatewayTimeout:
		return NewRequestTimeout(err, s, args...)
	}
	return wrapErr(err, s, args...)
}

func NotFoundf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusNotFound}
}

func NewNotFound(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusNotFound}
}

func MethodNotAllowedf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusMethodNotAllowed}
}

func NewMethodNotAllowed(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusMethodNotAllowed}
}

func NotAcceptablef(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusNotAcceptable}
}

func NewNotAcceptable(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusNotAcceptable}
}

func Conflictf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusConflict}
}

func NewConflict(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusConflict}
}

func Gonef(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusGone}
}

func NewGone(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusGone}
}

func Forbiddenf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusForbidden}
}

func NewForbidden(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusForbidden}
}

func NotImplementedf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusNotImplemented}
}

func NewNotImplemented(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusNotImplemented}
}

func BadRequestf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusBadRequest}
}

func NewBadRequest(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusBadRequest}
}

func Unauthorizedf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusUnauthorized}
}

func NewUnauthorized(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusUnauthorized}
}

func UnsupportedMediaTypef(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusUnsupportedMediaType}
}

func NewUnsupportedMedia(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusGone}
}

func NotHTTPVersionNotSupportedf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusHTTPVersionNotSupported}
}

func NewHTTPVersionNotSupported(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusHTTPVersionNotSupported}
}

func RequestTimeoutf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusRequestTimeout}
}

func NewRequestTimeout(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusRequestTimeout}
}

func BadGatewayf(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusBadGateway}
}

func NewBadGateway(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusBadGateway}
}

func ServiceUnavailablef(s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(nil, s, args...), s: http.StatusServiceUnavailable}
}

func NewServiceUnavailable(e error, s string, args ...any) *httpError {
	return &httpError{Err: *wrapErr(e, s, args...), s: http.StatusServiceUnavailable}
}

func isHttpError(e error) (*httpError, bool) {
	switch err := e.(type) {
	case *httpError:
		return err, true
	default:
		return nil, false
	}
}

func isHttpErrorWithStatus(e error, st int) bool {
	err, ok := isHttpError(e)
	if ok {
		return err.s == st
	}
	err = new(httpError)
	return As(e, &err) && err.s == st
}

func IsServiceUnavailable(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusServiceUnavailable)
}

func IsBadRequest(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusBadRequest)
}

func IsForbidden(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusForbidden)
}
func IsHTTPVersionNotSupported(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusHTTPVersionNotSupported)
}
func IsConflict(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusConflict)
}
func IsGone(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusGone)
}
func IsUnsupportedMediaType(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusUnsupportedMediaType)
}

func IsMethodNotAllowed(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusMethodNotAllowed)
}

func IsNotFound(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusNotFound)
}

func IsNotImplemented(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusNotImplemented)
}

func IsUnauthorized(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusUnauthorized)
}

func IsRequestTimeout(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusRequestTimeout)
}

func IsStatusNotAcceptable(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusNotAcceptable)
}

func IsBadGateway(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusBadGateway)
}

// Challenge adds a challenge token to be added to the HTTP response
func (h *httpError) Challenge(c string) *httpError {
	h.extra = c
	return h
}

// Challenge returns the challenge of the err parameter if it's an httpError type error
func Challenge(err error) string {
	un := new(httpError)
	if ok := As(err, &un); ok {
		return un.extra
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
		if IsRedirect(err) {
			w.Header().Set("Location", Location(err))
			status = HttpStatus(err)
		} else {
			if status, dat = RenderErrors(r, err); status == 0 {
				status = http.StatusInternalServerError
			}
			w.Header().Set("Content-Type", "application/json")
		}
	}
	w.WriteHeader(status)
	w.Write(dat)
}

func Location(err error) string {
	r := new(httpError)
	if As(err, &r) {
		return r.extra
	}
	return ""
}

// HandleError is a generic method to return an HTTP handler that passes an error up the chain
func HandleError(e error) ErrorHandlerFn {
	return func(w http.ResponseWriter, r *http.Request) error {
		return e
	}
}

// NotFound is a generic method to return an 404 error HTTP handler that
var NotFound = ErrorHandlerFn(func(w http.ResponseWriter, r *http.Request) error {
	return NotFoundf("%s not found", r.URL.Path)
})

type Http struct {
	Code    int        `jsonld:"status,omitempty"`
	Message string     `jsonld:"message"`
	Trace   StackTrace `jsonld:"trace,omitempty"`
}

func HttpErrors(err error) []Http {
	https := make([]Http, 0)

	withBacktrace := includeBacktrace.Load()
	load := func(err error) Http {
		var trace StackTrace
		var msg string
		if e, ok := err.(*Err); ok {
			msg = e.Error()
			if withBacktrace {
				trace = e.StackTrace()
			}
		} else {
			local := new(Err)
			if ok := As(err, &local); ok {
				if withBacktrace {
					trace = local.StackTrace()
				}
			}
			msg = err.Error()
		}

		return Http{
			Message: msg,
			Trace:   trace,
			Code:    HttpStatus(err),
		}
	}
	https = append(https, load(err))
	for {
		if err = Unwrap(err); err != nil {
			https = append(https, load(err))
		} else {
			break
		}
	}

	return https
}

func HttpStatus(e error) int {
	if IsRedirect(e) {
		r := new(httpError)
		if As(e, &r) {
			return r.s
		}
	}
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
	if IsStatusNotAcceptable(e) {
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
	//  http.StatusLengthRequires
	//  http.StatusPreconditionFailed
	//  http.StatusRequestEntityTooLarge
	//  http.StatusRequestURITooLong
	if IsUnsupportedMediaType(e) {
		return http.StatusUnsupportedMediaType
	}
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
	if IsServiceUnavailable(e) {
		return http.StatusServiceUnavailable
	}
	//  http.StatusGatewayTimeout
	if IsHTTPVersionNotSupported(e) {
		return http.StatusHTTPVersionNotSupported
	}

	if IsRequestTimeout(e) {
		return http.StatusGatewayTimeout
	}

	return 0
}

func errorFromStatus(status int) Error {
	err := new(httpError)
	err.s = status
	return err
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
		more := HttpErrors(err)
		errMap = append(errMap, more...)
		status = HttpStatus(err)
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
