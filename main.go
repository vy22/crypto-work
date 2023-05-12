package main

import (
	"fmt"
	"strconv"
)

// Створення типу даних BigNumber
type BigNumber struct {
	hex string
}

// Сеттер для встановлення значення в hex
func (b *BigNumber) setHex(hex string) {
	b.hex = hex
}

// Геттер для отримання значення в hex
func (b *BigNumber) getHex() string {
	return b.hex
}

// Функція для виконання операції побітового "XOR" для двох чисел BigNumber
func (b *BigNumber) bitXor(a BigNumber) BigNumber {
	// Отримаємо два числа у Hex
	hex1 := b.getHex()
	hex2 := a.getHex()
	// Створюємо результуюче число у Hex
	resultHex := ""

	// Проходимо по числам і використовуємо побітову операцію "OR"
	for i := 0; i < len(hex1); i++ {
		int1, _ := strconv.ParseInt(string(hex1[i]), 16, 0)
		int2, _ := strconv.ParseInt(string(hex2[i]), 16, 0)
		resultHex += fmt.Sprintf("%X", int1^int2)
	}
	// Створюємо результуючий BigNumber та встановлюємо йому результат
	var result2 BigNumber
	result2.setHex(resultHex)
	return result2
}

// повертає результат побітового АБО двох BigNumber
func XOR(bn1 *BigNumber, bn2 *BigNumber) *BigNumber {
	var result2 BigNumber
	h1 := bn1.getHex()
	h2 := bn2.getHex()
	// convert the hex strings to ints
	num1, _ := strconv.ParseInt(h1, 16, 64)
	num2, _ := strconv.ParseInt(h2, 16, 64)

	// perform bitwise OR
	res := num1 ^ num2

	// convert the result2 to a hex string
	hres := fmt.Sprintf("%x", res)

	// set the hex value of the result
	result2.setHex(hres)
	return &result2
}

// Функція для виконання операції побітового "І" для двох чисел BigNumber
func (b *BigNumber) bitAnd(a BigNumber) BigNumber {
	// Отримаємо два числа у Hex
	hex1 := b.getHex()
	hex2 := a.getHex()
	// Створюємо результуюче число у Hex
	resultHex := ""

	// Проходимо по числам і використовуємо побітову операцію "І"
	for i := 0; i < len(hex1); i++ {
		int1, _ := strconv.ParseInt(string(hex1[i]), 16, 0)
		int2, _ := strconv.ParseInt(string(hex2[i]), 16, 0)
		resultHex += fmt.Sprintf("%X", int1&int2)
	}
	// Створюємо результуючий BigNumber та встановлюємо йому результат
	var result BigNumber
	result.setHex(resultHex)
	return result
}

// повертає результат побітового I двох BigNumber
func AND(bn1 *BigNumber, bn2 *BigNumber) *BigNumber {
	var result BigNumber
	h1 := bn1.getHex()
	h2 := bn2.getHex()

	// convert the hex strings to ints
	num1, _ := strconv.ParseInt(h1, 16, 64)
	num2, _ := strconv.ParseInt(h2, 16, 64)

	// perform bitwise I
	res := num1 & num2

	// convert the result to a hex string
	hres := fmt.Sprintf("%x", res)

	// set the hex value of the result
	result.setHex(hres)
	return &result
}

// Функція для виконання операції побітового "OR" для двох чисел BigNumber
func (b *BigNumber) bitOr(a BigNumber) BigNumber {
	// Отримаємо два числа у Hex
	hex1 := b.getHex()
	hex2 := a.getHex()
	// Створюємо результуюче число у Hex
	resultHex := ""

	// Проходимо по числам і використовуємо побітову операцію "OR"
	for i := 0; i < len(hex1); i++ {
		int1, _ := strconv.ParseInt(string(hex1[i]), 16, 0)
		int2, _ := strconv.ParseInt(string(hex2[i]), 16, 0)
		resultHex += fmt.Sprintf("%X", int1|int2)
	}
	// Створюємо результуючий BigNumber та встановлюємо йому результат
	var result1 BigNumber
	result1.setHex(resultHex)
	return result1
}

// повертає результат побітового АБО двох BigNumber
func OR(bn1 *BigNumber, bn2 *BigNumber) *BigNumber {
	var result1 BigNumber
	h1 := bn1.getHex()
	h2 := bn2.getHex()
	// convert the hex strings to ints
	num1, _ := strconv.ParseInt(h1, 16, 64)
	num2, _ := strconv.ParseInt(h2, 16, 64)

	// perform bitwise OR
	res := num1 | num2

	// convert the result1 to a hex string
	hres := fmt.Sprintf("%x", res)

	// set the hex value of the result
	result1.setHex(hres)
	return &result1
}

func main() {
	// Створюємо два числа BigNumber
	var b1 BigNumber
	b1.setHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	var b2 BigNumber
	b2.setHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")

	// Виконуємо для них побітову операцію "XOR"
	var result2 = b1.bitXor(b2)

	// Виводимо результат
	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)

	fmt.Println("XOR: ", result2.getHex()) // Output: FFCD

	// Виконуємо для них побітову операцію "I"
	var result = b1.bitAnd(b2)
	fmt.Println("\nb1: ", b1)
	fmt.Println("b2: ", b2)
	// Виводимо результат
	fmt.Println("I:   ", result.getHex()) // Output: FFCD

	fmt.Println("\nb1: ", b1)
	fmt.Println("b2: ", b2)
	// Виконуємо для них побітову операцію "OR"
	var result1 = b1.bitOr(b2)

	// Виводимо результат
	fmt.Println("OR:  ", result1.getHex()) // Output: FFCD

}
