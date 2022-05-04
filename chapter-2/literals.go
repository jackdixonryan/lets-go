package main

import "fmt"

// hi, this is the top of the file.
// I can declare vars up here
var thisVar string = "is-package-level"

// but I can not use the shorthand declaration outside of a function.
// thatVar := "no you do not"

func main() {
	// LITERALS
	// literals()
	// OPERATORS
	// operators()

	// Constancy
	usingConst()

}

func literals() {
	var IntegerLiteral = 9999
	var BinaryIntegerLiteral = 0b10
	var OctalIntegerLiteral = 0o10
	var HexadecimalIntegerLiteral = 0x10

	var HardReadinIntegerLiteral = 1920334923534
	var EasyReadinIntegerLiteral = 1_920_334_923_534

	var FloatingPointLiteral = 6.66
	var FloatingPointLiteralWithExponent = 6.66e10

	var RuneLiteralUnicode = 'a'
	var RuneLiteralOctal = '\141'
	var RuneLiteralHexadecimal = '\x61'
	var RuneLiteralUnicodeNumber = '\U00000061'

	var RuneLiteralNewLine = '\n'
	var RuneLiteralTab = '\t'
	var RuneLiteralSingleQuote = '\''
	// var RuneLiteralDoubleQuote = '\"'
	var RuneLiteralBackslash = '\\'

	fmt.Println(IntegerLiteral)            // 9999
	fmt.Println(BinaryIntegerLiteral)      // 2
	fmt.Println(OctalIntegerLiteral)       // 8
	fmt.Println(HexadecimalIntegerLiteral) // 16

	// note that, despite the difference in how the variables are assigned, they will appear the same from println.
	fmt.Println(HardReadinIntegerLiteral)
	fmt.Println(EasyReadinIntegerLiteral)

	fmt.Println(FloatingPointLiteral)
	fmt.Println(FloatingPointLiteralWithExponent)

	fmt.Println(RuneLiteralUnicode)
	fmt.Println(RuneLiteralOctal)
	fmt.Println(RuneLiteralHexadecimal)
	fmt.Println(RuneLiteralUnicodeNumber)

	fmt.Println(RuneLiteralNewLine)
	fmt.Println(RuneLiteralTab)
	fmt.Println(RuneLiteralSingleQuote)
	// fmt.Println(RuneLiteralDoubleQuote)
	fmt.Println(RuneLiteralBackslash)

	// fmt.Println()

	fmt.Println("Here lie the string literals.")

	var InterpretedStringLiteral = "This is an interpreted string literal."
	var RawStringLiteral = `This one will also account for 
	new lines
	and understand "typically escaped characters"`

	fmt.Println(InterpretedStringLiteral)
	fmt.Println(RawStringLiteral)

	// some literals can be explicitly typed; for the most part, variables without an assignment default to a value of 0
	// this is apparently not the case with booleans.

	var unassigned bool
	fmt.Println(unassigned) // is declared false by default.

	// I can use types for just about everything, though Go does not enforce types until I, the developer, expressly tell it to do so.

	// these can be positive, or negative.
	var myAge int8
	var moneyInMySavingsAccount int16
	var ageOfTheBloodyEarth int32
	var numberOfGrainsOfSandOnBeach int64

	// these are always positive numbers.
	var ageOfElf uint8
	var dollarsSavedSoFar uint16
	var idealNetWorth int32
	var tooBigToEvenMatter uint64

	// some of these types have aliases. For example
	var byteAsInt uint8
	// could also be expressed as
	var byteAsByte byte

	var defaultNumberType int
	// all numbers, unless specified, are typed as int.
	// in almost every case, just use int.

	// because none of these were assigned a value, they all default to 0.
	println(myAge, moneyInMySavingsAccount, ageOfTheBloodyEarth, numberOfGrainsOfSandOnBeach, ageOfElf, dollarsSavedSoFar, idealNetWorth, tooBigToEvenMatter, byteAsInt, byteAsByte, defaultNumberType)
}

func operators() {
	// integers support all kinds of arithmetic operators.
	var x int = 10
	var y int = x * 2
	fmt.Println(x, y)

	// like in JS, you can run an operator and an assignment simultaneously.
	y *= 5
	fmt.Println(y)

	// go also has the same comparators as JS
	var isGreater bool = x > y
	fmt.Println(isGreater)

	// and some stranger ones, specific for bit manipulation.
	// God knows how they work.

	// floating point variables can be targeted with any operator except modulo
	// and like in JS, they are subject to binary imprecision and cannot be totally trusted.

	// Strings can be declared and reassigned
	var message string = "hello, world!"
	message = "some other thing" // this is allowed

	// but they cannot be modified?
	message += "Some additional text"

	// no, it appears they can...

	// type conversions in Go must always be explicit
	var newX int = 10
	var newY float64 = 30.2
	var z float64 = float64(newX) + newY
	var d int = x + int(z)

	fmt.Println(newX, newY, z, d)

	// unlike JS, you cannot treat any other type as a 0 or a non-value.
	// null != 0
	// "" != 0
	// '' != 0
	// 0.00 != 0

	// that's why you have to know the 0 values of variables.
	// For strings, the 0 value is ""

	// you can declare multiple vars at once like this
	var ten, twenty int = 10, 20
	// or with 0 values
	var thirty, forty int

	// or with multiple types...
	var fifty, message2 = 50, "hi!"

	// and... a crazy way called a declaration list
	var (
		mx     int
		my         = 20
		mz     int = 30
		md, me     = 40, "hello"
		mf, mg string
	)

	fmt.Println(ten, twenty, thirty, forty, fifty, message2, mx, my, mz, md, me, mg, mf)

	// there's also a shorthand.
	yeet := 10
	yeety, yeeter := "yeet", "yoot"

	fmt.Println(yeet, yeety, yeeter)

	// for the next bit, return to the top of the file.

	// most commonly, you will see := used within functions
	// and declaration lists with var used for rare situations where you need to declare package level variables.

	// some special cases:
	// when initializing a variable with its 0 value, use var.
	var isZero int

	// when initializing a variable where the default type would otherwise be wrong, use var and avoid type coercion.
	var isByte byte = 20
	// to explain further: without typing, this is declared as an int variable.
	// using this var scheme allows you to specify that this var is NOT an int
	// and is a little clearer than something like
	// isByte := byte(20)

	fmt.Println(isZero, isByte)

}

func usingConst() {
	// const looks like it will work the same, but does it?

	const x int64 = 10
	const (
		idKey   = "id"
		nameKey = "name"
	)

	const z = 20 * 10

	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	// there are reassignments to a constant variable. They will cause the compilation to fail
	// x = x + 1
	// y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
