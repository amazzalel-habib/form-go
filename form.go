package formgo

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

// EncodeFormValues encodes a struct into URL form values by extracting tagged fields.
//
// The function accepts a struct and extracts fields with the "form" tag. It then
// converts non-empty string and non-zero integer fields into URL form values.
//
// Parameters:
//   - data: The input struct to be encoded.
//
// Returns:
//   - url.Values: A map of field names to their string values in URL form.
//   - error: An error if the input data is not a struct or if encoding fails.
func EncodeFormValues(data interface{}) (url.Values, error) {
	v := reflect.ValueOf(data)

	// Check if the input data is a struct
	if v.Kind() != reflect.Struct {
		return nil, errors.New("data must be a struct")
	}

	values := url.Values{}

	// Iterate through the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldName := field.Tag.Get("form")

		// Skip fields without a "form" tag
		if fieldName == "" {
			continue
		}

		fieldValue := v.Field(i)

		// Handle string fields
		if fieldValue.Kind() == reflect.String {
			if fieldValue.String() != "" {
				values.Add(fieldName, fieldValue.String())
			}
		}

		// Handle integer fields
		if fieldValue.Kind() == reflect.Int {
			if fieldValue.Int() != 0 {
				values.Add(fieldName, strconv.Itoa(int(fieldValue.Int())))
			}
		}
	}

	return values, nil
}
