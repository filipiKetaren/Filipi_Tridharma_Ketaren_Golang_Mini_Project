### Project Mini PlantPal

Aplikasi ini adalah proyek mini yang dibangun dengan Golang untuk membantu pengguna dalam merawat tanaman. Fitur utama meliputi:

- Registrasi pengguna.
- Login.
- Pengelolaan data tanaman (CRUD).
- Pencatatan kondisi tanaman, seperti kelembapan, suhu, dan pencahayaan.
- Rekomendasi perawatan berbasis AI yang disesuaikan dengan kondisi tanaman.

## Entity Relationship Diagram (ERD)
Berikut adalah Entity Relationship Diagram (ERD) untuk aplikasi PlantPal yang mengelola data pengguna, tanaman, kondisi tanaman, dan saran perawatan.

### ERD Diagram
![ERD](./assets/image/ERD.png)

### Penjelasan ERD
#### **Tabel dan Relasi**
1. **users**
   - **Tabel ini menyimpan informasi pengguna aplikasi.**
     - `id` (integer): Primary key.
     - `username` (varchar): Nama pengguna.
     - `email` (varchar): Alamat email pengguna.
     - `password` (varchar): Password pengguna.
   - **Relasi:**
     - Satu pengguna memiliki banyak tanaman (one-to-many dengan tabel Plants).

2. **plants**
   - **Tabel ini menyimpan data tanaman milik pengguna.**
     - `id` (integer): Primary key.
     - `user_id` (integer): Foreign key ke tabel Users.
     - `plant_name` (varchar): Nama tanaman.
     - `species` (varchar): Jenis tanaman.
     - `lokasi` (varchar): Lokasi tanaman.
   - **Relasi:**
     - Satu tanaman dapat memiliki banyak data kondisi tanaman (one-to-many dengan tabel PlantsConditions).
     - Satu tanaman dapat memiliki banyak saran perawatan (one-to-many dengan tabel CareSuggestions).

3. **plants_conditions**
   - **Tabel ini menyimpan data kondisi tanaman berdasarkan waktu tertentu.**
     - `id` (integer): Primary key.
     - `plant_id` (integer): Foreign key ke tabel Plants.
     - `date` (date): Tanggal pencatatan kondisi.
     - `moisture_level` (varchar): Tingkat kelembapan.
     - `sunlight_exposure` (varchar): Paparan sinar matahari.
     - `temperature` (int): Suhu lingkungan.
     - `notes` (text): Catatan tambahan.
   - **Relasi:**
     - Satu tanaman dapat memiliki banyak data kondisi.

4. **care_suggestions**
   - **Tabel ini menyimpan saran perawatan untuk tanaman.**
     - `id` (integer): Primary key.
     - `plant_id` (integer): Foreign key ke tabel Plants.
     - `suggestion_text` (text): Teks saran perawatan.
   - **Relasi:**
     - Satu tanaman dapat memiliki banyak saran perawatan.

#### **Relasi Antar Tabel**
- **Users → Plants:** Relasi one-to-many.
- **Plants → PlantsConditions:** Relasi one-to-many.
- **Plants → CareSuggestions:** Relasi one-to-many.

---

## API Documentation

### **Daftar Endpoint Fitur Autentikasi**
| No | Method | Endpoint    | Request Body                                                                                 | Deskripsi                                                                                 |
|----|--------|-------------|---------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| 1  | POST   | `/register` | `{ "username": "filipi", "email": "filipi@gaming.com", "password": "123" }`           | Membuat akun pengguna                                                                     |
| 2  | POST   | `/login`    | `{ "username": "filipi", "email": "filipi@gaming.com", "password": "123" }`           | Membuat sebuah token yang akan digunakan untuk mengakses fitur dari aplikasi              |

### **Daftar Endpoint Fitur Manajemen Tanaman (CRUD)**
| No | Method  | Endpoint    | Request Body                                                                                 | Deskripsi                                                                                 |
|----|---------|-------------|---------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| 1  | GET     | `/plants`   | -                                                                                           | Mendapatkan seluruh data tanaman yang dimiliki oleh pengguna                              |
| 2  | GET     | `/plants/1` | -                                                                                           | Mendapatkan satu data tanaman yang dimiliki oleh pengguna berdasarkan ID                  |
| 3  | POST    | `/plants`   | `{ "plant_name": "Aloe Vera", "species": "Aloe", "location": "Living Room" }`        | Menambahkan data tanaman yang baru                                                       |
| 4  | PUT     | `/plants/1` | `{ "plant_name": "Aloe Vera Bandung", "species": "Aloe Stoe", "location": "Bed Room" }` | Mengubah data tanaman                                                                     |
| 5  | DELETE  | `/plants/1` | -                                                                                           | Menghapus data tanaman                                                                    |

### **Daftar Endpoint Fitur Pencatatan Kondisi Tanaman (CRUD)**
| No | Method  | Endpoint       | Request Body                                                                                                    | Deskripsi                                                                                 |
|----|---------|----------------|----------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| 1  | GET     | `/condition`   | -                                                                                                              | Mendapatkan seluruh data kondisi tanaman yang dimiliki oleh pengguna                      |
| 2  | GET     | `/condition/1` | -                                                                                                              | Mendapatkan satu data kondisi tanaman yang dimiliki oleh pengguna berdasarkan ID          |
| 3  | POST    | `/condition`   | `{ "plant_id": 2, "date": "2024-11-01", "moisture_level": 50, "sunlight_exposure": "medium", "temperature": 25, "notes": "Looks healthy" }` | Menambahkan data kondisi tanaman yang baru                                               |
| 4  | PUT     | `/condition/1` | `{ "plant_id": 2, "date": "2024-11-01", "moisture_level": 40, "sunlight_exposure": "high", "temperature": 28, "notes": "Looks healthy" }` | Mengubah data kondisi tanaman                                                            |
| 5  | DELETE  | `/condition/1` | -                                                                                                              | Menghapus data kondisi tanaman                                                           |

### **Daftar Endpoint Fitur Saran Perawatan Berbasis AI**
| No | Method | Endpoint                 | Request Body | Deskripsi                                                                                 |
|----|--------|--------------------------|--------------|-------------------------------------------------------------------------------------------|
| 1  | GET    | `/plants/1/care-suggestion` | -            | Mendapatkan saran dari AI berdasarkan kondisi tanaman                                      |
| 2  | GET    | `/suggestion`            | -            | Menampilkan seluruh data rekomendasi dari AI                                              |

#### **Cara Menggunakan**
1. Pastikan Anda memiliki Bearer Token dari proses autentikasi `/login`.
2. Sertakan token tersebut di header setiap permintaan yang membutuhkan autentikasi:
   ```plaintext
   Authorization: Bearer <token>
   ```

---
## Contoh Hasil Respon Fitur di Postman

### 1. User Autentikasi

#### Registrasi
![Registrasi](./assets/image/docs_restapi/register.jpg)

#### Login
![Login](./assets/image/docs_restapi/login.jpg)

---

### 2. CRUD Tanaman

#### Menambahkan Catatan Tanaman
![Create Plant](./assets/image/docs_restapi/create%20plant.jpg)

#### Menampilkan Catatan Tanaman
![Find Plant](./assets/image/docs_restapi/find%20plant.jpg)

#### Menampilkan Catatan Tanaman Berdasarkan ID
![Find By Id Plant](./assets/image/docs_restapi/findbyid.jpg)

#### Mengubah Catatan Tanaman
![Update Plant](./assets/image/docs_restapi/update%20plant.jpg)

#### Menghapus Catatan Tanaman
![Delete Plant](./assets/image/docs_restapi/delete%20plant.jpg)

---

### 3. CRUD Kondisi Tanaman

#### Menambahkan Catatan Kondisi Tanaman
![Create Plant Condition](./assets/image/docs_restapi/create%20condition.jpg)

#### Menampilkan Catatan Kondisi Tanaman
![Find Plant Condition](./assets/image/docs_restapi/find%20condition.jpg)

#### Menampilkan Catatan Kondisi Tanaman Berdasarkan ID
![Find By Id Plant Condition](./assets/image/docs_restapi/find%20by%20id%20condition.jpg)

#### Mengubah Catatan Kondisi Tanaman
![Update Plant Condition](./assets/image/docs_restapi/update%20condition.jpg)

#### Menghapus Catatan Kondisi Tanaman
![Delete Plant Condition](./assets/image/docs_restapi/delete%20condition.jpg)

---

### 4. Saran Perawatan Berdasarkan Kondisi dari AI

#### Meminta Saran Perawatan dari AI
![Care Suggestion AI](./assets/image/docs_restapi/get%20suggestion.jpg)

#### Menampilkan Catatan Saran Perawatan dari AI
![Show All Suggestion](./assets/image/docs_restapi/suggestion%20show.jpg)

---

### Catatan

Silahkan Import koleksi Postman yang tersedia di folder `assets/docs/Mini Project.postman_collection.json` untuk mendapatkan hasil respon yang lebih lengkap.

![All Response](./assets/image/docs_restapi/all%20response%202.png)

## High-Level Architecture Diagram (HLA)

### **Deskripsi Umum**
High-Level Architecture Diagram (HLA) untuk aplikasi PlantPal memberikan gambaran menyeluruh tentang alur kerja dan komponen utama dalam sistem backend. Diagram ini mencakup proses pengelolaan pengguna, tanaman, pencatatan kondisi tanaman, dan integrasi dengan layanan eksternal untuk rekomendasi perawatan berbasis AI.

### Komponen Utama

#### **User**
- Mengirimkan permintaan langsung ke backend untuk mengakses fitur aplikasi seperti autentikasi, manajemen tanaman, dan saran perawatan.

#### **Backend (Echo Framework)**
- Berperan sebagai inti sistem untuk mengelola logika bisnis dan komunikasi dengan komponen lain.
- **Modul Utama:**
  - **Autentikasi:** Mengelola registrasi dan login pengguna menggunakan JWT untuk keamanan.
  - **Manajemen Tanaman (CRUD):** Fitur untuk menambah, mengedit, melihat, dan menghapus data tanaman.
  - **Pencatatan Kondisi Tanaman:** Menyimpan kondisi tanaman seperti kelembaban, suhu, dan pencahayaan.
  - **Saran Perawatan Berbasis AI:** Mengintegrasikan API eksternal (Gemini) untuk memberikan rekomendasi perawatan.

#### **Database (RDS - MySQL)**
- Menyimpan data aplikasi termasuk:
  - **Pengguna:** Informasi akun pengguna.
  - **Tanaman:** Data tanaman yang dimiliki oleh pengguna.
  - **Kondisi Tanaman:** Catatan kondisi tanaman yang diinput secara berkala.
  - **Riwayat Saran Perawatan:** Data hasil rekomendasi AI untuk perawatan tanaman.

#### **API Eksternal (Gemini)**
- Layanan pihak ketiga yang digunakan untuk memberikan saran perawatan berbasis AI berdasarkan data kondisi tanaman yang dikirimkan dari backend.

### Alur Data

1. **User → Backend (Echo Framework):**
   - User mengirimkan permintaan langsung ke backend di EC2 (`http://54.209.170.93:8000`).

2. **Backend → Database:**
   - Backend membaca atau menulis data ke database (RDS).

3. **Backend → Gemini API:**
   - Backend mengirimkan permintaan ke layanan eksternal (Gemini API) untuk mendapatkan rekomendasi perawatan tanaman.

### **Diagram Visual**  
![HLA](assets/image/HLA.png)

# Instalasi Mini Project

Berikut adalah langkah-langkah untuk menginstal dan menjalankan Mini Project Golang:

## Langkah Instalasi

1. **Clone Repository**
   ```bash
   git clone https://github.com/filipiKetaren/Filipi_Tridharma_Ketaren_Golang_Mini_Project.git
   ```
   
   Masuk ke direktori proyek:
   ```bash
   cd Filipi_Tridharma_Ketaren_Golang_Mini_Project
   ```

2. **Atur Environment Variables**
   Buat file `.env` di root direktori proyek dan tambahkan konfigurasi berikut:
   ```env
   DB_HOST=localhost/ec2
   DB_PORT=your_port
   DB_USER=your_db_username
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   JWT_SECRET=your_jwt_secret
   API_KEY=your_api_key
   APP_PORT=8000
   ```
   Ganti `your_db_username`, `your_db_password`, dan `your_db_name` dengan informasi yang sesuai untuk database Anda.

3. **Install Dependencies**
   Jalankan perintah berikut untuk menginstal dependensi proyek:
   ```bash
   go mod tidy
   ```

4. **Migrasi Database**
   Jalankan perintah berikut untuk melakukan migrasi database:
   ```bash
   go run main.go migrate
   ```

5. **Menjalankan Aplikasi**
   Jalankan server dengan perintah berikut:
   ```bash
   go run main.go
   ```
   Aplikasi akan berjalan di `http://localhost:8000` secara default.

## Pengujian API

Gunakan **Postman** atau alat API lainnya untuk menguji endpoint yang tersedia. 

1. Import koleksi Postman yang tersedia di folder `assets/docs/Mini Project.postman_collection.json` jika disediakan.
2. Pastikan endpoint dapat diakses sesuai dengan konfigurasi yang telah Anda atur.

## Troubleshooting
- Jika Anda mengalami error saat koneksi ke database, periksa ulang konfigurasi file `.env`.
- Pastikan port yang digunakan tidak bentrok dengan aplikasi lain.

---

Jika Anda memiliki pertanyaan lebih lanjut, silakan hubungi [Filipi Tridharma Ketaren](mailto:filipi.ketaren@gmail.com) atau lihat dokumentasi di repository GitHub.
