package formgo

import (
	"net/url"
	"testing"
)

func TestEncodeFormValues_NonStructInput(t *testing.T) {
	nonStructData := "This is not a struct"
	_, err := EncodeFormValues(nonStructData)

	if err == nil {
		t.Error("Expected an error for non-struct input, but got nil.")
	} else if err.Error() != "data must be a struct" {
		t.Errorf("Expected error message 'data must be a struct', but got: %v", err.Error())
	}
}

type TestDataNoTag struct {
	StringField  string
	IntField     int
	Int8Field    int8
	Int16Field   int16
	Int32Field   int32
	Int64Field   int64
	UintField    uint
	Uint8Field   uint8
	Uint16Field  uint16
	Uint32Field  uint32
	Uint64Field  uint64
	BoolField    bool
	Float32Field float32
	Float64Field float64
}

func TestEncodeFormValues_SkipNoFormTag(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected url.Values
	}{
		{
			name: "ValidData",
			input: TestDataNoTag{
				StringField:  "abc",
				IntField:     123,
				Int8Field:    8,
				Int16Field:   16,
				Int32Field:   32,
				Int64Field:   64,
				UintField:    456,
				Uint8Field:   8,
				Uint16Field:  16,
				Uint32Field:  32,
				Uint64Field:  64,
				BoolField:    true,
				Float32Field: 1.23,
				Float64Field: 2.34,
			},
			expected: url.Values{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := EncodeFormValues(tc.input)
			if err != nil {
				t.Errorf("EncodeFormValues returned an error: %v", err)
			}

			if !urlValuesEqual(values, tc.expected) {
				t.Errorf("EncodeFormValues result mismatch. Got %v, expected %v", values, tc.expected)
			}
		})
	}
}

type TestData struct {
	StringField  string  `form:"str_field"`
	IntField     int     `form:"int_field"`
	Int8Field    int8    `form:"int8_field"`
	Int16Field   int16   `form:"int16_field"`
	Int32Field   int32   `form:"int32_field"`
	Int64Field   int64   `form:"int64_field"`
	UintField    uint    `form:"uint_field"`
	Uint8Field   uint8   `form:"uint8_field"`
	Uint16Field  uint16  `form:"uint16_field"`
	Uint32Field  uint32  `form:"uint32_field"`
	Uint64Field  uint64  `form:"uint64_field"`
	BoolField    bool    `form:"bool_field"`
	Float32Field float32 `form:"float32_field"`
	Float64Field float64 `form:"float64_field"`
}

func TestEncodeFormValues(t *testing.T) {
	FloatPrecision = 2
	testCases := []struct {
		name     string
		input    interface{}
		expected url.Values
	}{
		{
			name: "ValidData",
			input: TestData{
				StringField:  "abc",
				IntField:     123,
				Int8Field:    8,
				Int16Field:   16,
				Int32Field:   32,
				Int64Field:   64,
				UintField:    456,
				Uint8Field:   8,
				Uint16Field:  16,
				Uint32Field:  32,
				Uint64Field:  64,
				BoolField:    true,
				Float32Field: 1.23,
				Float64Field: 2.34,
			},
			expected: url.Values{
				"str_field":     []string{"abc"},
				"int_field":     []string{"123"},
				"int8_field":    []string{"8"},
				"int16_field":   []string{"16"},
				"int32_field":   []string{"32"},
				"int64_field":   []string{"64"},
				"uint_field":    []string{"456"},
				"uint8_field":   []string{"8"},
				"uint16_field":  []string{"16"},
				"uint32_field":  []string{"32"},
				"uint64_field":  []string{"64"},
				"bool_field":    []string{"true"},
				"float32_field": []string{"1.23"},
				"float64_field": []string{"2.34"},
			},
		},
		{
			name: "ZeroInt",
			input: TestData{
				StringField:  "abc",
				IntField:     0,
				Int8Field:    0,
				Int16Field:   0,
				Int32Field:   0,
				Int64Field:   0,
				UintField:    0,
				Uint8Field:   0,
				Uint16Field:  0,
				Uint32Field:  0,
				Uint64Field:  0,
				BoolField:    true,
				Float32Field: 1.23,
				Float64Field: 2.34,
			},
			expected: url.Values{
				"str_field":     []string{"abc"},
				"bool_field":    []string{"true"},
				"float32_field": []string{"1.23"},
				"float64_field": []string{"2.34"},
			},
		},
		{
			name: "ZeroAll",
			input: TestData{
				StringField:  "",
				IntField:     0,
				Int8Field:    0,
				Int16Field:   0,
				Int32Field:   0,
				Int64Field:   0,
				UintField:    0,
				Uint8Field:   0,
				Uint16Field:  0,
				Uint32Field:  0,
				Uint64Field:  0,
				BoolField:    false,
				Float32Field: 0,
				Float64Field: 0,
			},
			expected: url.Values{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := EncodeFormValues(tc.input)
			if err != nil {
				t.Errorf("EncodeFormValues returned an error: %v", err)
			}

			if !urlValuesEqual(values, tc.expected) {
				t.Errorf("EncodeFormValues result mismatch. Got %v, expected %v", values, tc.expected)
			}
		})
	}
}

type TestDataPtr struct {
	StringField  *string  `form:"str_field"`
	IntField     *int     `form:"int_field"`
	Int8Field    *int8    `form:"int8_field"`
	Int16Field   *int16   `form:"int16_field"`
	Int32Field   *int32   `form:"int32_field"`
	Int64Field   *int64   `form:"int64_field"`
	UintField    *uint    `form:"uint_field"`
	Uint8Field   *uint8   `form:"uint8_field"`
	Uint16Field  *uint16  `form:"uint16_field"`
	Uint32Field  *uint32  `form:"uint32_field"`
	Uint64Field  *uint64  `form:"uint64_field"`
	BoolField    *bool    `form:"bool_field"`
	Float32Field *float32 `form:"float32_field"`
	Float64Field *float64 `form:"float64_field"`
}

func TestEncodeFormValuesPointers(t *testing.T) {
	FloatPrecision = 2
	testCases := []struct {
		name     string
		input    interface{}
		expected url.Values
	}{
		{
			name: "ValidData",
			input: TestDataPtr{
				StringField:  String("abc"),
				IntField:     Int(123),
				Int8Field:    Int8(8),
				Int16Field:   Int16(16),
				Int32Field:   Int32(32),
				Int64Field:   Int64(64),
				UintField:    Uint(456),
				Uint8Field:   Uint8(8),
				Uint16Field:  Uint16(16),
				Uint32Field:  Uint32(32),
				Uint64Field:  Uint64(64),
				BoolField:    Bool(true),
				Float32Field: Float32(1.23),
				Float64Field: Float64(2.34),
			},
			expected: url.Values{
				"str_field":     []string{"abc"},
				"int_field":     []string{"123"},
				"int8_field":    []string{"8"},
				"int16_field":   []string{"16"},
				"int32_field":   []string{"32"},
				"int64_field":   []string{"64"},
				"uint_field":    []string{"456"},
				"uint8_field":   []string{"8"},
				"uint16_field":  []string{"16"},
				"uint32_field":  []string{"32"},
				"uint64_field":  []string{"64"},
				"bool_field":    []string{"true"},
				"float32_field": []string{"1.23"},
				"float64_field": []string{"2.34"},
			},
		},
		{
			name: "ZeroInt",
			input: TestDataPtr{
				StringField:  String("abc"),
				IntField:     Int(0),
				Int8Field:    Int8(0),
				Int16Field:   Int16(0),
				Int32Field:   Int32(0),
				Int64Field:   Int64(0),
				UintField:    Uint(0),
				Uint8Field:   Uint8(0),
				Uint16Field:  Uint16(0),
				Uint32Field:  Uint32(0),
				Uint64Field:  Uint64(0),
				BoolField:    Bool(true),
				Float32Field: Float32(1.23),
				Float64Field: Float64(2.34),
			},
			expected: url.Values{
				"str_field":     []string{"abc"},
				"bool_field":    []string{"true"},
				"float32_field": []string{"1.23"},
				"float64_field": []string{"2.34"},
			},
		},
		{
			name: "NilPointers",
			input: TestDataPtr{
				StringField:  nil,
				IntField:     nil,
				Int8Field:    nil,
				Int16Field:   nil,
				Int32Field:   nil,
				Int64Field:   nil,
				UintField:    nil,
				Uint8Field:   nil,
				Uint16Field:  nil,
				Uint32Field:  nil,
				Uint64Field:  nil,
				BoolField:    nil,
				Float32Field: nil,
				Float64Field: nil,
			},
			expected: url.Values{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := EncodeFormValues(tc.input)
			if err != nil {
				t.Errorf("EncodeFormValues returned an error: %v", err)
			}

			if !urlValuesEqual(values, tc.expected) {
				t.Errorf("EncodeFormValues result mismatch. Got %v, expected %v", values, tc.expected)
			}
		})
	}
}

type SubData struct {
	SubStringField string `form:"sub_str_field"`
}

type TestDataWithSubStruct struct {
	StringField string  `form:"str_field"`
	IntField    int     `form:"int_field"`
	SubData     SubData `form:"sub_data"`
}

func TestEncodeFormValuesWithSturct(t *testing.T) {
	FloatPrecision = 2
	testCases := []struct {
		name     string
		input    interface{}
		expected url.Values
	}{
		{
			name: "ValidData",
			input: TestDataWithSubStruct{
				StringField: "abc",
				IntField:    123,
				SubData: SubData{
					SubStringField: "sub-abc",
				},
			},
			expected: url.Values{
				"str_field": []string{"abc"},
				"int_field": []string{"123"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := EncodeFormValues(tc.input)
			if err != nil {
				t.Errorf("EncodeFormValues returned an error: %v", err)
			}

			if !urlValuesEqual(values, tc.expected) {
				t.Errorf("EncodeFormValues result mismatch. Got %v, expected %v", values, tc.expected)
			}
		})
	}
}

func urlValuesEqual(a, b url.Values) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valuesA := range a {
		valuesB, ok := b[key]
		if !ok {
			return false
		}
		if len(valuesA) != len(valuesB) {
			return false
		}
		for i, valueA := range valuesA {
			if valueA != valuesB[i] {
				return false
			}
		}
	}
	return true
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Int8(i int8) *int8 {
	return &i
}

func Int16(i int16) *int16 {
	return &i
}

func Int32(i int32) *int32 {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func Uint(u uint) *uint {
	return &u
}

func Uint8(u uint8) *uint8 {
	return &u
}

func Uint16(u uint16) *uint16 {
	return &u
}

func Uint32(u uint32) *uint32 {
	return &u
}

func Uint64(u uint64) *uint64 {
	return &u
}

func Bool(b bool) *bool {
	return &b
}

func Float32(f float32) *float32 {
	return &f
}

func Float64(f float64) *float64 {
	return &f
}
