/*	Nama : Ahmad Nur Rizal
	NIM : 1301194467
*/
package main

import "fmt"

func main() {
	var suku_n int
	var ganjil, genap, selisih, pembagi float64
	pembagi = 1
	suku_n = 2
	selisih = 4 - (float64(-4)/float64(3))
	for selisih > 0.00001 || selisih < -0.00001 {
		if suku_n%2 != 0 {
			ganjil = ganjil - 4/pembagi
			genap = ganjil + 4/float64(pembagi+2)
		} else if suku_n%2 == 0 {
			ganjil = ganjil + 4/pembagi
			genap = ganjil - 4/float64(pembagi+2)
		}
		selisih = ganjil - genap
		suku_n++
		pembagi = pembagi + 2
	}
	fmt.Println("Hasil PI :", ganjil)
	fmt.Println("Hasil PI :", genap)
	fmt.Println("Pada i ke :", suku_n)
	fmt.Println("Nama : Ahmad Nur Rizal")
	fmt.Println("NIM : 1301194467")
}
