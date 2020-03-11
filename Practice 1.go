/*  NIM  : 1301194467
	Nama : Ahmad Nur Rizal
	Dengan perogram ini, kita bisa mengetahui tahun keberapa saja yang merupakan tahun kabisat.
	Jika tahun yang dimasukkan adalah tahun kabisat maka akan muncul kata true, jika bukan maka muncul kata false
*/

package main
import "fmt"

func main() {

	var tahun int
	fmt.Print("Tahun : ")
	fmt.Scan(&tahun)
	fmt.Println("Kabisat :",tahun%400==0 || tahun%4==0 && tahun%100!=0)
    fmt.Println("Nama : Ahmad Nur Rizal")
    fmt.Print("NIM  : 1301194467")
}