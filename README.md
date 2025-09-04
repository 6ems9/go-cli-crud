# CRUD Mahasiswa (Golang CLI & WEB)

Aplikasi sederhana berbasis command line (CLI) dan WEB untuk melakukan operasi CRUD (Create, Read, Update, Delete) data mahasiswa menggunakan bahasa Go. Data mahasiswa disimpan secara otomatis ke file `mahasiswa.json` dalam format JSON.

## Fitur
- Tambah data mahasiswa (nama & umur)
- Lihat daftar mahasiswa
- Ubah data mahasiswa
- Hapus data mahasiswa
- Statistik jumlah, rata-rata, termuda, dan tertua
- Data otomatis tersimpan dan dimuat dari file JSON

## Cara Menjalankan Versi CLI
1. **Clone repository**
   ```bash
   git clone https://github.com/6ems9/go-cli-crud.git
   cd go-cli-crud
   ```
2. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```

## Cara Menjalankan Versi WEB
1. **Clone repository**
   ```bash
   git clone https://github.com/6ems9/go-cli-crud.git
   cd go-cli-crud
   ```
2. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```
3. **Buka browser** dan akses [http://localhost:8080](http://localhost:8080)

## Struktur Data
Setiap mahasiswa disimpan dalam format berikut:
```json
{
  "Nama": "Nama Mahasiswa",
  "Umur": 20
}
```

## Lisensi
MIT

---

> Dibuat untuk belajar dasar bahasa Go (Pemrograman CLI & WEB).
