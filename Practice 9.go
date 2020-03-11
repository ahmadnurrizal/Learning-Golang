/*	Nama : Ahmad Nur Rizal
	NIM : 1301194467
*/
package main

import "fmt"

const MAXSIZE int = 20192020

type RecType struct {
	count int
	size  int
}
type ArrType [MAXSIZE]RecType

func iSort(tab *ArrType, nsize int) { 
	var i, pass, temp int
	pass = 0
	for pass < nsize-1 && pass < MAXSIZE-1 {
		i = pass + 1
		temp = tab[i].count
		for i > 0 && temp < tab[i-1].count {
			tab[i].count = tab[i-1].count
			i = i - 1
		}
		tab[i].count = temp
		pass = pass + 1
	}
}

func mSort(tab *ArrType, nsize int) { 
	var i, pass, temp int
	pass = 0
	for pass < nsize-1 && pass < MAXSIZE-1 {
		i = pass + 1
		temp = tab[i].size
		for i > 0 && temp > tab[i-1].size {
			tab[i].size = tab[i-1].size
			i = i - 1
		}
		tab[i].size = temp
		pass = pass + 1
	}
}

func isFound(tab ArrType, n int, v int) bool { 
	var found int
	var min, mid, max int
	max = n
	found = -1
	min = 0

	for min <= max && found == -1 {
		mid = (min + max) / 2
		if v > tab[mid].count {
			min = mid + 1
		} else if v < tab[mid].count {
			max = mid - 1
		} else {
			found = mid
		}
	}
	if found != -1 {
		return true
	} else {
		return false
	}
}

func main() {
	var tab ArrType
	var i, nsize, n, v int

	iSort(&tab, nsize)
	fmt.Scanln(&v)
	fmt.Print(isFound(tab, n, v))
	fmt.Println("")
	fmt.Println("Nama : Ahmad Nur Rizal")
	fmt.Println("NIM : 1301194467")
}
