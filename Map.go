package main 

import "fmt"

func main() {

	// Map --> BUKAN PETA

	// Cara pertama:
	// var minuman map[tipe_data_key]tipe data value
	var makanan map[string]string
	makanan = map[string]string{}
	makanan["nama"] = "bubur ayam"
	fmt.Println(makanan)

	// Cara kedua:
	minuman := map[string]string{} // Percobaan
	minuman["nama"] = "susu"
	minuman["rasa"] = "vanilla"
	fmt.Println(minuman)

	// Inisialisasi nilai awal
	// Map bernilai awal nil
	var data2 = map[string]int{
		// Inisialisasi nilai awal disini
		// Cara vertikal
		"nama": 50,
		"jurusan":45,
		"datake-n":60,
	}
	// Cara horizontal
	var data3 = map[string]string{"nama":"Arkan", "Jurusan":"Informatika"}
	fmt.Println(data2)
	// data2["nama"] = 25 // Ini ngisi dataanya nanti, atau dapat digunakan untuk meng-update data
	fmt.Println(data3)
	var data_int = map[int]string{}
	data_int[01] = "halo"
	fmt.Println(data_int)

	// Data Campur
	var data_campur = map[string]interface{} {
		"nama":"Arkan",
		"jurusan":"Informatika",
		"Umur":18,
	}
	fmt.Println("Data campuran: ", data_campur)

	var data_campur_key = map[interface{}]interface{} {
		"nama":"Arkan",
		"jurusan":"Informatika",
		"Umur":18,
		2024: true,
	}
	fmt.Println("Data campuran: ", data_campur_key)

	// Nested Map
	var x = map[string]map[string]string{}
 
    x["buah"] = map[string]string{}
    x["warna"] = map[string]string{}
 
    x["buah"]["a"] = "apel"
    x["buah"]["b"] = "pisang"
 
    x["warna"]["r"] = "merah"
    x["warna"]["y"] = "kuning"
 
    fmt.Println(x)

}