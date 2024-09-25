package main

import "fmt"

func main() {
    // Deklarasi variabel
    var name string = "Shinta"  // variabel tipe string

    // Deklarasi variabel tanpa menyebutkan tipe
    city := "Surabaya"          // Go akan secara otomatis mengenali tipe datanya

    // Menampilkan nilai variabel
    fmt.Println("Nama:", name)
    fmt.Println("Asal:", city)

    // Mengubah nilai variabel
    name = "Putri"
    fmt.Println("Nama setelah diubah:", name)
}
