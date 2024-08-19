// package main  

// import (
// 	"fmt"
// )

// func main() {
	
	
	// For -> Perulangan yang ada jatahnya
	// Kemampuannya gabungan dari for, foreach, dan while
	// for i:= 0 (ngitung dari berapa); i < 5 (ngitung sampai berapa); i++;
	// i++ = ngitung normal

	// 1. For Standar
	// Pola: start counting; counting end; skip counting atau tidak
	// for i:= 0; i<=5; i++ {
	// 	fmt.Println("Angka ke-", i)
	// }

	// for i:= 1; i<=100; i+=3 {
	// 	fmt.Println("Angka ke-", i)
	// }

	// 2. For hanya dengan kondisi sebagai argumen
	// var i = 0 //untuk penghitung
	// for i < 5 {
	// 	fmt.Println("Saya sedang belajar bahasa Go.")
	// 	i++
	// }

	// 3. For tanpa kondisi
	// var i = 0
	// for {
	// 	fmt.Println(i,". Saya suka main sepuluh jam")
	// 	i++
	// 	if i == 10 {
	// 		break
	// 	}
	// }

	// 1. For
	// 2. Foreach
	// 3. While ... do ...
	// 4. Do ... while ...
	// 5. Repeat until

	// Additional
	// 1. Menggunakan string
	// var number = "10"
	// for i, v:= range number { //i = index; v = value
	// 	fmt.Println("Index ke- ", i, "Valuenya adalah: ", v)
	// }

	// var array1 = [5]int{10,20,30,40,50}
	// for _, value:= range array1 {
	// 	fmt.Println("Value: ", value)
	// }

	// Break & Continue
	// for start; end; skip counting/not {

	// }
	// for i:= 1; i <= 200; i+=3 {
	// 	if i % 2 == 1 {
	// 		continue
	// 	}
	// 	if i > 200 {
	// 		break
	// 	}
	// 	fmt.Println("Angka: ", i)
	// }

	// Nested Loop
	// Sampel
	// Tepuk tangan
	// for i:= 1; i < 4; i++ {
	// 	for i:= 1; i < 4; i++ {
	// 		fmt.Println("Tepuk tangan")
	// 	}
	// 	for i:= 1; i < 3; i++ {
	// 		fmt.Println("Tepuk pipi")
	// 	}
	// }

	// Pola segitiga bintang
	// fmt.Println("Segitiga terbalik")
	// for i:=0; i<10; i++ {
	// 	for j:=i; j<10; j++ {
	// 		fmt.Print("*", " ")
	// 	}

	// 	fmt.Println() //enter baris baru
	// }
	
	// fmt.Println("Segitiga normal")
	// for a:=10; a>0; a-=1 {
	// 	for b:=a; b<10; b++ {
	// 		fmt.Print("*", " ")
	// 	}

	// 	fmt.Println() //enter baris baru
	// }
	// fmt.Println("Segitiga Sama Kaki")
	// sisi := 10
	// jarak := (sisi/2)
	// jarak2 := strings.Repeat(" ", jarak)
	// for c:=sisi; c>0; c+=1 {
	// 	for d:=c; d<10; d++ {
	// 		if c % 2 == 1 {
	// 			count := strings.Repeat("*",c)
	// 			fmt.Print(jarak2, count)
	// 			jarak = jarak-1
				
	// 		} else {
	// 			continue
	// 		} if c>sisi {
	// 			break
	// 		}
				
	// 	}

	//  fmt.Println() //enter baris baru
	// }

	
// 	var number int
	
// 	fmt.Print("Masukkan jumlah baris: ")
// 	fmt.Scan(&number)
// 	fmt.Println("Output: ")
	
// 	for i := 1; i <= number; i++ {
// 		for j := 1; j <= number-i; j++ {
// 			fmt.Print(" ")
// 		}
// 		for k := 1; k <= i*2-1; k++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	} 

// }
	
