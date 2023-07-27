package go_utils_check

import (
	"fmt"
	"math"
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
	fmt.Println(TestIndent, "testInterface==nil ->", testInterface != nil, "Direct")
	fmt.Println(TestIndent, "testInterface==nil ->", IsNil(testInterface), "isNil")
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
				fmt.Println(TestIndent, "expected error:", tt.name, "--", err)
			}
		})
	}
}

func Test_LimitIntRange_int64(t *testing.T) {
	tests := []struct {
		value    int64
		min, max int64
		wantErr  bool
	}{
		{-1, 0, 5, true},
		{-1, -1, 5, false},
		{-1, -1, 1, false},
	}
	checkLimitIntRangeMaxMin[int](t, math.MinInt64, math.MaxInt64)
	for _, tt := range tests {
		checkLimitIntRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitIntRange_int32(t *testing.T) {
	tests := []struct {
		value    int32
		min, max int64
		wantErr  bool
	}{
		{-1, 0, 5, true},
		{-1, -1, 5, false},
		{-1, -1, 1, false},
	}
	checkLimitIntRangeMaxMin[int](t, math.MinInt32, math.MaxInt32)
	for _, tt := range tests {
		checkLimitIntRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitIntRange_int16(t *testing.T) {
	tests := []struct {
		value    int16
		min, max int64
		wantErr  bool
	}{
		{-1, 0, 5, true},
		{-1, -1, 5, false},
		{-1, -1, 1, false},
	}
	checkLimitIntRangeMaxMin[int](t, math.MinInt16, math.MaxInt16)
	for _, tt := range tests {
		checkLimitIntRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitIntRange_int8(t *testing.T) {
	tests := []struct {
		value    int8
		min, max int64
		wantErr  bool
	}{
		{-1, 0, 5, true},
		{-1, -1, 5, false},
		{-1, -1, 1, false},
	}
	checkLimitIntRangeMaxMin[int](t, math.MinInt8, math.MaxInt8)
	for _, tt := range tests {
		checkLimitIntRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitIntRange_int(t *testing.T) {
	tests := []struct {
		value    int
		min, max int64
		wantErr  bool
	}{
		{-1, 0, 5, true},
		{-1, -1, 5, false},
		{-1, -1, 1, false},
	}
	checkLimitIntRangeMaxMin[int](t, math.MinInt, math.MaxInt)
	for _, tt := range tests {
		checkLimitIntRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func checkLimitIntRangeMaxMin[T int64 | int32 | int16 | int8 | int](t *testing.T, min, max T) {
	min64, max64 := int64(min), int64(max)
	checkLimitIntRange(t, false, min, min64, max64)
	checkLimitIntRange(t, false, max, min64, max64)
	checkLimitIntRange(t, true, min, min64+1, max64)
	checkLimitIntRange(t, true, max, min64, max64-1)

}

func checkLimitIntRange[T int64 | int32 | int16 | int8 | int](t *testing.T, wantErr bool, inputValue T, min, max int64) {
	name := fmt.Sprintf("(%v,%v,%v)", inputValue, min, max)
	t.Run(name, func(t *testing.T) {
		actualValue, err := LimitIntRange(inputValue, name, min, max)
		if (err != nil) != wantErr {
			t.Errorf("LimitIntRange() error = %v, wantErr %v", err, wantErr)
		} else {
			expectedValue := inputValue
			if err != nil {
				fmt.Println(TestIndent, "expected error:", name, "--", err)
				expectedValue = 0
			}
			if actualValue != expectedValue {
				t.Errorf("LimitIntRange() expectedValue (%v != %v) actualValue", expectedValue, actualValue)
			}
		}
	})

}

// ////////////////////////////////////////////////////////////////////////////////////////////
func Test_LimitUintRange_uint64(t *testing.T) {
	tests := []struct {
		value    uint64
		min, max uint64
		wantErr  bool
	}{
		{0, 1, 5, true},
		{0, 0, 5, false},
		{1, 0, 1, false},
	}
	checkLimitUintRangeMaxMin[uint](t, 0, math.MaxUint64)
	for _, tt := range tests {
		checkLimitUintRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitUintRange_uint32(t *testing.T) {
	tests := []struct {
		value    uint32
		min, max uint64
		wantErr  bool
	}{
		{0, 1, 5, true},
		{0, 0, 5, false},
		{1, 0, 1, false},
	}
	checkLimitUintRangeMaxMin[uint](t, 0, math.MaxUint32)
	for _, tt := range tests {
		checkLimitUintRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitUintRange_uint16(t *testing.T) {
	tests := []struct {
		value    uint16
		min, max uint64
		wantErr  bool
	}{
		{0, 1, 5, true},
		{0, 0, 5, false},
		{1, 0, 1, false},
	}
	checkLimitUintRangeMaxMin[uint](t, 0, math.MaxUint16)
	for _, tt := range tests {
		checkLimitUintRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitUintRange_uint8(t *testing.T) {
	tests := []struct {
		value    uint8
		min, max uint64
		wantErr  bool
	}{
		{0, 1, 5, true},
		{0, 0, 5, false},
		{1, 0, 1, false},
	}
	checkLimitUintRangeMaxMin[uint](t, 0, math.MaxUint8)
	for _, tt := range tests {
		checkLimitUintRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func Test_LimitUintRange_uint(t *testing.T) {
	tests := []struct {
		value    uint
		min, max uint64
		wantErr  bool
	}{
		{0, 1, 5, true},
		{0, 0, 5, false},
		{1, 0, 1, false},
	}
	checkLimitUintRangeMaxMin[uint](t, 0, math.MaxUint)
	for _, tt := range tests {
		checkLimitUintRange(t, tt.wantErr, tt.value, tt.min, tt.max)
	}
}

func checkLimitUintRangeMaxMin[T uint64 | uint32 | uint16 | uint8 | uint](t *testing.T, min, max T) {
	min64, max64 := uint64(min), uint64(max)
	checkLimitUintRange(t, false, min, min64, max64)
	checkLimitUintRange(t, false, max, min64, max64)
	checkLimitUintRange(t, true, min, min64+1, max64)
	checkLimitUintRange(t, true, max, min64, max64-1)

}

func checkLimitUintRange[T uint64 | uint32 | uint16 | uint8 | uint](t *testing.T, wantErr bool, inputValue T, min, max uint64) {
	name := fmt.Sprintf("(%v,%v,%v)", inputValue, min, max)
	t.Run(name, func(t *testing.T) {
		actualValue, err := LimitUintRange(inputValue, name, min, max)
		if (err != nil) != wantErr {
			t.Errorf("LimitUintRange() error = %v, wantErr %v", err, wantErr)
		} else {
			expectedValue := inputValue
			if err != nil {
				fmt.Println(TestIndent, "expected error:", name, "--", err)
				expectedValue = 0
			}
			if actualValue != expectedValue {
				t.Errorf("LimitUintRange() expectedValue (%v != %v) actualValue", expectedValue, actualValue)
			}
		}
	})

}
