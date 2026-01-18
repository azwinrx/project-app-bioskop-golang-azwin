# Cinema Booking System - API Documentation

Base URL: `http://localhost:8080`

## Table of Contents

- [Authentication](#authentication)
- [Cinemas](#cinemas)
- [Seats](#seats)
- [Bookings](#bookings)
- [Payments](#payments)
- [Error Handling](#error-handling)

---

## Authentication

### Register User

Creates a new user account.

**Endpoint:** `POST /api/register`

**Request Body:**

```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (201 Created):**

```json
{
  "status": true,
  "message": "user registered successfully",
  "data": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com"
  }
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

---

### Login

Authenticates a user and returns a JWT token.

**Endpoint:** `POST /api/login`

**Request Body:**

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (200 OK):**

```json
{
  "status": true,
  "message": "login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "john_doe",
      "email": "john@example.com"
    }
  }
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

---

### Logout

Logs out the authenticated user.

**Endpoint:** `POST /api/logout`

**Headers:**

```
Authorization: Bearer <token>
```

**Response (200 OK):**

```json
{
  "status": true,
  "message": "logout successful",
  "data": null
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/api/logout \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

## Cinemas

### Get All Cinemas

Retrieves a paginated list of all cinemas.

**Endpoint:** `GET /api/cinemas`

**Query Parameters:**

- `page` (optional, default: 1) - Page number
- `limit` (optional, default: 10) - Items per page

**Response (200 OK):**

```json
{
  "status": true,
  "message": "cinemas retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Cinema XXI Grand Indonesia",
      "city": "Jakarta",
      "address": "Jl. MH Thamrin No.1, Jakarta Pusat"
    },
    {
      "id": 2,
      "name": "CGV Blitz Paris Van Java",
      "city": "Bandung",
      "address": "Jl. Sukajadi No.137-139, Bandung"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total_pages": 1,
    "total_items": 3
  }
}
```

**cURL Example:**

```bash
curl http://localhost:8080/api/cinemas?page=1&limit=10
```

---

### Get Cinema by ID

Retrieves details of a specific cinema.

**Endpoint:** `GET /api/cinemas/{cinemaId}`

**Path Parameters:**

- `cinemaId` (required) - Cinema ID

**Response (200 OK):**

```json
{
  "status": true,
  "message": "cinema retrieved successfully",
  "data": {
    "id": 1,
    "name": "Cinema XXI Grand Indonesia",
    "city": "Jakarta",
    "address": "Jl. MH Thamrin No.1, Jakarta Pusat"
  }
}
```

**cURL Example:**

```bash
curl http://localhost:8080/api/cinemas/1
```

---

## Seats

### Get Seats Availability

Retrieves seat availability for a specific cinema, date, and time.

**Endpoint:** `GET /api/cinemas/{cinemaId}/seats`

**Path Parameters:**

- `cinemaId` (required) - Cinema ID

**Query Parameters:**

- `date` (required) - Show date in YYYY-MM-DD format (e.g., 2026-01-25)
- `time` (required) - Show time in HH:MM format (e.g., 14:00)

**Response (200 OK):**

```json
{
  "status": true,
  "message": "seats availability retrieved successfully",
  "data": {
    "cinema_id": 1,
    "cinema_name": "Cinema XXI Grand Indonesia",
    "show_date": "2026-01-25",
    "show_time": "14:00",
    "seats": [
      {
        "id": 1,
        "seat_number": "A1",
        "is_booked": false
      },
      {
        "id": 2,
        "seat_number": "A2",
        "is_booked": true
      },
      {
        "id": 3,
        "seat_number": "A3",
        "is_booked": false
      }
    ]
  }
}
```

**cURL Example:**

```bash
curl "http://localhost:8080/api/cinemas/1/seats?date=2026-01-25&time=14:00"
```

---

## Bookings

### Create Booking

Creates a new booking for a seat. **Requires Authentication.**

**Endpoint:** `POST /api/booking`

**Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**

```json
{
  "cinema_id": 1,
  "seat_id": 1,
  "date": "2026-01-25",
  "time": "14:00",
  "payment_method": 1
}
```

**Response (201 Created):**

```json
{
  "status": true,
  "message": "booking created successfully",
  "data": {
    "booking_id": 1,
    "booking_code": "BK-abc12345",
    "cinema_id": 1,
    "seat_number": "A1",
    "show_date": "2026-01-25",
    "show_time": "14:00",
    "total_price": 50000,
    "status": "pending"
  }
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/api/booking \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "cinema_id": 1,
    "seat_id": 1,
    "date": "2026-01-25",
    "time": "14:00",
    "payment_method": 1
  }'
```

---

### Get User Bookings

Retrieves booking history for the authenticated user. **Requires Authentication.**

**Endpoint:** `GET /api/user/bookings`

**Headers:**

```
Authorization: Bearer <token>
```

**Response (200 OK):**

```json
{
  "status": true,
  "message": "user bookings retrieved successfully",
  "data": [
    {
      "booking_id": 1,
      "booking_code": "BK-abc12345",
      "cinema_name": "Cinema XXI Grand Indonesia",
      "show_date": "2026-01-25",
      "show_time": "14:00",
      "seat_numbers": ["A1"],
      "total_price": 50000,
      "status": "confirmed",
      "booking_date": "2026-01-18 10:30:00"
    }
  ]
}
```

**cURL Example:**

```bash
curl http://localhost:8080/api/user/bookings \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

---

## Payments

### Get Payment Methods

Retrieves all available payment methods.

**Endpoint:** `GET /api/payment-methods`

**Response (200 OK):**

```json
{
  "status": true,
  "message": "payment methods retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Credit Card",
      "description": "Pembayaran menggunakan kartu kredit"
    },
    {
      "id": 2,
      "name": "Debit Card",
      "description": "Pembayaran menggunakan kartu debit"
    },
    {
      "id": 3,
      "name": "E-Wallet",
      "description": "Pembayaran menggunakan dompet digital (GoPay, OVO, Dana)"
    }
  ]
}
```

**cURL Example:**

```bash
curl http://localhost:8080/api/payment-methods
```

---

### Process Payment

Processes payment for a booking. **Requires Authentication.**

**Endpoint:** `POST /api/pay`

**Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**

```json
{
  "booking_id": 1,
  "payment_method_id": 1
}
```

**Response (200 OK):**

```json
{
  "status": true,
  "message": "payment processed successfully",
  "data": {
    "payment_id": 1,
    "booking_id": 1,
    "amount": 50000,
    "payment_method": "Credit Card",
    "status": "success",
    "transaction_time": "2026-01-18 10:35:00"
  }
}
```

**cURL Example:**

```bash
curl -X POST http://localhost:8080/api/pay \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "booking_id": 1,
    "payment_method_id": 1
  }'
```

---

## Error Handling

### Error Response Format

All error responses follow this format:

```json
{
  "status": false,
  "message": "error message",
  "errors": "detailed error information"
}
```

### Common HTTP Status Codes

- **200 OK** - Request successful
- **201 Created** - Resource created successfully
- **400 Bad Request** - Invalid request data
- **401 Unauthorized** - Authentication required or invalid token
- **403 Forbidden** - Access denied
- **404 Not Found** - Resource not found
- **500 Internal Server Error** - Server error

### Example Error Responses

**400 Bad Request - Validation Error:**

```json
{
  "status": false,
  "message": "validation failed",
  "errors": "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

**401 Unauthorized:**

```json
{
  "status": false,
  "message": "unauthorized",
  "errors": "invalid or expired token"
}
```

**404 Not Found:**

```json
{
  "status": false,
  "message": "cinema not found",
  "errors": "cinema not found"
}
```

---

## Testing Flow

### Complete Booking Flow Example

1. **Register a new user:**

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'
```

2. **Login to get token:**

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

3. **Get list of cinemas:**

```bash
curl http://localhost:8080/api/cinemas
```

4. **Check seat availability:**

```bash
curl "http://localhost:8080/api/cinemas/1/seats?date=2026-01-25&time=14:00"
```

5. **Create booking:**

```bash
curl -X POST http://localhost:8080/api/booking \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"cinema_id":1,"seat_id":1,"date":"2026-01-25","time":"14:00","payment_method":1}'
```

6. **Get payment methods:**

```bash
curl http://localhost:8080/api/payment-methods
```

7. **Process payment:**

```bash
curl -X POST http://localhost:8080/api/pay \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"booking_id":1,"payment_method_id":1}'
```

8. **Check booking history:**

```bash
curl http://localhost:8080/api/user/bookings \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## Notes

- All dates should be in `YYYY-MM-DD` format
- All times should be in `HH:MM` format (24-hour)
- JWT tokens expire after 24 hours by default
- Tokens must be included in the `Authorization` header as `Bearer <token>`
- All request and response bodies use JSON format
- Pagination is available for endpoints returning lists
