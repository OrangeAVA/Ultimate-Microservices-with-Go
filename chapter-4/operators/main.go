package main

func main() {
	// Arithmetic Operators
	x := 5 + 5 // 10
	x = 5 - 5  // 0
	x = 5 * 5  // 25
	x = 5 / 5  // 1
	x = 5 % 5  // 0
	x++        // 6
	x--        // 5
	println(x)

	// Comparison Operators
	y := 5 == 5 // true
	y = 5 != 5  // false
	y = 5 > 5   // false
	y = 5 < 5   // false
	y = 5 >= 5  // true
	y = 5 <= 5  // true
	println(y)

	// Logical Operators
	a := true && true // true
	a = true && false // false
	a = true || true  // true
	a = true || false // true
	a = !true         // false
	a = !false        // true
	println(a)

	// Bitwise Operators
	b := 5 & 5 // 5
	b = 5 | 5  // 5
	b = 5 ^ 5  // 0
	b = 5 << 5 // 160
	b = 5 >> 5 // 0
	println(b)

	// Assignment Operators
	c := 5
	c += 5  // 10
	c -= 5  // 5
	c *= 5  // 25
	c /= 5  // 5
	c %= 5  // 0
	c &= 5  // 0
	c |= 5  // 5
	c ^= 5  // 5
	c <<= 5 // 160
	c >>= 5 // 5
	println(c)
}
