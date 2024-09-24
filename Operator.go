package main

import (
	"fmt"
)

func main() {
	// Operator Aritmatika
	var a, b float64
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	fmt.Printf("Hasil Penjumlahan: %.2f\n", a+b)
	fmt.Printf("Hasil Pengurangan: %.2f\n", a-b)
	fmt.Printf("Hasil Perkalian: %.2f\n", a*b)
	fmt.Printf("Hasil Pembagian: %.2f\n", a/b)
	fmt.Printf("Hasil Modulus: %.2f\n", int(a)%int(b)) // Modulus hanya untuk integer

	// Operator Bitwise
	var x, y int = 10, 4 // 10 (1010) dan 4 (0100)
	fmt.Printf("Hasil Bitwise AND: %d\n", x&y) // AND
	fmt.Printf("Hasil Bitwise OR: %d\n", x|y)  // OR
	fmt.Printf("Hasil Bitwise XOR: %d\n", x^y) // XOR
	fmt.Printf("Hasil Bitwise NOT: %d\n", ^x)   // NOT
	fmt.Printf("Hasil Left Shift: %d\n", x<<1)  // Left Shift
	fmt.Printf("Hasil Right Shift: %d\n", x>>1) // Right Shift

	// Operator Perbandingan
	fmt.Printf("%f < %f: %t\n", a, b, a < b)
	fmt.Printf("%f <= %f: %t\n", a, b, a <= b)
	fmt.Printf("%f > %f: %t\n", a, b, a > b)
	fmt.Printf("%f >= %f: %t\n", a, b, a >= b)
	fmt.Printf("%f == %f: %t\n", a, b, a == b)
	fmt.Printf("%f != %f: %t\n", a, b, a != b)

	// Operator Logika
	var xBool, yBool bool
	fmt.Print("Masukkan nilai boolean pertama (true/false): ")
	fmt.Scan(&xBool)
	fmt.Print("Masukkan nilai boolean kedua (true/false): ")
	fmt.Scan(&yBool)

	fmt.Printf("x && y: %t\n", xBool && yBool) // Logical AND
	fmt.Printf("x || y: %t\n", xBool || yBool) // Logical OR
	fmt.Printf("!x: %t\n", !xBool)              // Logical NOT

	// Pointer dan Referensi
	num := 20
	ptr := &num // Pointer ke num
	fmt.Printf("Nilai num: %d\n", num)
	fmt.Printf("Alamat num: %p\n", &num)
	fmt.Printf("Nilai melalui pointer: %d\n", *ptr) // Dereference pointer
}
