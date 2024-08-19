// package main

// import "fmt"

// func main() {
// 	// Slice
// 	// Copy Slice
// 	// Untuk menyalin slice - agar memiliki referensi memori yang berbeda
// 	makanan_baru := make([]string, 3)
// 	makanan := []string{"tempe", "coklat", "susu vanilla"}
// 	copy(makanan_baru, makanan)
// 	fmt.Println("Makanan: ", makanan)
// 	fmt.Println("Makanan baru: ", makanan_baru)
// 	println(makanan_baru)
// 	println(makanan)
	
// 	// Append -> menambah data
// 	n := append(makanan_baru, "bayam", "ayam betutu")
// 	fmt.Println(makanan)
// 	fmt.Println(n)

// 	// Cap
// 	// Menghitung kapasitas maksimum dari slice setelah operasi dilakukan
// 	cap1 := cap(makanan_baru)
// 	fmt.Println(cap1)

// 	// 1. Mengambil data dalam range tertentu (dengan dua index)
// 	minuman := []string{"susu cokelat", "susu vanilla", "banana uyu", "kopi"}
// 	fmt.Println(minuman)
// 	fmt.Println(minuman[0:2]) //Mengambil dari index ke 0 sampai sebelum index ke-1
	
// 	// 2. Akses data dengan tiga index
// 	// Kapasitas tidak boleh melebihi jumlah elemen slice-nya
// 	fmt.Println("Tiga index: ", minuman[0:2:3]) //Nama slice[start:end:capasity]
// }

