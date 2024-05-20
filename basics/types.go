package basics

import (
	"fmt"
	"math"
)

type DataType struct {
	name string
	max  interface{}
	min  interface{}
}

func Types() {

	var uint8_default_value uint8
	var uint16_default_value uint16
	var uint32_default_value uint32

	// byte: Represents a single 8-bit unsigned integer (0 to 255).
	// It's primarily used for storing raw byte data.
	var byte_default_value byte

	// rune: Represents a Unicode code point, which can be a single byte or multiple bytes depending on the character. It's used for representing characters in text.
	var success rune = '🟢'
	var warning rune = '🟡'
	var danger rune = '🔴'

	type_bool := DataType{name: "bool", max: true, min: false}

	type_byte := DataType{name: "byte", max: 255, min: byte_default_value}

	type_uint8 := DataType{name: "uint8", max: math.MaxUint8, min: uint8_default_value}
	type_uint16 := DataType{name: "uint16", max: math.MaxUint16, min: uint16_default_value}
	type_uint32 := DataType{name: "uint32", max: math.MaxUint32, min: uint32_default_value}

	type_int := DataType{name: "int", max: math.MaxInt, min: math.MinInt}
	type_int8 := DataType{name: "int8", max: math.MaxInt8, min: math.MinInt8}
	type_int16 := DataType{name: "int16", max: math.MaxInt16, min: math.MinInt16}
	type_int32 := DataType{name: "int32", max: math.MaxInt32, min: math.MinInt32}
	type_int64 := DataType{name: "int64", max: math.MaxInt64, min: math.MinInt64}

	type_float32 := DataType{name: "float32", max: math.MaxFloat32, min: math.SmallestNonzeroFloat32}
	type_float64 := DataType{name: "float64", max: math.MaxFloat64, min: math.SmallestNonzeroFloat64}

	fmt.Println("Unicode:", string('\u2705'))
	fmt.Println("Unicode:", string('\u274C'))

	fmt.Println("Rune:", string(success))
	fmt.Println("Rune:", string(warning))
	fmt.Println("Rune:", string(danger))

	fmt.Printf("%+v \n", type_bool)

	fmt.Printf("%+v \n", type_byte)

	fmt.Printf("%+v \n", type_uint8)
	fmt.Printf("%+v \n", type_uint16)
	fmt.Printf("%+v \n", type_uint32)

	fmt.Printf("%+v \n", type_int)
	fmt.Printf("%+v \n", type_int8)
	fmt.Printf("%+v \n", type_int16)
	fmt.Printf("%+v \n", type_int32)
	fmt.Printf("%+v \n", type_int64)

	fmt.Printf("%+v \n", type_float32)
	fmt.Printf("%+v \n", type_float64)
}
