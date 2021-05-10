package any_test

import (
	"reflect"
	"testing"

	"github.com/hyqe/any"
)

func TestOne(t *testing.T) {
	type args struct {
		list []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "small list",
			args: args{
				list: []interface{}{1, "a", 1.4},
			},
		},
		{
			name: "nil",
			args: args{
				list: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := any.One(tt.args.list...)
			if got == nil && len(tt.args.list) > 0 {
				t.Errorf("One() = %v, want onr of: %v", got, tt.args.list)
			}
			if !has(tt.args.list, got) && len(tt.args.list) > 0 {
				t.Errorf("One() = %v, want onr of: %v", got, tt.args.list)
			}
		})
	}
}

func has(src []interface{}, v interface{}) bool {
	for _, x := range src {
		if reflect.DeepEqual(x, v) {
			return true
		}
	}
	return false
}
