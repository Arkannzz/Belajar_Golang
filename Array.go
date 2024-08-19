// package main

// import "fmt"

// func main() {
// 	// Struktur array satu data:
// 	// isinya tidak bolrh lebih banyak dari jumlah yang di tentukan 
// 	// var nama_array [jumlah_array] tipe_data
// 	// isinya array

// 	var name [7] string
// 	name [0] = "Arkan" 
// 	name [1] = "Zayed"
// 	name [2] = "Bagas"
// 	name [3] = "Raden"
// 	name [4] = "Ayu"
// 	name [5] = "Adit"

// 	fmt.Println("Array standar: ",name)

// 	// Append di array -> hanya bisa dipakai jika jumlah array nya tidak di tentukan
// 	var angka = [] int {0,1,2,3,4}
// 	var angka_baru = append(angka, 100)
// 	fmt.Println("Array append: ",angka_baru)

// 	// Mengisi nilai awal array
// 	var makanan = [3]string{"roti","nasi","jagung"}
// 	fmt.Println("Jumlah data: ", len(makanan))
// 	fmt.Println("Makanan ke-3: ", makanan[2])

// 	// Cara vertikal
// 	var minuman = [3]string{
// 		"air,",
// 		"susu coklat,",
// 		"teh beras",
// 	}
// 	fmt.Println("Minuman: ", minuman)

// 	// Mengisi nilai awal array tanpa jumlah elemen
// 	// var anime = [...]string{"Detective Conan", "DDemon Slayer", "Shokugeki no Souma", "data ke-n"}
// 	// fmt.Println(anime)

// 	// Array campuran
// 	array1 := []interface{}{1,2,true,"Saya makan"}
// 	fmt.Println(array1[3])

// 	identitas := []interface{}{
// 		"M. Arkan Deraya Dariawan,", 18, 29, "Maret", 2006, "Bandung.", "mahasiswa", "Telkom University,", "Informatika", "S1 Informatika,", "103012400268.", 3.87, "olahraga."}

// 	fmt.Println("Halo, perkenalkan nama aku ", identitas[0], "aku berusia", identitas[1], "aku lahir pada tanggal", identitas[2], identitas[3], identitas[4], "di", identitas[5], "Saat ini aku sedang penempuh pendidikan sebagai", identitas[6], "di", identitas[7],"aku tergabung dalam fakultas", identitas[8], "dengan program studi", identitas[9],"aku memiliki NIM", identitas[10], "Pada semester pertama, aku belajar dengan keras sehingga mendapat IPK", identitas[11], "dan aku sangat bangga akan hal itu karena disisi lain, aku juga berhasil menyeimbangkan belajar dengan hobiku yaitu ", identitas[12], "Sekarang aku masih berusaha untuk mengejar mimpiku menjadi seorang programmer yang bekerja diperusahaan luar negeri, dengan usaha dan doa yang kuat aku yakin bisa meraih semua keinginanku.")

// 	// Array Multidimensi
// 	// Jumlah elemennya boleh tidak ditulis
// 	var number = [2][3]int{[3]int{4,5,6}, [3]int{9,1,10}}
// 	// Pola: [jumlah array pertama][jumlah array ke-2/di dalam]tipe data{isinya}
// 	fmt.Println("Versi biasa: ",number)
// 	var number2 = [2][3]int{
// 		[3]int{4,5,6},
// 		[3]int{9,1,10},
// 	}
// 	fmt.Println("Versi vertical: ",number2)

// 	// Loop Array
// 	var anime = [...]string{"Detective Conan", "Demon Slayer", "Shokugeki no Souma", "data ke-n"}
// 	for i:=0; i < len(anime); i++ {
// 		fmt.Println("Anime ke-", i, ":", anime[i])
// 	}
// 	fmt.Println("Loop For-Range")
// 	for i, anime := range anime {
// 		fmt.Println("Anime ke-", i, ":", anime)
// 	}
// 	fmt.Println("Loop For-Range tanpa index")
// 	for _, anime := range anime {
// 		fmt.Println("Anime:", anime)
// 	}

// 	// Cara kedua membuat array
// 	var nama = make([]interface{}, 2)
// 	nama[0] = 10
// 	nama[1] = "Amanda"

// 	fmt.Println(nama)

// 	var ganjil = make([]int, 0)
// 	var genap = make([]int, 0)

// 	var angka_biasa = [15] int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
// 	fmt.Println(angka_biasa)
	
// 	for i := 0; i < len(angka_biasa); i++ {
// 		if i % 2 == 1 {
// 			ganjil = append(ganjil, i)
				
// 		} else {
// 			genap = append(genap, i)
			
// 		}
// 	} 
// 	fmt.Println(ganjil)
// 	fmt.Println(genap)


// }
