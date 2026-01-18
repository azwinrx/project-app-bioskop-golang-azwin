# Cinema Booking System - RESTful API

Aplikasi Cinema Booking System berbasis RESTful API menggunakan bahasa pemrograman Golang. Aplikasi ini membantu pengguna (customer) dalam melakukan pendaftaran akun, login, memilih bioskop, mengecek ketersediaan kursi, melakukan pemesanan kursi, serta melakukan pembayaran tiket.

## Teknologi yang Digunakan

- **Go 1.25.3** - Bahasa pemrograman utama
- **Chi Router** - HTTP router untuk RESTful API
- **PostgreSQL** - Database relational
- **pgx/v5** - PostgreSQL driver
- **Zap** - Structured logging
- **Viper** - Configuration management
- **JWT** - Token-based authentication
- **Validator** - Input validation
- **Bcrypt** - Password hashing

## Arsitektur

Project menggunakan **Clean Architecture** dengan layer:
- **Handler/Controller** - HTTP handlers
- **Usecase/Service** - Business logic
- **Repository** - Data access layer
- **Entity** - Database models
- **DTO** - Data Transfer Objects
- **Middleware** - HTTP middlewares

## Fitur

### 1. User Registration and Authentication
- POST /api/register - Mendaftarkan customer baru
- POST /api/login - Login customer (returns JWT token)
- POST /api/logout - Logout customer (requires authentication)

### 2. Cinema Selection
- GET /api/cinemas - Mendapatkan daftar semua bioskop (dengan pagination)
- GET /api/cinemas/{cinemaId} - Mendapatkan detail bioskop

### 3. Seat Availability
- GET /api/cinemas/{cinemaId}/seats?date={date}&time={time} - Mendapatkan status kursi

### 4. Booking Seat
- POST /api/booking - Melakukan pemesanan kursi (requires authentication)

### 5. Payment Methods
- GET /api/payment-methods - Mendapatkan daftar metode pembayaran
- POST /api/pay - Memproses pembayaran (requires authentication)

### 6. User Booking History
- GET /api/user/bookings - Mendapatkan riwayat pemesanan (requires authentication)

## Instalasi dan Setup

### Prerequisites
- Go 1.25.3 atau lebih baru
- PostgreSQL 12 atau lebih baru

### Langkah Instalasi

1. Clone Repository dan Install Dependencies
```bash
cd project-app-bioskop-golang-azwin
go mod download
go mod tidy
```

2. Setup Database
```bash
# Buat database
createdb cinema_booking

# Jalankan schema
psql -d cinema_booking -f schema.sql
```

3. Setup Environment Variables
```bash
cp .env.example .env
# Edit .env sesuai konfigurasi
```

4. Jalankan Aplikasi
```bash
go run main.go
```

Server akan berjalan di http://localhost:8080

## Database Schema

Lihat file `schema.sql` untuk detail lengkap database schema dan sample data.

## API Documentation

Lihat section API Endpoints di bawah untuk detail lengkap setiap endpoint.

## License

MIT License
