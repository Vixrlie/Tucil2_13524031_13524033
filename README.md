# Tucil 2 Strategi Algoritma - Mamba Voxelizer

## Penjelasan Singkat Program
**Mamba Voxelizer** adalah sebuah program berbasis CLI yang ditulis menggunakan bahasa pemrograman Go. Program ini berfungsi untuk melakukan proses *voxelization* terhadap sebuah model 3D (berupa titik dan *face* dari file `.obj`). Program ini mengkonversi poligon 3D  menjadi representasi kubus-kubus *voxel* menggunakan struktur data Octree dan algoritma *Divide and Conquer*. Hasil akhirnya kemudian disimpan ke dalam file output yang namanya disesuaikan oleh pengguna.

## Requirement dan Instalasi
Untuk dapat melakukan kompilasi dan menjalankan program ini, pastikan terinstall:
- [Golang](https://go.dev/dl/) minimal versi **1.26** atau yang lebih baru.
- Tidak ada tambahan third-party extension.

## Cara Mengkompilasi Program
Untuk mengkompilasi program menjadi *executable* *binary*, buka terminal pada _root directory_ program (folder yang berisi `go.mod`), lalu jalankan command berikut:
```bash
go build -o bin/mamba.exe src/main.go
```
(Catatan: gunakan penamaan output `mamba` tanpa `.exe` jika berjalan pada Linux/MacOS).

## Cara Menjalankan dan Menggunakan Program
Terdapat dua cara untuk menjalankan program, melalui *executable* yang telah dikompilasi, atau *run* secara langsung.

**Melalui file executable hasil proses kompilasi:**
```bash
./bin/mamba.exe
```

**Atau jalankan secara instan dengan go run:**
```bash
go run src/main.go
```

**Cara Penggunaan:**
1. Masukkan *file source* Anda (yang berisikan data poligon vertex dan face) ke dalam direktori `tc/`.
2. Jalankan program dengan salah satu *command* di atas.
3. Saat *prompt* `Please enter the filename you desire` muncul, ketikkan nama file beserta ekstensinya (contoh: `cow.obj`).
4. Program akan mulai mengubah model ke dalam bentuk voxel.
5. Setelah selesai, info kedalaman data dan durasi waktu akan ditampilkan. File output hasil akan disimpan ke dalam direktori `test/` (dengan prefix nama `v_`).

## Author / Identitas Pembuat
Program ini dibuat untuk memenuhi Tugas Kecil 2 IF2211 Strategi Algoritma:

| NIM | Nama |
| :---: | :--- |
| **13524031** | Vincent Rionarlie |
| **13524033** | Ray Owen Martin |