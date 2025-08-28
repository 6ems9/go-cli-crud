# CRUD Mahasiswa (Golang CLI)

Aplikasi sederhana berbasis command line (CLI) untuk melakukan operasi CRUD (Create, Read, Update, Delete) data mahasiswa menggunakan bahasa Go. Data mahasiswa disimpan secara otomatis ke file `mahasiswa.json` dalam format JSON.

## Fitur
- Tambah data mahasiswa (nama & umur)
- Lihat daftar mahasiswa
- Ubah data mahasiswa
- Hapus data mahasiswa
- Data otomatis tersimpan dan dimuat dari file JSON

## Cara Menjalankan
1. **Clone repository**
   ```bash
   git clone https://github.com/6ems9/go-cli-crud.git
   cd go-cli-crud
   ```
2. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```

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

> Dibuat untuk belajar dasar bahasa Go dan pemrograman CLI.
