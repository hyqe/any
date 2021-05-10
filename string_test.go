package any_test

import (
	"fmt"
	"testing"

	"github.com/hyqe/any"
)

func ExampleString() {
	x := any.String(
		any.Hex(
			any.Between(31, 52),
		),
		any.Base64(
			any.Between(10, 12),
		),
	)
	fmt.Println(x)
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		list []string
	}{
		{
			name: "abc's",
			list: []string{"a", "c", "b"},
		},
		{
			name: "none",
			list: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := any.String(tt.list...)
			if got == "" && len(tt.list) > 0 {
				t.Fatalf("got: %v, want one of: %v", got, tt.list)
			}
			if !hasString(tt.list, got) && len(tt.list) > 0 {
				t.Fatalf("got: %v, want one of: %v", got, tt.list)
			}
		})
	}
}

func TestBytes(t *testing.T) {
	if x := any.Bytes(); x == nil {
		t.Fatal("did not generate a bytes")
	}
}

func TestHex(t *testing.T) {
	if x := any.Hex(); x == "" {
		t.Fatal("did not generate a string")
	}
}

func TestBas64(t *testing.T) {
	if x := any.Base64(); x == "" {
		t.Fatal("did not generate a string")
	}
	if x := any.Base64(1, 2, 3); x == "" {
		t.Fatal("did not generate a string")
	}
}

func hasString(src []string, v string) bool {
	for _, x := range src {
		if x == v {
			return true
		}
	}
	return false
}
