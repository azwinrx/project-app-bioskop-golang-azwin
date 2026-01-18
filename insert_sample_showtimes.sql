-- Insert sample showtimes for testing
-- Run this after the main backup to add new showtimes with future dates

-- Showtimes for Cinema 1 (January 25-27, 2026)
INSERT INTO showtimes (movie_id, cinema_id, show_date, show_time, price) VALUES
(1, 1, '2026-01-25', '14:00:00', 50000.00),
(1, 1, '2026-01-25', '19:00:00', 60000.00),
(2, 1, '2026-01-26', '15:30:00', 55000.00),
(2, 1, '2026-01-26', '20:00:00', 65000.00),
(1, 1, '2026-01-27', '16:00:00', 50000.00);

-- Showtimes for Cinema 2 (January 25-27, 2026)
INSERT INTO showtimes (movie_id, cinema_id, show_date, show_time, price) VALUES
(1, 2, '2026-01-25', '13:00:00', 45000.00),
(2, 2, '2026-01-25', '18:00:00', 55000.00),
(1, 2, '2026-01-26', '14:30:00', 45000.00),
(2, 2, '2026-01-27', '19:30:00', 55000.00);

-- Check the inserted data
SELECT 
    s.id,
    m.title as movie,
    c.name as cinema,
    s.show_date,
    s.show_time,
    s.price
FROM showtimes s
JOIN movies m ON s.movie_id = m.id
JOIN cinemas c ON s.cinema_id = c.id
WHERE s.deleted_at IS NULL
ORDER BY s.show_date, s.show_time;
