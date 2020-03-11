/*	Nama : Ahmad Nur Rizal
	NIM : 1301194467
*/
package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

const N = 10000

var points [N]point
var numpoint int

func jarak(p1, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func bacaTitik() {
	var titik point
	fmt.Scan(&titik.x, &titik.y)
	for titik.x != 0.0 || titik.y != 0.0 {
		points[numpoint] = titik
		numpoint = numpoint + 1
		fmt.Scan(&titik.x, &titik.y)
	}
}

func ambilTitikTerdekat(p1, p2 *point) {
	var indeksSatu, indeksDua int
	var jarakTerdekat float64
	indeksSatu, indeksDua = 0, 1
	jarakTerdekat = jarak(points[indeksSatu], points[indeksDua])
	*p1 = points[indeksSatu]
	*p2 = points[indeksDua]
	for indeksSatu < numpoint-1 {
		if jarak(points[indeksSatu], points[indeksDua]) < jarakTerdekat {
			jarakTerdekat = jarak(points[indeksSatu], points[indeksDua])
			*p1 = points[indeksSatu]
			*p2 = points[indeksDua]
		}
		indeksDua++
		if indeksDua == numpoint {
			indeksSatu = indeksSatu + 1
			indeksDua = indeksSatu + 1
		}
	}
}

func main() {
	var titikTerdekat1, titikTerdekat2 point
	bacaTitik()
	ambilTitikTerdekat(&titikTerdekat1, &titikTerdekat2)
	fmt.Println("Titik terdekat adalah (", titikTerdekat1.x, ",", titikTerdekat1.y,
		") dan (", titikTerdekat2.x, ",", titikTerdekat2.y,
		") dengan jarak ", jarak(titikTerdekat1, titikTerdekat2))
	fmt.Println("")
	fmt.Println("Nama : Ahmad Nur Rizal")
	fmt.Println("NIM : 1301194467")
}
