// Package main adalah package utama yang akan dieksekusi
package main

import (
	"bufio"         // Membaca input dari user via terminal
	"encoding/json" // Encode/decode data ke/dari format JSON
	"fmt"           // Menampilkan output ke layar
	"os"            // Operasi file dan akses ke sistem operasi
	"strconv"       // Konversi string ke tipe data lain
	"strings"       // Manipulasi string
)

// Mahasiswa merepresentasikan data mahasiswa
type Mahasiswa struct {
	Nama string // Nama mahasiswa
	Umur int    // Umur mahasiswa
}

// data adalah slice untuk menyimpan daftar mahasiswa
var data []Mahasiswa

const dataFile = "mahasiswa.json" // Nama file untuk menyimpan data

// Fungsi utama program
func main() {
	// Load data dari file JSON jika ada
	loadDataJSON()

	scanner := bufio.NewScanner(os.Stdin) // Membaca input dari user
	for {
		// Menampilkan menu utama
		fmt.Println("\n=== Menu CRUD Mahasiswa ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Lihat Data Mahasiswa")
		fmt.Println("3. Ubah Data Mahasiswa")
		fmt.Println("4. Hapus Data Mahasiswa")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		pilihan := scanner.Text()

		// Menangani pilihan menu
		switch pilihan {
		case "1":
			tambahMahasiswa(scanner) // Menambah data mahasiswa
		case "2":
			lihatMahasiswa() // Melihat data mahasiswa
		case "3":
			ubahMahasiswa(scanner) // Mengubah data mahasiswa
		case "4":
			hapusMahasiswa(scanner) // Menghapus data mahasiswa
		case "5":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// tambahMahasiswa menambah data mahasiswa baru ke slice data
func tambahMahasiswa(scanner *bufio.Scanner) {
	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Masukkan umur: ")
	scanner.Scan()
	umurStr := scanner.Text()
	umur, err := strconv.Atoi(strings.TrimSpace(umurStr)) // Konversi umur ke integer
	if err != nil {
		fmt.Println("Umur harus berupa angka.")
		return
	}
	data = append(data, Mahasiswa{Nama: nama, Umur: umur}) // Tambah ke slice data
	simpanDataJSON()                                       // Simpan ke file JSON
	fmt.Println("Data berhasil ditambahkan.")
}

// lihatMahasiswa menampilkan seluruh data mahasiswa
func lihatMahasiswa() {
	if len(data) == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	fmt.Println("\nDaftar Mahasiswa:")
	for i, m := range data {
		fmt.Printf("%d. %s (Umur: %d)\n", i+1, m.Nama, m.Umur)
	}
}

// ubahMahasiswa mengubah data mahasiswa berdasarkan nomor urut
func ubahMahasiswa(scanner *bufio.Scanner) {
	lihatMahasiswa() // Tampilkan data terlebih dahulu
	if len(data) == 0 {
		return
	}
	fmt.Print("Masukkan nomor mahasiswa yang ingin diubah: ")
	scanner.Scan()
	idxStr := scanner.Text()
	idx, err := strconv.Atoi(strings.TrimSpace(idxStr))
	if err != nil || idx < 1 || idx > len(data) {
		fmt.Println("Nomor tidak valid.")
		return
	}
	idx-- // Ubah ke index slice (mulai dari 0)
	fmt.Print("Masukkan nama baru: ")
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Masukkan umur baru: ")
	scanner.Scan()
	umurStr := scanner.Text()
	umur, err := strconv.Atoi(strings.TrimSpace(umurStr))
	if err != nil {
		fmt.Println("Umur harus berupa angka.")
		return
	}
	data[idx] = Mahasiswa{Nama: nama, Umur: umur} // Update data
	simpanDataJSON()                              // Simpan ke file JSON
	fmt.Println("Data berhasil diubah.")
}

// hapusMahasiswa menghapus data mahasiswa berdasarkan nomor urut
func hapusMahasiswa(scanner *bufio.Scanner) {
	lihatMahasiswa() // Tampilkan data terlebih dahulu
	if len(data) == 0 {
		return
	}
	fmt.Print("Masukkan nomor mahasiswa yang ingin dihapus: ")
	scanner.Scan()
	idxStr := scanner.Text()
	idx, err := strconv.Atoi(strings.TrimSpace(idxStr))
	if err != nil || idx < 1 || idx > len(data) {
		fmt.Println("Nomor tidak valid.")
		return
	}
	idx--                                      // Ubah ke index slice (mulai dari 0)
	data = append(data[:idx], data[idx+1:]...) // Hapus data dari slice
	simpanDataJSON()                           // Simpan ke file JSON
	fmt.Println("Data berhasil dihapus.")
}

// simpanDataJSON menyimpan slice data ke file JSON
func simpanDataJSON() {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Gagal menyimpan data ke file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Gagal encode data ke JSON:", err)
	}
}

// loadDataJSON membaca data dari file JSON ke slice data
func loadDataJSON() {
	file, err := os.Open(dataFile)
	if err != nil {
		// File belum ada, tidak masalah
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Gagal decode data dari JSON:", err)
	}
}
