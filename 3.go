package main
import "fmt"
const NMAX int = 1024
type arrBaju[NMAX]string
func main(){
	var A arrBaju
	var warnaBaju, searcKey string
	var i, size int
	i = 0
	size = 1
	fmt.Scan(&warnaBaju)
	for warnaBaju != "."{
		A[i] = warnaBaju
		i++
		size++
		fmt.Scan(&warnaBaju)
	}
	fmt.Scan(&searcKey)
	cekBaju(A, size, searcKey)
}
func cekBaju(A arrBaju, size int, searchKey string){
	var i int
	var b bool
	b = false
	for i = 0; i < size; i++{
		if A[i] == searchKey{
			b = true
		}
	}
	if b == true{
		fmt.Print("Baju Ada")
	}else{
		fmt.Print("Baju Tidak Ada")
	}
}