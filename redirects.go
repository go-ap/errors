package errors

import "net/http"

func SeeOtherf(s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(nil, s, args...), s: http.StatusSeeOther}
}
func NewSeeOther(e error, s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(e, s, args...), s: http.StatusSeeOther}
}
func Foundf(s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(nil, s, args...), s: http.StatusFound}
}
func NewFound(e error, s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(e, s, args...), s: http.StatusFound}
}
func MovedPermanentlyf(s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(nil, s, args...), s: http.StatusMovedPermanently}
}
func NewMovedPermanently(e error, s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(e, s, args...), s: http.StatusMovedPermanently}
}
func NotModifiedf(s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(nil, s, args...), s: http.StatusNotModified}
}
func NewNotModified(e error, s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(e, s, args...), s: http.StatusNotModified}
}

func TemporaryRedirectf(s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(nil, s, args...), s: http.StatusTemporaryRedirect}
}
func NewTemporaryRedirect(e error, s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(e, s, args...), s: http.StatusTemporaryRedirect}
}
func PermanentRedirectf(s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(nil, s, args...), s: http.StatusPermanentRedirect}
}
func NewPermanentRedirect(e error, s string, args ...interface{}) *redirect {
	return &redirect{Err: wrapErr(e, s, args...), s: http.StatusPermanentRedirect}
}

type redirect struct {
	Err
	s int
}

func (f *redirect) UnmarshalJSON(data []byte) error {
	return f.Err.UnmarshalJSON(data)
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a forbidden to its underlying type Err.
func (f *redirect) As(err interface{}) bool {
	switch x := err.(type) {
	case **redirect:
		*x = f
	case *redirect:
		*x = *f
	case *Err:
		return f.Err.As(x)
	default:
		return false
	}
	return true
}

func (f redirect) Is(e error) bool {
	return IsRedirect(e)
}

func IsRedirect(e error) bool {
	_, okp := e.(*redirect)
	_, oks := e.(redirect)
	return okp || oks || As(e, &redirect{})
}

func (f redirect) Unwrap() error {
	return f.Err.Unwrap()
}
