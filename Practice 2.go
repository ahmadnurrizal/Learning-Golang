/*Nama : Ahmad Nur Rizal
  NIM : 1301194467
  Program ini dapat menentukan kapan sepeda motor pak Andi akan oleng
  berdasarkan selisih massa isi 2 kantong terpal. Jika selisihnya lebih besar atau sama dengan 9 Kg, 
  maka motor pak Andi akan oleng, begitupun sebaliknya.
*/
package main

import "fmt"

func main() {
	var x1 float32
	var x2 float32
	var selisih bool
	var jumlah float32

	fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
	fmt.Scanln(&x1, &x2)
	jumlah = x1 + x2

	for jumlah <= 150 && x1 >= 0 && x2 >= 0 {
		selisih = x2-x1 >= 9 || x1-x2 >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng :", selisih)
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&x1, &x2)
		jumlah = x1 + x2
	}
	fmt.Println("Proses Selesai.")
	fmt.Println("")
	fmt.Println("Nama : Ahmad Nur Rizal")
	fmt.Println("NIM : 1301194467")
}
