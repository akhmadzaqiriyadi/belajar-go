# 🇮🇩 Latihan Go: FizzBuzz & Pengecek Palindrom

Repositori ini berisi solusi untuk dua tantangan pemrograman klasik yang diimplementasikan menggunakan bahasa Go. Proyek ini bertujuan untuk melatih logika dasar, perulangan, kondisional, dan manipulasi string.

---

## 챌린지Tantangan yang Diselesaikan

### 1. FizzBuzz
Program mencetak angka dari 1 hingga 100 dengan aturan berikut:
-   Untuk kelipatan 3, cetak **"Fizz"**.
-   Untuk kelipatan 5, cetak **"Buzz"**.
-   Untuk kelipatan 3 dan 5, cetak **"FizzBuzz"**.
-   Selain itu, cetak angkanya.

### 2. Pengecek Palindrom
Program memeriksa apakah sebuah kata merupakan **palindrom** (dibaca sama dari depan maupun belakang).
-   Contoh: `katak`, `level`, `malam`.

---

## 🚀 Cara Menjalankan

1.  Pastikan **Go** sudah terinstal di komputer Anda.
2.  Clone repositori ini atau salin kode dari `main.go`.
3.  Buka terminal, masuk ke direktori proyek.
4.  Jalankan program dengan perintah:
    ```sh
    go run main.go
    ```
5.  Anda akan melihat output gabungan dari kedua tantangan tersebut.

---

## 🛠️ Cara Kerja

### Logika FizzBuzz
Algoritma ini menggunakan perulangan `for` untuk menghitung dari 1 sampai 100. Di setiap angka, ia menggunakan blok `if-else if-else` untuk memeriksa kondisi kelipatan. Kunci utamanya adalah memeriksa kondisi yang paling spesifik (`kelipatan 3 && 5`) terlebih dahulu sebelum memeriksa kondisi yang lebih umum.

```go
func fizzBuzz(limit int) {
    for i := 1; i <= limit; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
```

### Logika Palindrom
Algoritma ini secara efisien membandingkan karakter dari kedua ujung string, bergerak ke tengah.
-   **Dua Penunjuk**: Satu penunjuk (`i`) dimulai dari awal, dan penunjuk lainnya (`len(input)-1-i`) dimulai dari akhir.
-   **Perulangan Hingga Setengah**: Perulangan hanya berjalan sampai titik tengah string untuk efisiensi.
-   **Keluar Lebih Awal**: Jika ada pasangan karakter yang tidak cocok, fungsi langsung mengembalikan `false`. Jika perulangan selesai, fungsi mengembalikan `true`.

```go
func isPalindrome(input string) bool {
    for i := 0; i < len(input)/2; i++ {
        if input[i] != input[len(input)-1-i] {
            return false
        }
    }
    return true
}
```