// +build !go1.13

package errors

const (
	homeVal = "$HOME"
	goPathVal = "$GOPATH"
)

func stack() []byte {
	goPath := os.Getenv("GOPATH")
	hPath := os.Getenv("HOME")

	sStack := debug.Stack()
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	if path.IsAbs(goPath) {
		sStack = bytes.Replace(sStack, []byte(goPath), []byte(goPathVal), -1)
	}
	if path.IsAbs(hPath) {
		sStack = bytes.Replace(sStack, []byte(hPath), []byte(homeVal), -1)
		err.f = strings.Replace(err.f, hPath, homeVal, -1)
	}
	return sStack
}
