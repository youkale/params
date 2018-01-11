package params

import (
	"net/url"
	"reflect"
	"errors"
	"strings"
	"strconv"
	"fmt"
)

var (
	IllegalStruct = errors.New("query: Unmarshal() expects struct input. ")
)

//Unmarshal url.Values to struct
func Unmarshal(values url.Values, s interface{}) error {
	val := reflect.ValueOf(s)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return IllegalStruct
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return IllegalStruct
	}
	return reflectValueFromTag(values, val)
}

func reflectValueFromTag(values url.Values, val reflect.Value) error {
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		kt := typ.Field(i)
		tag := kt.Tag.Get("param")
		if tag == "-" {
			continue
		}
		sv := val.Field(i)
		uv := getVal(values, tag)
		switch sv.Kind() {
		case reflect.String:
			sv.SetString(uv)
		case reflect.Bool:
			b, err := strconv.ParseBool(uv)
			if err != nil {
				return errors.New(fmt.Sprintf("convert has error, expect type: %v ,val: %v ,query key: %v", sv.Type(), uv, tag))
			}
			sv.SetBool(b)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			n, err := strconv.ParseUint(uv, 10, 64)
			if err != nil || sv.OverflowUint(n) {
				return errors.New(fmt.Sprintf("convert has error, expect type: %v ,val: %v ,query key: %v", sv.Type(), uv, tag))
			}
			sv.SetUint(n)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(uv, 10, 64)
			if err != nil || sv.OverflowInt(n) {
				return errors.New(fmt.Sprintf("convert has error, expect type: %v ,val: %v ,query key: %v", sv.Type(), uv, tag))
			}
			sv.SetInt(n)
		case reflect.Float32, reflect.Float64:
			n, err := strconv.ParseFloat(uv, sv.Type().Bits())
			if err != nil || sv.OverflowFloat(n) {
				return errors.New(fmt.Sprintf("convert has error, expect type: %v ,val: %v ,query key: %v", sv.Type(), uv, tag))
			}
			sv.SetFloat(n)
		}
	}
	return nil
}

//get val, if absent get from tag default val
func getVal(values url.Values, tag string) (string) {
	name, opts := parseTag(tag)
	uv := values.Get(name)
	optsLen := len(opts)
	if optsLen > 0 {
		if optsLen == 1 && uv == "" {
			uv = opts[0]
		}
	}
	return uv
}

type tagOptions []string

func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}
