# GORM
## Review
Gorm adalah salah satu ORM yang populer untuk bahasa Go. Sama seperti ORM yang lain, Gorm dapat digunakan untuk memetakan database kedalam objek Go. Dengan ini maka kita bisa lebih mudah untuk berinteraksi dengan database.

## Alur Pembuatan
Untuk menggunakan Gorm kita bisa langsung membuat koneksi ke database yang sudah ada. Namun kita juga bisa mengenerate database menggunakan Gorm. Sehingga alurnya akan menjadi seperti berikut:
1. --- Jika Database Sudah Ada ---
2. Buat koneksi ke database
3. Buat schema struct berdasarkan table yang ada pada database. Setiap table harus dibuatkan struct sendiri
4. --- Jika Database Belum Ada ---
5. Buat schema struct sesuai dengan table yang ingin dibuat (Bisa mengimplementasikan gorm.Model jika ingin mengimplementasikan data-data seperti ID, Created_at dll dalam struct)
6. Buat koneksi ke database
7. Jalankan db.AutoMigrate() dengan mengirimkan struct-struct yang kita buat tadi
8. --- Query dan Exec ---
9. Untuk melakukan query dan execution, kita tinggal memanggil method-method yang ada pada variabel koneksi
10. Untuk memilih tabel mana yang ingin diquery, kita bisa menggunakan struct dan objek dari schema yang kita buat sebelumnya sesuai dengan kebutuhan