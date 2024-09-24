![image](https://github.com/user-attachments/assets/81754b2f-fc5e-497e-9823-4cf28e740382)
![image](https://github.com/user-attachments/assets/7daf19dd-5368-4187-a698-b7b380468827)
![image](https://github.com/user-attachments/assets/72654c70-3d1c-47a6-a95d-7aa4b8f1b4a0)
![image](https://github.com/user-attachments/assets/0ceb4248-069c-47cf-82a7-9a3a3a712f5e)


    package main

    import "fmt"

    func main() {
        // Integer types
        a := int8(-128)                // menggunakan := dan cast ke int8
        b := int16(-32768)             // int16
        c := int32(-2147483648)        // int32
        d := int64(-9223372036854775808) // int64
        e := 2147483647                // otomatis jadi int (sesuai arsitektur)

    // Unsigned integer types
    f := uint8(255)               // uint8 (alias byte)
    g := uint16(65535)            // uint16
    h := uint32(4294967295)       // uint32
    i := uint64(18446744073709551615) // uint64
    j := uint(4294967295)         // otomatis jadi uint

    // Floating-point types
    k := float32(3.14)            // float32
    l := 3.141592653589793        // otomatis jadi float64

    // Complex number types
    m := complex(2, 3)            // complex64 (otomatis dari real + imag float32)
    n := complex(2.5, 7.5)        // complex128 (otomatis dari real + imag float64)

    // Menampilkan hasil
    fmt.Println("Contoh penggunaan tipe data integer:")
    fmt.Printf("int8: %d\n", a)
    fmt.Printf("int16: %d\n", b)
    fmt.Printf("int32: %d\n", c)
    fmt.Printf("int64: %d\n", d)
    fmt.Printf("int (32/64-bit): %d\n", e)

    fmt.Println("\nContoh penggunaan tipe data unsigned integer:")
    fmt.Printf("uint8 (byte): %d\n", f)
    fmt.Printf("uint16: %d\n", g)
    fmt.Printf("uint32: %d\n", h)
    fmt.Printf("uint64: %d\n", i)
    fmt.Printf("uint (32/64-bit): %d\n", j)

    fmt.Println("\nContoh penggunaan tipe data floating-point:")
    fmt.Printf("float32: %f\n", k)
    fmt.Printf("float64: %f\n", l)

    fmt.Println("\nContoh penggunaan tipe data complex:")
    fmt.Printf("complex64: %v\n", m)
    fmt.Printf("complex128: %v\n", n)

    // Mengakses bagian real dan imaginer dari bilangan kompleks
    fmt.Printf("Real part of complex64: %f\n", real(m))
    fmt.Printf("Imaginary part of complex64: %f\n", imag(m))
    fmt.Printf("Real part of complex128: %f\n", real(n))
    fmt.Printf("Imaginary part of complex128: %f\n", imag(n))
    }
