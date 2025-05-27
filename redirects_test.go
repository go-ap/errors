package errors

import (
	"fmt"
	"testing"
)

func Test_redirect_Unwrap(t *testing.T) {
	type fields struct {
		c error
		u string
		s int
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   nil,
		},
		{
			name:   "fmt.Error",
			fields: fields{c: fmt.Errorf("test")},
			want:   fmt.Errorf("test"),
		},
		{
			name:   "Err",
			fields: fields{c: Errorf("test1")},
			want:   Errorf("test1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := redirect{
				c: tt.fields.c,
				u: tt.fields.u,
				s: tt.fields.s,
			}
			err := r.Unwrap()
			if (err == nil) != (tt.want == nil) {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.want)
				return
			} else if err != nil && tt.want != nil {
				if tt.want.Error() != err.Error() {
					t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.want)
				}
			}
		})
	}
}
