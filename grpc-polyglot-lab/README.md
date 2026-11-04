# LANGKAH PRAKTIKUM gRPC POLYGLOT

### A. Persiapan Lingkungan Praktikum

Sebelum memulai praktikum, siapkan beberapa perangkat lunak yang dibutuhkan untuk membangun komunikasi gRPC menggunakan beberapa bahasa pemrograman.

1. **Git**: Digunakan untuk mengelola versi kode dan repositori proyek.
2. **Python versi 3.10 atau lebih baru**: Digunakan untuk membangun aplikasi client Python.
3. **Node.js versi LTS**: Digunakan untuk membangun aplikasi client Node.js.
4. **Go (Golang)**: Digunakan untuk membangun server gRPC.
5. **Protocol Buffers Compiler (protoc)**: Digunakan untuk mengompilasi file `.proto`.

#### Pemeriksaan Instalasi

Buka terminal atau Git Bash, kemudian jalankan perintah berikut untuk memastikan semua perangkat lunak telah terinstal.

```bash
git --version
python --version
node -v
npm -v
go version
protoc --version
```

### B. Inisialisasi Workspace dan Manajemen Repositori Git

Tahap ini dilakukan untuk membuat direktori utama proyek dan menginisialisasi repositori Git.

**1. Membuat direktori proyek**

```bash
mkdir grpc-polyglot-lab
cd grpc-polyglot-lab
```

**2. Inisialisasi repositori Git**

```bash
git init
```

**3. Membuat file `.gitignore`**

File ini digunakan untuk mengabaikan file sementara dan direktori dependensi agar tidak ikut disimpan ke repositori.

```bash
echo "node_modules/" >> .gitignore
echo "__pycache__/" >> .gitignore
echo ".venv/" >> .gitignore
echo "*.exe" >> .gitignore
```

**4. Menyimpan commit awal**

```bash
git add .gitignore
git commit -m "initial commit: add gitignore rules"
```

### C. Perancangan Kontrak API (`user.proto`)

Tahap ini menggunakan pendekatan *schema-first development*, yaitu mendefinisikan struktur komunikasi antar layanan terlebih dahulu sebelum membuat kode program.

**1. Membuat file `user.proto`**

Buat file bernama `user.proto` pada direktori utama proyek.

**2. Menuliskan struktur Protocol Buffers**


### D. Implementasi gRPC Server Menggunakan Go

Tahap ini bertujuan membangun server gRPC menggunakan bahasa Go yang menerima permintaan dari client dan mengirimkan data profil pengguna.

**1. Membuat direktori server Go**

```bash
mkdir go-server
cd go-server
```

**2. Inisialisasi modul Go**

```bash
go mod init go-server
```

**3. Menginstal dependensi gRPC**

```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

**4. Menginstal plugin generator Go**

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Pastikan direktori instalasi Go (`GOPATH/bin`) sudah masuk ke PATH sistem.

**5. Mengompilasi file Protocol Buffers**

Jalankan perintah berikut dari direktori `go-server`:

```bash
protoc --proto_path=../ --go_out=. --go-grpc_out=. ../user.proto
```

Proses ini menghasilkan kode Go berdasarkan struktur yang telah didefinisikan pada file `user.proto`.

**6. Membuat file `main.go`**

Buat file `main.go` di dalam direktori `go-server`, kemudian masukkan kode implementasi server Go sesuai dengan kontrak `UserService`.

Server bertugas menerima `user_id`, memeriksa data pengguna, dan mengembalikan informasi nama, email, serta nomor telepon.

**7. Menjalankan server Go**

```bash
go mod tidy
go run main.go
```

Jika berhasil, server akan berjalan pada port `50051`.

### E. Implementasi gRPC Client Menggunakan Python

Tahap ini membangun aplikasi client Python untuk mengirimkan permintaan data kepada server Go.

**1. Kembali ke direktori utama**

```bash
cd ..
mkdir python-client
cd python-client
```

**2. Menginstal pustaka gRPC Python**

```bash
pip install grpcio grpcio-tools
```

**3. Mengompilasi file Protocol Buffers**

```bash
python -m grpc_tools.protoc --proto_path=../ --python_out=. --grpc_python_out=. ../user.proto
```

Perintah tersebut menghasilkan file `user_pb2.py` dan `user_pb2_grpc.py` yang digunakan oleh client Python.

**4. Membuat file `client.py`**

Buat file `client.py`, kemudian tuliskan kode program untuk menghubungkan client Python dengan server Go melalui alamat `localhost:50051`.

Client mengirimkan permintaan dengan `user_id = 1` dan menampilkan hasil respons berupa nama, email, dan nomor telepon pengguna.

**5. Menjalankan client Python**

```bash
python client.py
```

### F. Implementasi gRPC Client Menggunakan Node.js

Tahap ini membangun client alternatif menggunakan Node.js untuk menguji komunikasi lintas bahasa pemrograman.

**1. Kembali ke direktori utama**

```bash
cd ..
mkdir node-client
cd node-client
```

**2. Inisialisasi proyek Node.js**

```bash
npm init -y
```

**3. Menginstal dependensi gRPC**

```bash
npm install @grpc/grpc-js @grpc/proto-loader
```

**4. Membuat file `client.js`**

Buat file `client.js` di dalam direktori `node-client`.

File ini digunakan untuk membaca skema `user.proto` secara langsung menggunakan pustaka `@grpc/proto-loader`, kemudian menghubungkan client Node.js dengan server Go.

Client mengirimkan permintaan menggunakan `user_id = 1` dan menampilkan data nama, email, serta nomor telepon.

**5. Menjalankan client Node.js**

```bash
node client.js
```

### G. Pengujian dan Hasil Verifikasi Sistem Polyglot

Pengujian dilakukan untuk memastikan server Go dapat menerima permintaan dari dua client berbeda, yaitu Python dan Node.js.

Pengujian membutuhkan tiga terminal yang dijalankan secara terpisah.

**1. Terminal 1 – Menjalankan Server Go**

```bash
cd go-server
go run main.go
```

Output yang diharapkan:

```text
gRPC server (Go) berjalan di port :50051...
```

**2. Terminal 2 – Menjalankan Client Python**

```bash
cd python-client
python client.py
```

Output yang diharapkan:

```text
Mengirim request ke Go server menggunakan Python...
Respon dari server:
Nama  : riatunnisa
Email : riatunnisa@universitas.ac.id
Phone : +6281234567890
```

**3. Terminal 3 – Menjalankan Client Node.js**

```bash
cd node-client
node client.js
```

Output yang diharapkan:

```text
Mengirim request ke Go server menggunakan Node.js...
Respon dari server:
Nama  : riatunnisa
Email : riatunnisa@universitas.ac.id
Phone : +6281234567890
```

Jika kedua client berhasil menampilkan data yang sama, berarti komunikasi gRPC antara bahasa Go, Python, dan Node.js telah berhasil dilakukan.

---

## LATIHAN

### A. Modifikasi Kontrak Data (`user.proto`)

Latihan pertama dilakukan dengan menambahkan field baru berupa nomor telepon pada struktur pesan `UserResponse`.

**1. Membuka file `user.proto`**

Buka file `user.proto` pada direktori utama proyek.

**2. Menambahkan field phone**

Tambahkan field berikut pada struktur `UserResponse`:

```protobuf
string phone = 3;
```

Penambahan field tersebut bertujuan agar server dapat mengirimkan informasi nomor telepon kepada client.

### B. Kompilasi Ulang File Protocol Buffers

Setelah struktur kontrak data diperbarui, lakukan kompilasi ulang agar kode perantara pada setiap bahasa pemrograman sesuai dengan perubahan terbaru.

**1. Regenerasi kode Go**

Jalankan dari direktori `go-server`:

```bash
protoc --proto_path=../ --go_out=. --go-grpc_out=. ../user.proto
```

**2. Regenerasi kode Python**

Jalankan dari direktori `python-client`:

```bash
python -m grpc_tools.protoc --proto_path=../ --python_out=. --grpc_python_out=. ../user.proto
```

Proses regenerasi diperlukan agar field `phone` dapat dikenali oleh server dan client.

### C. Penyesuaian Implementasi Server dan Client

Setelah proses kompilasi ulang selesai, lakukan penyesuaian pada kode program.

**1. Penyesuaian server Go**

Buka file `main.go`, kemudian sesuaikan fungsi RPC agar mengembalikan informasi nomor telepon pengguna.

Contoh penambahan data:

```go
return &pb.UserResponse{
    Name:  "riatunnisa",
    Email: "riatunnisa@universitas.ac.id",
    Phone: "081234567890",
}, nil
```

**2. Penyesuaian client Python**

Buka file `client.py`, kemudian tambahkan atribut `phone` pada bagian output.

```python
print("Respon dari server:")
print(f"Nama  : {response.name}")
print(f"Email : {response.email}")
print(f"Phone : {response.phone}")
```

**3. Pengujian kembali**

Jalankan server Go terlebih dahulu, kemudian jalankan client Python dan Node.js untuk memastikan field `phone` berhasil dikirim dan diterima.

### D. Analisis Pendekatan Kompilasi (`README.md`)

Pada tahap ini dilakukan analisis mengenai dua pendekatan pemrosesan file Protocol Buffers, yaitu *static code generation* dan *dynamic runtime generation*.

**1. Static Code Generation**

Pendekatan ini digunakan pada bahasa Go dan Python dengan bantuan kompiler `protoc` atau `grpc_tools.protoc`.

Kelebihannya adalah menghasilkan kode perantara sebelum aplikasi dijalankan, mendukung keamanan tipe data (*type safety*), serta membantu mendeteksi kesalahan lebih awal.

Kekurangannya adalah memerlukan proses kompilasi ulang setiap kali terdapat perubahan pada file `.proto`.

**2. Dynamic Runtime Generation**

Pendekatan ini digunakan pada Node.js dengan bantuan pustaka `@grpc/proto-loader`.

File `.proto` dibaca dan dimuat secara langsung ketika aplikasi dijalankan tanpa perlu membuat kode perantara secara terpisah.

Kelebihannya adalah lebih fleksibel dan praktis ketika melakukan perubahan skema. Kekurangannya adalah pemeriksaan kesalahan struktur dan tipe data lebih banyak dilakukan ketika aplikasi berjalan (*runtime*).

**Kesimpulan Analisis:**

Static code generation lebih sesuai ketika membutuhkan pemeriksaan tipe data dan kode perantara yang dihasilkan sebelum aplikasi dijalankan. Sementara itu, dynamic runtime generation memberikan fleksibilitas lebih tinggi karena file `.proto` dapat dibaca secara langsung saat aplikasi dijalankan.

### E. Commit dan Push ke GitHub

Tahap terakhir adalah menyimpan seluruh perubahan proyek ke dalam repositori Git dan mengunggahnya ke GitHub.

**1. Memeriksa perubahan file**

```bash
git status
```

**2. Menambahkan seluruh perubahan**

```bash
git add .
```

**3. Membuat commit**

```bash
git commit -m "feat: tambahkan field phone pada user.proto dan sesuaikan server client"
```

**4. Mengunggah perubahan ke GitHub**

```bash
git push origin main
```

Pastikan seluruh perubahan berhasil dikirim ke repositori GitHub.

---

