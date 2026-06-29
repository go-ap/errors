package errors

import (
	"net/http"
)

func SeeOther(u string) *httpError {
	return &httpError{s: http.StatusSeeOther, extra: u}
}

func NewSeeOther(e error, u string) *httpError {
	return &httpError{Err: Err{c: e}, s: http.StatusSeeOther, extra: u}
}

func Found(u string) *httpError {
	return &httpError{s: http.StatusFound, extra: u}
}

func NewFound(e error, u string) *httpError {
	return &httpError{Err: Err{c: e}, s: http.StatusFound, extra: u}
}

func MovedPermanently(u string) *httpError {
	return &httpError{s: http.StatusMovedPermanently, extra: u}
}

func NewMovedPermanently(e error, u string) *httpError {
	return &httpError{Err: Err{c: e}, s: http.StatusMovedPermanently, extra: u}
}

func NotModified(u string) *httpError {
	return &httpError{s: http.StatusNotModified, extra: u}
}

func NewNotModified(e error, u string) *httpError {
	return &httpError{Err: Err{c: e}, s: http.StatusNotModified, extra: u}
}

func TemporaryRedirect(u string) *httpError {
	return &httpError{s: http.StatusTemporaryRedirect, extra: u}
}

func NewTemporaryRedirect(e error, u string) *httpError {
	return &httpError{Err: Err{c: e}, s: http.StatusTemporaryRedirect, extra: u}
}

func PermanentRedirect(u string) *httpError {
	return &httpError{s: http.StatusPermanentRedirect, extra: u}
}

func NewPermanentRedirect(e error, u string) *httpError {
	return &httpError{Err: Err{c: e}, s: http.StatusPermanentRedirect, extra: u}
}

func IsNotModified(e error) bool {
	return isHttpErrorWithStatus(e, http.StatusNotModified)
}

func IsRedirect(e error) bool {
	err, ok := isHttpError(e)
	if !ok {
		return false
	}
	return err.s >= http.StatusMultipleChoices && err.s <= http.StatusPermanentRedirect
}
