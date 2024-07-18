package util

import (
	"fmt"
	"net/url"
	"reflect"
)

// BuildQuery creates a URL query string from a struct.
// It uses the "url" struct tag to determine the query parameter names.
// All other tags will be ignored
func BuildQuery(data interface{}) (string, error) {
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

		// Get the query parameter name from the "url" tag
		tag := field.Tag.Get("url")
		if tag == "" {
			continue // Skip if no tag is present
		}

		// Check if value is the zero value of its type and skip if it is
		if value.IsZero() {
			continue
		}
		values.Add(tag, fmt.Sprintf("%v", value.Interface()))
	}

	return values.Encode(), nil
}
