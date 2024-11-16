## Entity Relationship Diagram (ERD)
Berikut adalah Entity Relationship Diagram (ERD) untuk aplikasi PlantPal yang mengelola data pengguna, tanaman, kondisi tanaman, dan saran perawatan.

### ERD Diagram
![ERD](./assets/image/ERD.png)

### Penjelasan ERD
#### **Tabel dan Relasi**
1. **Users**
   - **Tabel ini menyimpan informasi pengguna aplikasi.**
     - `id` (integer): Primary key.
     - `username` (varchar): Nama pengguna.
     - `email` (varchar): Alamat email pengguna.
     - `password` (varchar): Password pengguna.
   - **Relasi:**
     - Satu pengguna memiliki banyak tanaman (one-to-many dengan tabel Plants).

2. **Plants**
   - **Tabel ini menyimpan data tanaman milik pengguna.**
     - `id` (integer): Primary key.
     - `user_id` (integer): Foreign key ke tabel Users.
     - `plant_name` (varchar): Nama tanaman.
     - `species` (varchar): Jenis tanaman.
     - `lokasi` (varchar): Lokasi tanaman.
   - **Relasi:**
     - Satu tanaman dapat memiliki banyak data kondisi tanaman (one-to-many dengan tabel PlantsConditions).
     - Satu tanaman dapat memiliki banyak saran perawatan (one-to-many dengan tabel CareSuggestions).

3. **PlantsConditions**
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

4. **CareSuggestions**
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

## High-Level Architecture Diagram (HLA)

### **Deskripsi Umum**
High-Level Architecture Diagram (HLA) untuk aplikasi PlantPal memberikan gambaran menyeluruh tentang alur kerja dan komponen utama dalam sistem backend. Diagram ini mencakup proses pengelolaan pengguna, tanaman, pencatatan kondisi tanaman, dan integrasi dengan layanan eksternal untuk rekomendasi perawatan berbasis AI.

### **Komponen Utama**
1. **User**
   - Representasi pengguna aplikasi (melalui web atau aplikasi mobile).
   - Mengirimkan permintaan ke API untuk mengakses fitur aplikasi seperti autentikasi, manajemen tanaman, dan saran perawatan.

2. **API Gateway**
   - Berfungsi sebagai pintu masuk utama untuk semua permintaan API.
   - Mengatur routing, load balancing, dan pengelolaan permintaan ke backend.

3. **Backend (Echo Framework)**
   - Berperan sebagai inti sistem untuk mengelola logika bisnis dan komunikasi dengan komponen lain.
   - **Modul Utama:**
     - **Autentikasi:** Mengelola registrasi dan login pengguna menggunakan JWT untuk keamanan.
     - **Manajemen Tanaman (CRUD):** Fitur untuk menambah, mengedit, melihat, dan menghapus data tanaman.
     - **Pencatatan Kondisi Tanaman:** Menyimpan kondisi tanaman seperti kelembaban, suhu, dan pencahayaan.
     - **Saran Perawatan Berbasis AI:** Mengintegrasikan API eksternal (Gemini) untuk memberikan rekomendasi perawatan.

4. **Database (RDS - PostgreSQL)**
   - Menyimpan data aplikasi termasuk:
     - **Pengguna:** Informasi akun pengguna.
     - **Tanaman:** Data tanaman yang dimiliki oleh pengguna.
     - **Kondisi Tanaman:** Catatan kondisi tanaman yang diinput secara berkala.
     - **Riwayat Saran Perawatan:** Data hasil rekomendasi AI untuk perawatan tanaman.

5. **API Eksternal (Gemini)**
   - Layanan pihak ketiga yang digunakan untuk memberikan saran perawatan berbasis AI berdasarkan data kondisi tanaman yang dikirimkan dari backend.

### **Alur Data**

1. **User → API Gateway**  
   - User mengirimkan permintaan (HTTP Request) ke API Gateway.  
   - API Gateway meneruskan permintaan ke backend.  

2. **Backend → Database**  
   - Backend membaca dan menulis data ke database menggunakan ORM (GORM).  

3. **Backend → Gemini API**  
   - Backend mengirimkan data kondisi tanaman ke Gemini API untuk mendapatkan rekomendasi perawatan.  

4. **Database → Backend**  
   - Data hasil rekomendasi disimpan dan dikirimkan kembali ke user melalui API Gateway.  

### **Diagram Visual**  
![HLA](assets/image/HLA.png)