package params

import (
	"reflect"
	"strconv"
	"strings"
)

// Convert converts the values in the http.Request's query string to the fields of the provided struct.
func Convert(values map[string][]string, v interface{}) error {
	rv := reflect.ValueOf(v).Elem()
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		// Skip unexported fields.
		if ft.PkgPath != "" {
			continue
		}

		tag := ft.Tag.Get("param")
		if tag == "" {
			continue
		}

		name, defaultValue := parseTag(tag)

		if values[name] == nil {
			if defaultValue != "" {
				values[name] = []string{defaultValue}
			} else {
				continue
			}
		}

		// Set the field's value.
		switch fv.Kind() {
		case reflect.String:
			fv.SetString(values[name][0])
		case reflect.Bool:
			b, err := strconv.ParseBool(values[name][0])
			if err != nil {
				return err
			}
			fv.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(values[name][0], 10, 64)
			if err != nil {
				return err
			}
			fv.SetInt(n)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			n, err := strconv.ParseUint(values[name][0], 10, 64)
			if err != nil {
				return err
			}
			fv.SetUint(n)
		case reflect.Float32, reflect.Float64:
			n, err := strconv.ParseFloat(values[name][0], fv.Type().Bits())
			if err != nil {
				return err
			}
			fv.SetFloat(n)
		}
	}

	return nil
}

func parseTag(tag string) (name string, defaultValue string) {
	parts := strings.Split(tag, ",")
	name = parts[0]
	if len(parts) > 1 {
		defaultValue = parts[1]
	}
	return name, defaultValue
}
