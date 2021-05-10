package any_test

import (
	"fmt"
	"testing"

	"github.com/hyqe/any"
)

func ExampleInt() {
	x := any.Int(
		any.Between(10, 20),
		any.Between(30, 40),
	)
	fmt.Println(x)
}

func TestInt(t *testing.T) {
	if x := any.Int(); x == 0 {
		t.Fatal("did not generate a int")
	}
	if x := any.Int(1, 2, 3); x == 0 {
		t.Fatal("did not generate a int")
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "high",
			args: args{
				min: 5000000,
				max: 5000001,
			},
		},
		{
			name: "eq",
			args: args{
				min: 3,
				max: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := any.Between(tt.args.min, tt.args.max); tt.args.min > got || tt.args.max < got {
				t.Errorf("Between() = %v, want %v to %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}
