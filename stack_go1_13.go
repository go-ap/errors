// +build go1.13

package errors

import "runtime/debug"

// For >= Go1.13 we rely on passing the -trimpath flag to the build flags
func stack() []byte {
	return debug.Stack()
}
