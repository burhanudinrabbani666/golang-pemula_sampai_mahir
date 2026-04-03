# Benchmark

- Selain unit test, Go-Lang testing package juga mendukung melakukan benchmark
- Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
- Benchmark di Go-Lang dilakukan dengan cara secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
- Kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari testing package

## testing.B

- testing.B adalah struct yang digunakan untuk melakukan benchmark.
- testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Error(), Fatal() dan lain-lain
- Yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
- Salah satunya adalah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark

## Cara Kerja Benchmark

- Cara kerja benchmark di Go-Lang sangat sederhana
- Gimana kita hanya perlu membuat perulangan sejumlah N attribute
- Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmark nya dalam waktu

Next: [Membuat benchmark](./13-membuat-benchmark.md)
