package main

import "fmt"

// --- Challenge 1: FizzBuzz ---

// fizzBuzz adalah fungsi yang menjalankan logika FizzBuzz hingga batas tertentu.
// - Menerima satu argumen: `limit` bertipe `int` (misal: 100).
// - Tidak mengembalikan nilai, karena tugasnya hanya mencetak ke layar.
func fizzBuzz(limit int) {
	fmt.Println("--- Tantangan 1: FizzBuzz ---")
	// Lakukan perulangan dari 1 sampai `limit`.
	for i := 1; i <= limit; i++ {
		// Cek kondisi yang paling spesifik terlebih dahulu.
		// Apakah `i` kelipatan 3 DAN 5?
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 { // Jika tidak, apakah kelipatan 3?
			fmt.Println("Fizz")
		} else if i%5 == 0 { // Jika tidak, apakah kelipatan 5?
			fmt.Println("Buzz")
		} else { // Jika tidak memenuhi semua kondisi di atas.
			fmt.Println(i)
		}
	}
}

// --- Challenge 2: Palindrome Checker ---

// isPalindrome adalah fungsi untuk memeriksa sebuah string.
// - Menerima satu argumen: `input` bertipe `string`.
// - Mengembalikan satu nilai: `bool` (true jika palindrom, false jika tidak).
func isPalindrome(input string) bool {
	// Lakukan perulangan hanya sampai setengah panjang string untuk efisiensi.
	for i := 0; i < len(input)/2; i++ {
		// Bandingkan karakter dari kiri (input[i]) dengan karakter
		// yang bersesuaian dari kanan (input[len(input)-1-i]).
		if input[i] != input[len(input)-1-i] {
			// Jika ada satu saja yang tidak cocok, langsung kembalikan false.
			return false
		}
	}
	// Jika loop selesai tanpa masalah, berarti ini adalah palindrom.
	return true
}

// --- Entry Point Program ---

// main adalah fungsi utama tempat program kita akan mulai berjalan.
func main() {
	// Menjalankan tantangan FizzBuzz hingga 100.
	fizzBuzz(100)

	// Membuat pemisah agar output rapi.
	fmt.Println("\n--- Tantangan 2: Pengecek Palindrom ---")

	// Menyiapkan beberapa kata untuk diuji.
	kata1 := "katak"
	kata2 := "malam"
	kata3 := "go"

	// Memanggil fungsi isPalindrome untuk setiap kata dan mencetak hasilnya.
	fmt.Printf("%s adalah palindrom? %t\n", kata1, isPalindrome(kata1))
	fmt.Printf("%s adalah palindrom? %t\n", kata2, isPalindrome(kata2))
	fmt.Printf("%s adalah palindrom? %t\n", kata3, isPalindrome(kata3))
}