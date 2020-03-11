package main

import "fmt"

type RecType struct {
	count int
	size  int
}
type ArrType [100]RecType

func isFound(tab ArrType, n int, v int) int {
	var found int
	var bawah, tengah, atas int
	found = -1
	bawah = 0
	atas = n
	for bawah <= atas && found == -1 {
		tengah = (bawah + atas) / 2
		if v > tab[tengah].count {
			bawah = tengah + 1
		} else if v < tab[tengah].count {
			atas = tengah - 1
		} else {
			found = tengah
		}
	}
	return found
}

func main() {
	var tab ArrType
	var nsize, n, v int

	tab[0].count = 1
	tab[1].count = 2
	tab[2].count = 3
	tab[3].count = 4
	tab[4].count = 5
	tab[5].count = 6
	tab[6].count = 7
	tab[7].count = 8
	nsize = 8
	n = nsize

	fmt.Scanln(&v)
	fmt.Println(isFound(tab, n, v))

}
