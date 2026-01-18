-- Cinema Booking System Database Schema

-- Table: users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_users_email ON users(email);

-- Table: cinemas
CREATE TABLE IF NOT EXISTS cinemas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_cinemas_city ON cinemas(city);

-- Table: movies
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    genre VARCHAR(50),
    poster_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Table: showtimes
CREATE TABLE IF NOT EXISTS showtimes (
    id SERIAL PRIMARY KEY,
    movie_id INTEGER NOT NULL REFERENCES movies(id),
    cinema_id INTEGER NOT NULL REFERENCES cinemas(id),
    show_date DATE NOT NULL,
    show_time TIME NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_showtimes_cinema_date ON showtimes(cinema_id, show_date);
CREATE INDEX idx_showtimes_movie ON showtimes(movie_id);

-- Table: seats
CREATE TABLE IF NOT EXISTS seats (
    id SERIAL PRIMARY KEY,
    cinema_id INTEGER NOT NULL REFERENCES cinemas(id),
    seat_number VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(cinema_id, seat_number)
);

CREATE INDEX idx_seats_cinema ON seats(cinema_id);

-- Table: bookings
CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    showtime_id INTEGER NOT NULL REFERENCES showtimes(id),
    total_price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, confirmed, cancelled
    booking_code VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bookings_user ON bookings(user_id);
CREATE INDEX idx_bookings_showtime ON bookings(showtime_id);
CREATE INDEX idx_bookings_code ON bookings(booking_code);

-- Table: booking_seats
CREATE TABLE IF NOT EXISTS booking_seats (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    seat_id INTEGER NOT NULL REFERENCES seats(id),
    UNIQUE(booking_id, seat_id)
);

CREATE INDEX idx_booking_seats_booking ON booking_seats(booking_id);
CREATE INDEX idx_booking_seats_seat ON booking_seats(seat_id);

-- Table: payment_methods
CREATE TABLE IF NOT EXISTS payment_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: payments
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER NOT NULL REFERENCES bookings(id),
    payment_method_id INTEGER NOT NULL REFERENCES payment_methods(id),
    amount DECIMAL(10, 2) NOT NULL,
    transaction_time TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, success, failed
    UNIQUE(booking_id)
);

CREATE INDEX idx_payments_booking ON payments(booking_id);

-- Sample Data Insertion

-- Insert sample cinemas
INSERT INTO cinemas (name, city, address) VALUES
('Cinema XXI Grand Indonesia', 'Jakarta', 'Jl. MH Thamrin No.1, Jakarta Pusat'),
('CGV Blitz Paris Van Java', 'Bandung', 'Jl. Sukajadi No.137-139, Bandung'),
('Cinepolis Surabaya Town Square', 'Surabaya', 'Jl. Adityawarman No.55, Surabaya');

-- Insert sample movies
INSERT INTO movies (title, description, duration_minutes, genre, poster_url) VALUES
('Avatar: The Way of Water', 'Sekuel epik dari Avatar dengan dunia bawah laut Pandora', 192, 'Sci-Fi', 'https://example.com/avatar.jpg'),
('Top Gun: Maverick', 'Pete Mitchell kembali sebagai instruktur di Top Gun', 130, 'Action', 'https://example.com/topgun.jpg'),
('The Batman', 'Batman menghadapi Riddler dalam Gotham City', 176, 'Action', 'https://example.com/batman.jpg');

-- Insert sample showtimes (adjust dates as needed)
INSERT INTO showtimes (movie_id, cinema_id, show_date, show_time, price) VALUES
(1, 1, '2026-01-25', '14:00', 50000),
(1, 1, '2026-01-25', '17:00', 50000),
(1, 1, '2026-01-25', '20:00', 50000),
(2, 2, '2026-01-25', '15:00', 45000),
(2, 2, '2026-01-25', '18:00', 45000),
(3, 3, '2026-01-25', '16:00', 55000),
(3, 3, '2026-01-25', '19:00', 55000);

-- Insert sample seats for each cinema
-- Cinema 1 seats (A1-A10, B1-B10)
INSERT INTO seats (cinema_id, seat_number) 
SELECT 1, seat_num FROM (
    SELECT 'A' || generate_series(1, 10) AS seat_num
    UNION ALL
    SELECT 'B' || generate_series(1, 10)
) AS seats;

-- Cinema 2 seats
INSERT INTO seats (cinema_id, seat_number) 
SELECT 2, seat_num FROM (
    SELECT 'A' || generate_series(1, 10) AS seat_num
    UNION ALL
    SELECT 'B' || generate_series(1, 10)
) AS seats;

-- Cinema 3 seats
INSERT INTO seats (cinema_id, seat_number) 
SELECT 3, seat_num FROM (
    SELECT 'A' || generate_series(1, 10) AS seat_num
    UNION ALL
    SELECT 'B' || generate_series(1, 10)
) AS seats;

-- Insert payment methods
INSERT INTO payment_methods (name, description) VALUES
('Credit Card', 'Pembayaran menggunakan kartu kredit'),
('Debit Card', 'Pembayaran menggunakan kartu debit'),
('E-Wallet', 'Pembayaran menggunakan dompet digital (GoPay, OVO, Dana)'),
('Bank Transfer', 'Pembayaran melalui transfer bank'),
('Cash', 'Pembayaran tunai di kasir');
