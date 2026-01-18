# Quick Start Guide - Cinema Booking System

## 🚀 Setup Cepat dalam 5 Menit

### 1. Install Dependencies

```bash
cd project-app-bioskop-golang-azwin
go mod download
```

### 2. Setup Database

```bash
# Buat database PostgreSQL
createdb cinema_booking

# Import schema dan sample data
psql -d cinema_booking -f schema.sql
```

### 3. Konfigurasi Environment

```bash
# Copy .env.example ke .env
cp .env.example .env

# Edit .env dan sesuaikan dengan konfigurasi database Anda:
# DATABASE_NAME=cinema_booking
# DATABASE_USERNAME=postgres
# DATABASE_PASSWORD=your_password
# DATABASE_HOST=localhost
```

### 4. Jalankan Aplikasi

```bash
go run main.go
```

Server akan berjalan di: `http://localhost:8080`

---

## 🧪 Testing API

### Menggunakan cURL

**1. Register User:**

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
```

**2. Login:**

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

Simpan `token` dari response untuk request selanjutnya.

**3. Get Cinemas:**

```bash
curl http://localhost:8080/api/cinemas
```

**4. Check Seat Availability:**

```bash
curl "http://localhost:8080/api/cinemas/1/seats?date=2026-01-25&time=14:00"
```

**5. Create Booking (dengan token):**

```bash
curl -X POST http://localhost:8080/api/booking \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{"cinema_id":1,"seat_id":1,"date":"2026-01-25","time":"14:00","payment_method":1}'
```

### Menggunakan Postman

1. Import file `Cinema_Booking_API.postman_collection.json` ke Postman
2. Set environment variable `base_url` ke `http://localhost:8080`
3. Jalankan request secara berurutan

---

## 📁 Struktur Project

```
project-app-bioskop-golang-azwin/
├── cmd/                          # Entry points aplikasi
├── internal/
│   ├── adaptor/
│   │   ├── handler/             # HTTP handlers (controllers)
│   │   │   ├── auth_handler.go
│   │   │   ├── cinema_handler.go
│   │   │   ├── booking_handler.go
│   │   │   └── payment_handler.go
│   │   └── routes/              # Routing configuration
│   │       └── routes.go
│   ├── data/
│   │   ├── entity/              # Database models
│   │   └── repository/          # Data access layer
│   │       ├── users.go
│   │       ├── cinemas.go
│   │       ├── seats.go
│   │       ├── showtimes.go
│   │       ├── bookings.go
│   │       ├── booking_seats.go
│   │       ├── payments.go
│   │       └── payment_methods.go
│   ├── dto/                     # Data Transfer Objects
│   ├── middleware/              # HTTP middlewares
│   │   ├── auth.go
│   │   └── logging.go
│   ├── usecase/                 # Business logic
│   │   ├── auth.go
│   │   ├── cinemas.go
│   │   ├── bookings.go
│   │   └── payments.go
│   └── wire/                    # Dependency injection
│       └── wire.go
├── pkg/
│   ├── database/                # Database connection
│   └── utils/                   # Utilities
│       ├── config.go
│       ├── jwt.go
│       ├── logger.go
│       ├── password_hash.go
│       ├── response.go
│       └── validator.go
├── logs/                        # Application logs
├── .env.example                 # Environment template
├── schema.sql                   # Database schema
├── main.go                      # Application entry point
├── go.mod                       # Go modules
├── API_DOCUMENTATION.md         # API documentation
└── README.md                    # Main documentation
```

---

## 📚 API Endpoints

### Public Endpoints (No Authentication Required)

- `POST /api/register` - Register user
- `POST /api/login` - Login user
- `GET /api/cinemas` - Get all cinemas
- `GET /api/cinemas/{id}` - Get cinema by ID
- `GET /api/cinemas/{id}/seats` - Get seat availability
- `GET /api/payment-methods` - Get payment methods

### Protected Endpoints (Authentication Required)

- `POST /api/logout` - Logout user
- `POST /api/booking` - Create booking
- `GET /api/user/bookings` - Get user booking history
- `POST /api/pay` - Process payment

---

## 🔑 Environment Variables

```env
# Application
APP_NAME=Cinema Booking System
PORT=8080
DEBUG=true
LIMIT=10
PATH_LOGGING=./logs

# Database
DATABASE_NAME=cinema_booking
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=postgres
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_MAX_CONN=20

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRATION=24h
```

---

## ✅ Fitur yang Sudah Diimplementasikan

- ✅ User Registration dengan password hashing (bcrypt)
- ✅ User Login dengan JWT token
- ✅ User Logout
- ✅ Get daftar cinema dengan pagination
- ✅ Get detail cinema
- ✅ Check ketersediaan kursi berdasarkan cinema, date, time
- ✅ Create booking (dengan authentication)
- ✅ Get booking history user (dengan authentication)
- ✅ Get payment methods
- ✅ Process payment (dengan authentication)
- ✅ Middleware untuk authentication (JWT)
- ✅ Middleware untuk logging (Zap)
- ✅ Input validation (go-playground/validator)
- ✅ Repository pattern
- ✅ Clean Architecture
- ✅ Structured logging dengan Zap
- ✅ Configuration management dengan Viper
- ✅ Database connection pooling dengan pgx/v5
- ✅ HTTP routing dengan Chi
- ✅ Standardized JSON response

---

## 🐛 Troubleshooting

### Database Connection Error

```
Error: failed to connect to postgres database
```

**Solusi:** Pastikan PostgreSQL sudah running dan kredensial di `.env` sudah benar.

### Port Already in Use

```
Error: bind: address already in use
```

**Solusi:** Ubah `PORT` di file `.env` atau matikan aplikasi yang menggunakan port 8080.

### Token Invalid/Expired

```
Error: invalid or expired token
```

**Solusi:** Login ulang untuk mendapatkan token baru.

---

## 📞 Support

Untuk pertanyaan atau issue, silakan buat issue di repository atau hubungi maintainer.

---

## 📝 Next Steps

Setelah aplikasi berjalan, Anda bisa:

1. Explore API menggunakan Postman collection yang disediakan
2. Membaca dokumentasi lengkap di `API_DOCUMENTATION.md`
3. Lihat database schema di `schema.sql`
4. Customize sesuai kebutuhan project Anda

Happy Coding! 🎉
