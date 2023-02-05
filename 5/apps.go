package main

import "fmt"

type Kendaraan struct {
	Suara string
}

func (k *Kendaraan) Akselerasi() string {
	return k.Suara
}

type Sepeda struct {
	Kendaraan
	Rantai string
}

func (s *Sepeda) Akselerasi() string {
	s.Rantai = "Perlu perbaikan"
	return s.Kendaraan.Akselerasi()
}

type Mobil struct {
	Kendaraan
	Bensin string
}

func (m *Mobil) Akselerasi() string {
	m.Bensin = "Habis"
	return m.Kendaraan.Akselerasi()
}

func main() {
	sepeda := Sepeda{
		Kendaraan: Kendaraan{
			Suara: "Swoosh",
		},
		Rantai: "Normal",
	}
	mobil := Mobil{
		Kendaraan: Kendaraan{
			Suara: "Vroom",
		},
		Bensin: "Penuh",
	}

	fmt.Println("Sebelum sepeda berakselerasi:", sepeda.Rantai)
	fmt.Println("Sepeda berakselerasi:", sepeda.Akselerasi())
	fmt.Println("Setelah sepeda berakselerasi:", sepeda.Rantai)

	fmt.Println()

	fmt.Println("Sebelum mobil berakselerasi:", mobil.Bensin)
	fmt.Println("Mobil berakselerasi:", mobil.Akselerasi())
	fmt.Println("Setelah mobil berakselerasi:", mobil.Bensin)
}
