package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegipanjang PersegiPanjang) HitungLuas() int {
	return persegipanjang.Panjang * persegipanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001
}

func main() {
	persegi := Persegi{Sisi: 4}
	luas := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi: ", luas)

	persegiPanjangg := PersegiPanjang{Panjang: 4, Lebar: 5}
	luas = SeberapaLuas(persegiPanjangg)
	fmt.Println("Luas Persegi Panjang: ", luas)

	asal := Asal{Angka: 5}
	luas = SeberapaLuas(asal)
	fmt.Println("Luas Asal: ", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
