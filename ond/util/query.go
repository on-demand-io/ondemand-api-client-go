package util

import (
	"fmt"
	"net/url"
	"reflect"
)

// BuildQueryParamsString creates a URL query string from a struct.
// It uses the "url" struct tag to determine the query parameter names.
// All other tags will be ignored
// It handles slices of strings in addition to other basic types.
func BuildQueryParamsString(data any) (string, error) {
	if isNil(data) {
		return "", nil
	}

	values := url.Values{}
	val := reflect.ValueOf(data)

	// Check if data is a pointer, if so, dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Make sure data is a struct
	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("data must be a struct or a pointer to a struct")
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("url")
		if tag == "" {
			continue // Skip if no tag is present
		}

		// Check if value is the zero value of its type and skip if it is
		if value.IsZero() {
			continue
		}

		switch value.Kind() {
		case reflect.Slice:
			if value.Type().Elem().Kind() == reflect.String {
				for j := 0; j < value.Len(); j++ {
					values.Add(tag, value.Index(j).String())
				}
			}
		default:
			values.Add(tag, fmt.Sprintf("%v", value.Interface()))
		}
	}

	return values.Encode(), nil
}

func isNil(a any) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}
