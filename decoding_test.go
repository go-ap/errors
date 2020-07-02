package errors

import (
	"reflect"
	"testing"
)

func Test_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		args    []byte
		want    []Err
		wantErr bool
	}{
		{
			name:    "nil",
			args:    nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty",
			args:    []byte{'{', '}'},
			want:    []Err(nil),
			wantErr: true,
		},
		{
			name:    "no errors",
			args:    []byte(`{"errors":[]}`),
			want:    []Err{},
			wantErr: false,
		},
		{
			name: "rl-example",
			args: []byte(`{"@context":"https://fedbox.git/ns#errors","errors":[{"status":403,"message":"Error processing OAuth2 request: The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.","trace":[{"file":"/home/habarnam/workspace/go-ap/errors/errors.go","line":79,"calee":"github.com/go-ap/errors.wrap(0x114baa0, 0xc0005075c0, 0x1041334, 0x23, 0xc0003e5078, 0x1, 0x1, 0x0, 0x0, 0x0, ...)"},{"file":"/home/habarnam/workspace/go-ap/errors/errors.go","line":42,"calee":"github.com/go-ap/errors.Annotatef(0x114baa0, 0xc0005075c0, 0x1041334, 0x23, 0xc0003e5078, 0x1, 0x1, 0x0)"},{"file":"/home/habarnam/workspace/go-ap/errors/http.go","line":54,"calee":"github.com/go-ap/errors.wrapErr(0x114baa0, 0xc0005075c0, 0x1041334, 0x23, 0xc0003e5078, 0x1, 0x1, 0x0, 0x0, 0x0, ...)"},{"file":"/home/habarnam/workspace/go-ap/errors/http.go","line":130,"calee":"github.com/go-ap/errors.NewForbidden(0x114baa0, 0xc0005075c0, 0x1041334, 0x23, 0xc0003e5078, 0x1, 0x1, 0x0)"},{"file":"/home/habarnam/workspace/fedbox/app/oauth.go","line":163,"calee":"github.com/go-ap/fedbox/app.annotatedRsError(0x193, 0x114baa0, 0xc0005075c0, 0x1041334, 0x23, 0xc0003e5078, 0x1, 0x1, 0x0, 0x0)"},{"file":"/home/habarnam/workspace/fedbox/app/oauth.go","line":177,"calee":"github.com/go-ap/fedbox/app.redirectOrOutput(0xc0000a0120, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/home/habarnam/workspace/fedbox/app/oauth.go","line":83,"calee":"github.com/go-ap/fedbox/app.(*oauthHandler).Authorize(0xc00026dec0, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000206a50, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":431,"calee":"github.com/go-chi/chi.(*Mux).routeHTTP(0xc0002012c0, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000206a60, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":70,"calee":"github.com/go-chi/chi.(*Mux).ServeHTTP(0xc0002012c0, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":298,"calee":"github.com/go-chi/chi.(*Mux).Mount.func1(0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000209ec0, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":431,"calee":"github.com/go-chi/chi.(*Mux).routeHTTP(0xc000200f60, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000206890, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/home/habarnam/workspace/fedbox/app/middleware.go","line":61,"calee":"github.com/go-ap/fedbox/app.ActorFromAuthHeader.func1.1(0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cc00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000211940, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/middleware/get_head.go","line":37,"calee":"github.com/go-chi/chi/middleware.GetHead.func1(0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000209980, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/middleware/realip.go","line":34,"calee":"github.com/go-chi/chi/middleware.RealIP.func1(0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc0002099a0, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":70,"calee":"github.com/go-chi/chi.(*Mux).ServeHTTP(0xc000200f60, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":298,"calee":"github.com/go-chi/chi.(*Mux).Mount.func1(0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000209f80, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":431,"calee":"github.com/go-chi/chi.(*Mux).routeHTTP(0xc0002004e0, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000206b80, 0x7fb8c9b0baa8, 0xc00009a6c0, 0xc00021cb00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/middleware/logger.go","line":46,"calee":"github.com/go-chi/chi/middleware.RequestLogger.func1.1(0x1158d00, 0xc000010028, 0xc00021ca00)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000274ae0, 0x1158d00, 0xc000010028, 0xc00021ca00)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/middleware/request_id.go","line":76,"calee":"github.com/go-chi/chi/middleware.RequestID.func1(0x1158d00, 0xc000010028, 0xc00021c900)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000209fa0, 0x1158d00, 0xc000010028, 0xc00021c900)"},{"file":"/home/habarnam/workspace/fedbox/app/middleware.go","line":21,"calee":"github.com/go-ap/fedbox/app.Repo.func1.1(0x1158d00, 0xc000010028, 0xc00021c800)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2012,"calee":"net/http.HandlerFunc.ServeHTTP(0xc000274b10, 0x1158d00, 0xc000010028, 0xc00021c800)"},{"file":"/home/habarnam/.local/share/go/pkg/mod/github.com/go-chi/chi@v4.1.2+incompatible/mux.go","line":86,"calee":"github.com/go-chi/chi.(*Mux).ServeHTTP(0xc0002004e0, 0x1158d00, 0xc000010028, 0xc00021c800)"},{"file":"/usr/lib/go/src/net/http/server.go","line":2807,"calee":"net/http.serverHandler.ServeHTTP(0xc0002820e0, 0x1158d00, 0xc000010028, 0xc000524100)"},{"file":"/usr/lib/go/src/net/http/server.go","line":3381,"calee":"net/http.initALPNRequest.ServeHTTP(0x115d000, 0xc000090630, 0xc000080380, 0xc0002820e0, 0x1158d00, 0xc000010028, 0xc000524100)"},{"file":"/usr/lib/go/src/net/http/h2_bundle.go","line":5720,"calee":"net/http.(*http2serverConn).runHandler(0xc000582600, 0xc000010028, 0xc000524100, 0xc00029e000)"},{"file":"/usr/lib/go/src/net/http/h2_bundle.go","line":5454,"calee":"created by net/http.(*http2serverConn).processHeaders"}],"location":"/home/habarnam/workspace/go-ap/errors/http.go:54"},{"status":500,"message":"urls don't validate: https://brutalinks.git/auth/fedbox/callback / http://brutalinks.git/auth/fedbox/callback"}]}`),
			want: []Err{
				Err{
					m: "urls don't validate: https://brutalinks.git/auth/fedbox/callback / http://brutalinks.git/auth/fedbox/callback",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalJSON(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}
