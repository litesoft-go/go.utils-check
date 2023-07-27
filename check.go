package go_utils_check

import (
	"errors"
	"fmt"
	"reflect"
)

var TestIndent = "             "

// IsNil -- checks 'any' thing to determine if it is directly 'nil' or
// indirectly 'nil': Interface(s), Slice(s), Map(s), Func(s), & Ptr(s)
// -
func IsNil(it any) bool {
	if it == nil {
		return true
	}
	reflectValue := reflect.ValueOf(it)
	if !reflectValue.IsValid() {
		return true
	}
	switch reflectValue.Kind() {
	case reflect.Interface, reflect.Slice, reflect.Map, reflect.Func, reflect.Ptr: // Order I think is most likely!
		return reflectValue.IsNil()
	default:
		return false
	}
}

// StringVisibleAsciiOrNonAsciiUTF8 -- implements a check of the passed in string to (simplistically) enforce
// that it consists of only visible characters.  It errors on the following conditions:
// * string is empty
// * any rune in the string is:
// ** <= ' '
// ** == 127 (DEL)
//
// Known simplistic aspect is that there is no checking for non-ascii whitespace!
// -
func StringVisibleAsciiOrNonAsciiUTF8(str string) error {
	if str == "" {
		return errors.New("was empty")
	}
	for _, r := range []rune(str) {
		if r == ' ' {
			return errors.New("contains a space")
		}
		if r < ' ' {
			return errors.New("contains an ascii control character")
		}
		if r == 127 {
			return errors.New("contains the non-visible 'DEL' (127) ascii character")
		}
	}
	return nil
}

// LimitIntRange -- determines if any kind of signed integer is 'min' <= 'value' <= 'max' (inclusive)
// Note: 'min' and 'max' are 'int64's
// -
func LimitIntRange[T int64 | int32 | int16 | int8 | int](value T, what string, min, max int64) (T, error) {
	err := checkRange(int64(value), what, min, max)
	if err != nil {
		value = 0
	}
	return value, err
}

// LimitUintRange -- determines if any kind of unsigned integer is 'min' <= 'value' <= 'max' (inclusive)
// Note: 'min' and 'max' are 'uint64's
// -
func LimitUintRange[T uint64 | uint32 | uint16 | uint8 | uint](value T, what string, min, max uint64) (T, error) {
	err := checkRange(uint64(value), what, min, max)
	if err != nil {
		value = 0
	}
	return value, err
}

func checkRange[T int64 | uint64](value T, what string, min, max T) error {
	if value < min {
		return fmt.Errorf("value (%v < %v) minimum for: %v", value, min, what)
	}
	if max < value {
		return fmt.Errorf("value (%v > %v) maximum for: %v", value, max, what)
	}
	return nil
}
