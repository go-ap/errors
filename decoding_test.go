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
		willErr bool
	}{
		{
			name:    "nil",
			args:    nil,
			want:    nil,
			willErr: true,
		},
		{
			name:    "empty",
			args:    []byte{'{', '}'},
			want:    []Err(nil),
			willErr: true,
		},
		{
			name:    "no errors",
			args:    []byte(`{"errors":[]}`),
			want:    []Err{},
			willErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalJSON(tt.args)
			if tt.willErr {
				if err == nil {
					t.Errorf("UnmarshalJSON() didn't return an error when it was expected")
				}
			} else {
				if err != nil {
					t.Errorf("UnmarshalJSON() returned an error = %v", err)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}
