package check

import (
	"fmt"
	"testing"
)

type TestEnum int

const (
	Option1 TestEnum = iota + 1
	// geInvalidTestEnum -- Not needed for this test!
)

type TestInterface interface {
	Process() error
}

func Test_IsNil(t *testing.T) {
	var testInterface TestInterface = nil
	fmt.Println("testInterface==nil ->", testInterface != nil, "Direct")
	fmt.Println("testInterface==nil ->", IsNil(testInterface), "isNil")
	tests := []struct {
		name string
		arg  any
		want bool
	}{
		{"string", "", false},
		{"int", 0, false},
		{"Option1", Option1, false},
		{"nil", nil, true},
		{"testInterface", testInterface, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.arg); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringVisibleAsciiOrNonAsciiUTF8(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{"empty", "", true},
		{"space", " ", true},
		{"a", "a", false},
		{"a-b", "a-b", false},
		{"a b", "a b", true},
		{"a{nl}b", "a\nb", true},
		{"a{31}b", string([]byte{'a', 31, 'b'}), true},
		{"a{127}b", string([]byte{'a', 127, 'b'}), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StringVisibleAsciiOrNonAsciiUTF8(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringVisibleAsciiOrNonAsciiUTF8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				fmt.Println("    error:", tt.name, "--", err)
			}
		})
	}
}
