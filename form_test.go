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

type TestDataWithOmitEmpty struct {
	StringField     string   `form:"str_field,omitempty"`
	IntField        int      `form:"int_field,omitempty"`
	Int8Field       int8     `form:"int8_field,omitempty"`
	Int16Field      int16    `form:"int16_field,omitempty"`
	Int32Field      int32    `form:"int32_field,omitempty"`
	Int64Field      int64    `form:"int64_field,omitempty"`
	UintField       uint     `form:"uint_field,omitempty"`
	Uint8Field      uint8    `form:"uint8_field,omitempty"`
	Uint16Field     uint16   `form:"uint16_field,omitempty"`
	Uint32Field     uint32   `form:"uint32_field,omitempty"`
	Uint64Field     uint64   `form:"uint64_field,omitempty"`
	BoolField       bool     `form:"bool_field,omitempty"`
	Float32Field    float32  `form:"float32_field,omitempty"`
	Float64Field    float64  `form:"float64_field,omitempty"`
	PtrStringField  *string  `form:"ptr_str_field,omitempty"`
	PtrIntField     *int     `form:"ptr_int_field,omitempty"`
	PtrInt8Field    *int8    `form:"ptr_int8_field,omitempty"`
	PtrInt16Field   *int16   `form:"ptr_int16_field,omitempty"`
	PtrInt32Field   *int32   `form:"ptr_int32_field,omitempty"`
	PtrInt64Field   *int64   `form:"ptr_int64_field,omitempty"`
	PtrUintField    *uint    `form:"ptr_uint_field,omitempty"`
	PtrUint8Field   *uint8   `form:"ptr_uint8_field,omitempty"`
	PtrUint16Field  *uint16  `form:"ptr_uint16_field,omitempty"`
	PtrUint32Field  *uint32  `form:"ptr_uint32_field,omitempty"`
	PtrUint64Field  *uint64  `form:"ptr_uint64_field,omitempty"`
	PtrBoolField    *bool    `form:"ptr_bool_field,omitempty"`
	PtrFloat32Field *float32 `form:"ptr_float32_field,omitempty"`
	PtrFloat64Field *float64 `form:"ptr_float64_field,omitempty"`
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

type SubData struct {
	SubStringField string `form:"sub_str_field"`
}

type TestDataWithSubStruct struct {
	StringField string  `form:"str_field"`
	IntField    int     `form:"int_field"`
	SubData     SubData `form:"sub_data"`
}

func TestEncodeFormValues(t *testing.T) {
	FloatPrecision = 2
	testCases := []struct {
		name     string
		input    interface{}
		expected url.Values
	}{
		{
			name: "ValidDataSkipNoFormTag",
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
				"int_field":     []string{"0"},
				"int8_field":    []string{"0"},
				"int16_field":   []string{"0"},
				"int32_field":   []string{"0"},
				"int64_field":   []string{"0"},
				"uint_field":    []string{"0"},
				"uint8_field":   []string{"0"},
				"uint16_field":  []string{"0"},
				"uint32_field":  []string{"0"},
				"uint64_field":  []string{"0"},
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
			expected: url.Values{
				"str_field":     []string{""},
				"int_field":     []string{"0"},
				"int8_field":    []string{"0"},
				"int16_field":   []string{"0"},
				"int32_field":   []string{"0"},
				"int64_field":   []string{"0"},
				"uint_field":    []string{"0"},
				"uint8_field":   []string{"0"},
				"uint16_field":  []string{"0"},
				"uint32_field":  []string{"0"},
				"uint64_field":  []string{"0"},
				"bool_field":    []string{"false"},
				"float32_field": []string{"0"},
				"float64_field": []string{"0"},
			},
		},
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
		{
			name: "ValidData",
			input: TestDataWithOmitEmpty{
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
			input: TestDataWithOmitEmpty{
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
			input: TestDataWithOmitEmpty{
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
		{
			name: "ValidDataWithPointers",
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
			name: "ValidDataWithPointersZeroInt",
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
				"int_field":     []string{"0"},
				"int8_field":    []string{"0"},
				"int16_field":   []string{"0"},
				"int32_field":   []string{"0"},
				"int64_field":   []string{"0"},
				"uint_field":    []string{"0"},
				"uint8_field":   []string{"0"},
				"uint16_field":  []string{"0"},
				"uint32_field":  []string{"0"},
				"uint64_field":  []string{"0"},
				"bool_field":    []string{"true"},
				"float32_field": []string{"1.23"},
				"float64_field": []string{"2.34"},
			},
		},
		{
			name: "ValidDataWithPointersNilPointers",
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
			expected: url.Values{
				"str_field":     []string{""},
				"int_field":     []string{""},
				"int8_field":    []string{""},
				"int16_field":   []string{""},
				"int32_field":   []string{""},
				"int64_field":   []string{""},
				"uint_field":    []string{""},
				"uint8_field":   []string{""},
				"uint16_field":  []string{""},
				"uint32_field":  []string{""},
				"uint64_field":  []string{""},
				"bool_field":    []string{""},
				"float32_field": []string{""},
				"float64_field": []string{""},
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
