// Package main adalah package utama yang akan dieksekusi
package main

import (
	"encoding/json" // Encode/decode data ke/dari format JSON
	"fmt"           // Menampilkan output ke terminal (log)
	"html/template" // Render template HTML
	"net/http"      // HTTP server dan handler
	"os"            // Operasi file
	"strconv"       // Konversi string ke tipe data lain
	"strings"       // Manipulasi string
	"sync"          // Sinkronisasi akses data (mutex)
)

// Mahasiswa merepresentasikan entitas mahasiswa
type Mahasiswa struct {
	Nama string // Nama mahasiswa
	Umur int    // Umur mahasiswa
}

// data adalah slice global untuk menyimpan daftar mahasiswa
var data []Mahasiswa

const dataFile = "mahasiswa.json" // Nama file untuk menyimpan data JSON

// Mutex untuk menghindari race condition saat akses data
var mu sync.Mutex

func main() {
	loadDataJSON()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tambah", tambahHandler)
	http.HandleFunc("/hapus", hapusHandler)
	// Untuk static file jika perlu
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type Statistik struct {
	Total    int
	RataRata float64
	Min      int
	Max      int
}

type PageData struct {
	Mahasiswa []Mahasiswa
	Statistik Statistik
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	tmpl := template.Must(template.New("index.html").Funcs(template.FuncMap{"inc": func(i int) int { return i + 1 }}).ParseFiles("templates/index.html"))
	stat := hitungStatistik()
	data := PageData{Mahasiswa: data, Statistik: stat}
	tmpl.Execute(w, data)
}

func tambahHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	nama := r.FormValue("nama")
	umurStr := r.FormValue("umur")
	umur, err := strconv.Atoi(strings.TrimSpace(umurStr))
	if err != nil || nama == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data = append(data, Mahasiswa{Nama: nama, Umur: umur})
	simpanDataJSON()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func hapusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	idxStr := r.FormValue("id")
	idx, err := strconv.Atoi(idxStr)
	if err != nil || idx < 0 || idx >= len(data) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data = append(data[:idx], data[idx+1:]...)
	simpanDataJSON()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func hitungStatistik() Statistik {
	if len(data) == 0 {
		return Statistik{Total: 0, RataRata: 0, Min: 0, Max: 0}
	}
	total := len(data)
	sumUmur := 0
	minUmur := data[0].Umur
	maxUmur := data[0].Umur
	for _, m := range data {
		sumUmur += m.Umur
		if m.Umur < minUmur {
			minUmur = m.Umur
		}
		if m.Umur > maxUmur {
			maxUmur = m.Umur
		}
	}
	rataRata := float64(sumUmur) / float64(total)
	return Statistik{Total: total, RataRata: rataRata, Min: minUmur, Max: maxUmur}
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
