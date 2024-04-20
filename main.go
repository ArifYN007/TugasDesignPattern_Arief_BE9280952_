// PENERAPAN CREATIONAL DESIGN PATTTERNS (SINGLETON)
// https://refactoring.guru/design-patterns/singleton
package main

import "fmt"

//strukture mewakili suatu negara
type Goverment struct {
	name string
}

//instance adalah instansi tunggal dari pemerintah (singleton), karena dari awal sudah ada pemerintahan maka tidak boleh ada pembentukan pemerintahan baru
var instance = &Goverment{name: "The Goverment of X"}

//mengembalikan instansi tunggal dari pemerintah DAN HARUS CUMAN 1 INSTANSI
func GetInstance() *Goverment {
	return instance
}

func main() {
	//mendapatkan instansi tunggal dari pemerintah dan mencetak nama pemerintah YANG SEKARANG
	goverment := GetInstance()
	fmt.Println("Current goverment: ", goverment.name)
	// mencoba membuat instasi baru, yang harusnya mengembalikan instansi yang sama
	anotherGoverment := GetInstance()
	fmt.Println("Is it the same goverment instance,?", goverment == anotherGoverment)
}
