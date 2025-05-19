package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pengeluaran struct {
	Kategori  string
	Deskripsi string
	Jumlah    float64
}

type Budget struct {
	Anggaran     float64
	Pengeluarans []Pengeluaran
}

func (b *Budget) TambahPengeluaran(p Pengeluaran) {
	b.Pengeluarans = append(b.Pengeluarans, p)
}

func (b *Budget) UbahPengeluaran(index int, p Pengeluaran) {
	if index >= 0 && index < len(b.Pengeluarans) {
		b.Pengeluarans[index] = p
	}
}

func (b *Budget) HapusPengeluaran(index int) {
	if index >= 0 && index < len(b.Pengeluarans) {
		b.Pengeluarans = append(b.Pengeluarans[:index], b.Pengeluarans[index+1:]...)
	}
}

func (b *Budget) TotalPengeluaran() float64 {
	total := 0.0
	for _, p := range b.Pengeluarans {
		total += p.Jumlah
	}
	return total
}

// Fungsi rekursif untuk menghitung total pengeluaran
func TotalRekursif(pengeluarans []Pengeluaran, index int) float64 {
	if index == len(pengeluarans) {
		return 0
	}
	return pengeluarans[index].Jumlah + TotalRekursif(pengeluarans, index+1)
}

func (b *Budget) SaranHemat() string {
	return "Kurangi pengeluaran pada kategori makanan dan transportasi untuk lebih hemat."
}

func SequentialSearch(data []Pengeluaran, kategori string) []Pengeluaran {
	hasil := []Pengeluaran{}
	for _, p := range data {
		if strings.ToLower(p.Kategori) == strings.ToLower(kategori) {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

func BinarySearch(data []Pengeluaran, kategori string) []Pengeluaran {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Kategori < data[j].Kategori
	})

	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid].Kategori == kategori {
			hasil := []Pengeluaran{data[mid]}
			for i := mid - 1; i >= 0 && data[i].Kategori == kategori; i-- {
				hasil = append(hasil, data[i])
			}
			for i := mid + 1; i < len(data) && data[i].Kategori == kategori; i++ {
				hasil = append(hasil, data[i])
			}
			return hasil
		} else if data[mid].Kategori < kategori {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return []Pengeluaran{}
}

func SelectionSortJumlah(data []Pengeluaran) {
	n := len(data)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].Jumlah < data[minIdx].Jumlah {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func InsertionSortKategori(data []Pengeluaran) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Kategori > key.Kategori {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func (b *Budget) Laporan() {
	fmt.Println("Laporan Pengeluaran")
	kategoriMap := make(map[string]float64)
	for _, p := range b.Pengeluarans {
		kategoriMap[p.Kategori] += p.Jumlah
	}
	for k, v := range kategoriMap {
		fmt.Printf("Kategori: %s - Total: %.2f\n", k, v)
	}
	total := TotalRekursif(b.Pengeluarans, 0)
	sel := b.Anggaran - total
	fmt.Printf("Total Pengeluaran (Rekursif): %.2f\n", total)
	fmt.Printf("Selisih Anggaran: %.2f\n", sel)
}

func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func menu() {
	fmt.Println("\n=== Aplikasi Pengelolaan Budget Traveling ===")
	fmt.Println("1. Kelola Pengeluaran")
	fmt.Println("2. Analisis Budget")
	fmt.Println("3. Cari Pengeluaran")
	fmt.Println("4. Urutkan Pengeluaran")
	fmt.Println("5. Laporan Budget")
	fmt.Println("0. Keluar")
}

func main() {
	budget := Budget{Anggaran: 5000000}
	for {
		menu()
		pilih := input("Pilih menu: ")
		switch pilih {
		case "1":
			fmt.Println("1. Tambah, 2. Edit, 3. Hapus")
			a := input("Pilihan: ")
			switch a {
			case "1":
				kat := input("Kategori: ")
				desk := input("Deskripsi: ")
				jumlahStr := input("Jumlah: ")
				jumlah, _ := strconv.ParseFloat(jumlahStr, 64)
				budget.TambahPengeluaran(Pengeluaran{kat, desk, jumlah})
			case "2":
				idxStr := input("Index pengeluaran: ")
				idx, _ := strconv.Atoi(idxStr)
				kat := input("Kategori baru: ")
				desk := input("Deskripsi baru: ")
				jumlahStr := input("Jumlah baru: ")
				jumlah, _ := strconv.ParseFloat(jumlahStr, 64)
				budget.UbahPengeluaran(idx-1, Pengeluaran{kat, desk, jumlah})
			case "3":
				idxStr := input("Index pengeluaran: ")
				idx, _ := strconv.Atoi(idxStr)
				budget.HapusPengeluaran(idx - 1)
			}
		case "2":
			fmt.Printf("Total Pengeluaran: %.2f\n", budget.TotalPengeluaran())
			fmt.Println("Saran: ", budget.SaranHemat())
		case "3":
			kat := input("Kategori yang dicari: ")
			fmt.Println("1. Sequential, 2. Binary")
			opsi := input("Metode: ")
			var hasil []Pengeluaran
			if opsi == "1" {
				hasil = SequentialSearch(budget.Pengeluarans, kat)
			} else {
				hasil = BinarySearch(budget.Pengeluarans, kat)
			}
			for _, p := range hasil {
				fmt.Println(p)
			}
		case "4":
			fmt.Println("1. Berdasarkan Jumlah (Selection)")
			fmt.Println("2. Berdasarkan Kategori (Insertion)")
			opsi := input("Pilihan: ")
			if opsi == "1" {
				SelectionSortJumlah(budget.Pengeluarans)
			} else {
				InsertionSortKategori(budget.Pengeluarans)
			}
			for _, p := range budget.Pengeluarans {
				fmt.Println(p)
			}
		case "5":
			budget.Laporan()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
