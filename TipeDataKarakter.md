![image](https://github.com/user-attachments/assets/fbb13920-0d33-4489-b749-9810c71227f5)

    package main

    import "fmt"

    func main() {
    // Deklarasi karakter menggunakan rune (unicode code point)
    var char rune = 'A'  // rune untuk karakter 'A'
    var anotherChar rune = '😊' // rune untuk emoji '😊'
   
       // Menggunakan len() untuk mendapatkan panjang string
    length := len(rune)

    // Menampilkan karakter dan nilai numeriknya (kode Unicode)
    fmt.Printf("Panjang string: %d\n", length)
    fmt.Printf("Karakter pertama: %c, nilai Unicode: %U\n", char, char)
    fmt.Printf("Karakter kedua: %c, nilai Unicode: %U\n", anotherChar, anotherChar)
    }
