/*	Nama : Ahmad Nur Rizal
	NIM : 1301194467
*/
package main

import "fmt"

const N int = 2019

type RecType struct {
	f1 string
	f2 int
	f3 float64
}
type ArrType [N]RecType

// Mencari Nilai terbesar f3
func rmax(data ArrType) float64 {
	var i, idxCari, idxMax int
	i = 0
	idxCari = 1
	idxMax = 0
	for i < N-1 { // N-1 karena idxCari = i + 1 sehingga data terakhir masih diperhitungkan
		if data[idxCari].f3 > data[idxMax].f3 {
			idxMax = idxCari
		}
		idxCari++
		i++
	}
	return data[idxMax].f3
}
func main() {
	var data ArrType
	fmt.Println(rmax(data))
	fmt.Println(" ")
	fmt.Println("Nama : Ahmad Nur Rizal")
	fmt.Println("NIM : 1301194467")

}
