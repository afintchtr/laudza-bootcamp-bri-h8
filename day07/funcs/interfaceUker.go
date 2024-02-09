package funcs

import "fmt"

type UnitKerja interface {
	Kur() string
	Tabungan() string
}

type KantorCabang struct {
	Kode   string
	Alamat string
}

type KantorUnit struct {
	Kode   string
	Alamat string
}

func (cabang KantorCabang) Kur() string {
	result := fmt.Sprintf("KUR untuk Cabang %v sebesar 80%%", cabang.Kode)
	return result
}

func (cabang KantorCabang) Tabungan() string {
	result := fmt.Sprintf("Tabungan untuk Cabang beralamat di %v sebesar 80%%", cabang.Alamat)
	return result
}

func (unit KantorUnit) Kur() string {
	result := fmt.Sprintf("KUR untuk Unit %v sebesar 80%%", unit.Kode)
	return result
}

func (unit KantorUnit) Tabungan() string {
	result := fmt.Sprintf("Tabungan untuk Unit beralamat di %v sebesar 80%%", unit.Alamat)
	return result
}

func (unit KantorUnit) Kupedes() string {
	result := fmt.Sprintf("Kupedes untuk Unit %v sebesar 80%%", unit.Kode)
	return result
}

var cabang1 UnitKerja = KantorCabang {
	Kode: "077",
	Alamat: "Jl. Pasar Baru",
}

var unit1 UnitKerja = KantorUnit {
	Kode: "012",
	Alamat: "Jl. Pasar Lama",
}

func PrintAllUkerFuncs() {
	fmt.Println(cabang1.Kur())
	fmt.Println(cabang1.Tabungan())
	fmt.Println(unit1.Kur())
	fmt.Println(unit1.Tabungan())
	fmt.Println(unit1.(KantorUnit).Kupedes())
}