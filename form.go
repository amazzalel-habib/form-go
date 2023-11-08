package formgo

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

// FloatPrecision is the precision used when formatting float values.
var FloatPrecision = -1

// EncodeFormValues encodes a struct into URL form values by extracting tagged fields.
//
// The function accepts a struct and extracts fields with the "form" tag. It then
// converts non-empty string, non-zero integer, true boolean, and non-zero float fields
// into URL form values. It also handles pointer values for these primitive types.
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

		// Handle pointer values
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()

			// Check if the pointer is nil
			if !fieldValue.IsValid() {
				continue
			}
		}

		switch fieldValue.Kind() {
		case reflect.String:
			if fieldValue.String() != "" {
				values.Add(fieldName, fieldValue.String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fieldValue.Int() != 0 {
				values.Add(fieldName, strconv.FormatInt(fieldValue.Int(), 10))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if fieldValue.Uint() != 0 {
				values.Add(fieldName, strconv.FormatUint(fieldValue.Uint(), 10))
			}
		case reflect.Bool:
			if fieldValue.Bool() {
				values.Add(fieldName, "true")
			}
		case reflect.Float32, reflect.Float64:
			if fieldValue.Float() != 0.0 {
				values.Add(fieldName, strconv.FormatFloat(fieldValue.Float(), 'f', FloatPrecision, 64))
			}
		}
	}

	return values, nil
}
