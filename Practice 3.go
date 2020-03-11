/*	Nama	: Ahmad Nur Rizal
	NIM		: 1301194467
*/
package main

import "fmt"

func main() {
	var nam float32
	var nmk string
	fmt.Print("Nilai akhir mata kuliah : ")
	fmt.Scan(&nam)
	if nam >= 0 && nam <= 100 {
		if nam > 80 && nam <= 100 {
			nmk = "A"
		} else if nam > 72.5 && nam <= 80 {
			nmk = "AB"
		} else if nam > 65 && nam <= 72.5 {
			nmk = "B"
		} else if nam > 57.5 && nam <= 65 {
			nmk = "C"
		} else if nam > 50 && nam <= 57.5 {
			nmk = "BC"
		} else if nam > 40 && nam <= 50 {
			nmk = "D"
		} else {
			nmk = "E"
		}
		fmt.Println("NIlai mata kuliah :", nmk)
	}
	fmt.Println("")
	fmt.Println("Nama : Ahmad Nur Rizal")
	fmt.Println("NIM : 1301194467")
}
