package go_utils_check

import (
	"errors"
	"reflect"
)

var TestIndent = "             "

func IsNil(it any) bool {
	if it == nil {
		return true
	}
	itrv := reflect.ValueOf(it)
	if !itrv.IsValid() {
		return true
	}
	switch itrv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return itrv.IsNil()
	default:
		return false
	}
}

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
