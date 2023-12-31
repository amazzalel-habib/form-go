/*
MIT License

Copyright (c) 2023 AMAZZAL EL-HABIB

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package formgo

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// FloatPrecision is the precision used when formatting float values.
var FloatPrecision = -1

// EncodeFormValues encodes a struct into URL form values by extracting tagged fields.
//
// The function accepts a struct and extracts fields with the "form" tag. It then
// converts non-empty string, non-zero integer, true boolean, and non-zero float fields
// into URL form values. It also handles pointer values for these primitive types.
//
// The "omitempty" tag can be used to skip fields that have the "omitempty" tag and are empty.
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
		parts := strings.Split(field.Tag.Get("form"), ",")
		fieldName := strings.TrimSpace(parts[0])
		omitempty := false
		if len(parts) > 1 && strings.TrimSpace(parts[1]) == "omitempty" {
			omitempty = true
		}

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
				if omitempty {
					continue
				}

				// For non-nil pointers with omitempty not set, add an empty string
				values.Add(fieldName, "")
				continue
			}
		}

		switch fieldValue.Kind() {
		case reflect.String:
			if fieldValue.String() != "" {
				values.Add(fieldName, fieldValue.String())
			} else if !omitempty {
				values.Add(fieldName, "")
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fieldValue.Int() != 0 {
				values.Add(fieldName, strconv.FormatInt(fieldValue.Int(), 10))
			} else if !omitempty {
				values.Add(fieldName, "0")
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if fieldValue.Uint() != 0 {
				values.Add(fieldName, strconv.FormatUint(fieldValue.Uint(), 10))
			} else if !omitempty {
				values.Add(fieldName, "0")
			}
		case reflect.Bool:
			if fieldValue.Bool() {
				values.Add(fieldName, "true")
			} else if !omitempty {
				values.Add(fieldName, "false")
			}
		case reflect.Float32, reflect.Float64:
			if fieldValue.Float() != 0.0 {
				values.Add(fieldName, strconv.FormatFloat(fieldValue.Float(), 'f', FloatPrecision, 64))
			} else if !omitempty {
				values.Add(fieldName, "0")
			}
		}
	}

	return values, nil
}
