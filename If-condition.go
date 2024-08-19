// package main

// import"fmt"

// func main() {
	// If dengan satu kondisi
	/*if kondisi {
		aksi
	}*/

	// other language:
	// if(kondisi) {
	// 	aksi
	// }

	// var number int = 100
	// Kondisi baru harus banget ditulis di samping
	// if number < 30 {
	// 	fmt.Println("Jawaban Anda salah")
	// } else if number > 150 { //else if harus di samping
	// 	fmt.Println("Jawaban Anda benar")
	// } else {
	// 	fmt.Println("Invalid")
	// }

	// Project nilai
	// var nilai int

	// fmt.Print("Silahkan masukkan nilai: ")
	// fmt.Scan(&nilai)

	// if nilai > 90 {
	// 	fmt.Println("A")

	// } else if nilai > 79 && nilai < 90 {
	// 	fmt.Println("B")

	// } else if nilai > 59 && nilai < 80 {
	// 	fmt.Println("C")

	// } else if nilai > 49 && nilai < 60 {
	// 	fmt.Println("D")

	// } else {
	// 	fmt.Println("E")
	// }
	
	// var nilai_total float32 = 86.5

	// if rata_rata:= nilai_total / 10; rata_rata > 7 {
	// 	fmt.Println("Selamat Anda lulus")
	// } else if rata_rata > 6 && rata_rata < 7 {
	// 	fmt.Println("Maaf nilai Anda biasa aja, tapi lulus")
	// } else {
	// 	fmt.Println("Maaf nilai Anda terlalu kecil, Anda tidak lulus")
	// }

	// Switch case
	// Hanya berfokus pada satu variabel
	// Kalau sudah nemu 1 yang betulcase lain tidak akan di cek
	// var point = 3

	// switch point { //ngecek variabel point
	// case 8: //kalau nilainya 8 gimana
	// 	fmt.Println("Horee")
	// case 7:
	// 	fmt.Println("Good job")
	// case 6,5,4: //kalau nilainya 6,5,atau 4
	// 	fmt.Println("Practice more")
	// case 3:
	// 	fmt.Println("Wih hebat")
	// 	// Memaksa switch case untuk benar-benar mengecek semuanya
	// 	fallthrough
	// default:
	// 	fmt.Println("Nilai tidak ada")
	// }
	
	// Nested If-Switch
// 	var point = 0
// 	if point > 7 {
// 		switch point {
// 		case 10:
// 			fmt.Println("Perfect")
// 		default:
// 			fmt.Println("Great Job")
// 		}
// 	} else { // di bawah 7
// 		if point == 6 {
// 			fmt.Println("Good Job")
// 		} else if point == 5 {
// 			fmt.Println("Belajar lagi")
// 		} else { // 4,3,2,1
// 			fmt.Println("You can do it")
// 			if point == 0 {
// 				fmt.Println("Waduh, dimarahin mama")
// 			}
// 		}
// 	}
// }