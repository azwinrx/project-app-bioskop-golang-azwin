--
-- PostgreSQL database dump
--

\restrict ahkLi4uyxkJK2E7LyJZUFBjMOnjpZadknxNge8ksATxpQxDHfNXJvLvsFrrK41B

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

-- Started on 2026-01-13 15:19:56

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 234 (class 1259 OID 19679)
-- Name: booking_seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.booking_seats (
    id integer NOT NULL,
    booking_id integer,
    seat_id integer
);


ALTER TABLE public.booking_seats OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 19678)
-- Name: booking_seats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.booking_seats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.booking_seats_id_seq OWNER TO postgres;

--
-- TOC entry 5138 (class 0 OID 0)
-- Dependencies: 233
-- Name: booking_seats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.booking_seats_id_seq OWNED BY public.booking_seats.id;


--
-- TOC entry 232 (class 1259 OID 19654)
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    id integer NOT NULL,
    user_id integer,
    showtime_id integer,
    total_price numeric(10,2) NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying,
    booking_code character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 19653)
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bookings_id_seq OWNER TO postgres;

--
-- TOC entry 5139 (class 0 OID 0)
-- Dependencies: 231
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- TOC entry 226 (class 1259 OID 19602)
-- Name: cinemas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cinemas (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    city character varying(100) NOT NULL,
    address text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.cinemas OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 19601)
-- Name: cinemas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cinemas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cinemas_id_seq OWNER TO postgres;

--
-- TOC entry 5140 (class 0 OID 0)
-- Dependencies: 225
-- Name: cinemas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cinemas_id_seq OWNED BY public.cinemas.id;


--
-- TOC entry 222 (class 1259 OID 19576)
-- Name: email_verifications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.email_verifications (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    otp_code character varying(10) NOT NULL,
    expired_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.email_verifications OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 19575)
-- Name: email_verifications_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.email_verifications_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.email_verifications_id_seq OWNER TO postgres;

--
-- TOC entry 5141 (class 0 OID 0)
-- Dependencies: 221
-- Name: email_verifications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.email_verifications_id_seq OWNED BY public.email_verifications.id;


--
-- TOC entry 224 (class 1259 OID 19588)
-- Name: movies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    duration_minutes integer NOT NULL,
    genre character varying(100),
    poster_url character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.movies OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 19587)
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_id_seq OWNER TO postgres;

--
-- TOC entry 5142 (class 0 OID 0)
-- Dependencies: 223
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- TOC entry 236 (class 1259 OID 19697)
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_methods (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    logo_url character varying(255),
    is_active boolean DEFAULT true,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payment_methods OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 19696)
-- Name: payment_methods_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_methods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_id_seq OWNER TO postgres;

--
-- TOC entry 5143 (class 0 OID 0)
-- Dependencies: 235
-- Name: payment_methods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_methods_id_seq OWNED BY public.payment_methods.id;


--
-- TOC entry 238 (class 1259 OID 19708)
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id integer NOT NULL,
    booking_id integer,
    payment_method_id integer,
    amount numeric(10,2) NOT NULL,
    transaction_time timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    status character varying(20) DEFAULT 'success'::character varying
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 19707)
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payments_id_seq OWNER TO postgres;

--
-- TOC entry 5144 (class 0 OID 0)
-- Dependencies: 237
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- TOC entry 228 (class 1259 OID 19616)
-- Name: seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seats (
    id integer NOT NULL,
    cinema_id integer,
    seat_number character varying(10) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.seats OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 19615)
-- Name: seats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_id_seq OWNER TO postgres;

--
-- TOC entry 5145 (class 0 OID 0)
-- Dependencies: 227
-- Name: seats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.seats_id_seq OWNED BY public.seats.id;


--
-- TOC entry 230 (class 1259 OID 19631)
-- Name: showtimes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.showtimes (
    id integer NOT NULL,
    movie_id integer,
    cinema_id integer,
    show_date date NOT NULL,
    show_time time without time zone NOT NULL,
    price numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.showtimes OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 19630)
-- Name: showtimes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.showtimes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.showtimes_id_seq OWNER TO postgres;

--
-- TOC entry 5146 (class 0 OID 0)
-- Dependencies: 229
-- Name: showtimes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.showtimes_id_seq OWNED BY public.showtimes.id;


--
-- TOC entry 220 (class 1259 OID 19556)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    is_verified boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 19555)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 5147 (class 0 OID 0)
-- Dependencies: 219
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4922 (class 2604 OID 19682)
-- Name: booking_seats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seats ALTER COLUMN id SET DEFAULT nextval('public.booking_seats_id_seq'::regclass);


--
-- TOC entry 4918 (class 2604 OID 19657)
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- TOC entry 4910 (class 2604 OID 19605)
-- Name: cinemas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas ALTER COLUMN id SET DEFAULT nextval('public.cinemas_id_seq'::regclass);


--
-- TOC entry 4905 (class 2604 OID 19579)
-- Name: email_verifications id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.email_verifications ALTER COLUMN id SET DEFAULT nextval('public.email_verifications_id_seq'::regclass);


--
-- TOC entry 4907 (class 2604 OID 19591)
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- TOC entry 4923 (class 2604 OID 19700)
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_methods_id_seq'::regclass);


--
-- TOC entry 4926 (class 2604 OID 19711)
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- TOC entry 4913 (class 2604 OID 19619)
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seats_id_seq'::regclass);


--
-- TOC entry 4915 (class 2604 OID 19634)
-- Name: showtimes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes ALTER COLUMN id SET DEFAULT nextval('public.showtimes_id_seq'::regclass);


--
-- TOC entry 4901 (class 2604 OID 19559)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 5128 (class 0 OID 19679)
-- Dependencies: 234
-- Data for Name: booking_seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.booking_seats (id, booking_id, seat_id) FROM stdin;
\.


--
-- TOC entry 5126 (class 0 OID 19654)
-- Dependencies: 232
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bookings (id, user_id, showtime_id, total_price, status, booking_code, created_at, updated_at) FROM stdin;
\.


--
-- TOC entry 5120 (class 0 OID 19602)
-- Dependencies: 226
-- Data for Name: cinemas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cinemas (id, name, city, address, created_at, updated_at, deleted_at) FROM stdin;
1	CGV Grand Indonesia	Jakarta	Jl. M.H. Thamrin No.1	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
2	XXI Summarecon Bekasi	Bekasi	Summarecon Mall Bekasi Lt. 3	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
\.


--
-- TOC entry 5116 (class 0 OID 19576)
-- Dependencies: 222
-- Data for Name: email_verifications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.email_verifications (id, email, otp_code, expired_at, created_at) FROM stdin;
\.


--
-- TOC entry 5118 (class 0 OID 19588)
-- Dependencies: 224
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.movies (id, title, description, duration_minutes, genre, poster_url, created_at, updated_at, deleted_at) FROM stdin;
1	Avengers: Infinity War	The Avengers and their allies must be willing to sacrifice all in an attempt to defeat the powerful Thanos.	149	Action, Sci-Fi	https://example.com/avengers.jpg	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
2	Avatar: The Way of Water	Jake Sully lives with his newfound family formed on the extrasolar moon Pandora.	192	Adventure, Sci-Fi	https://example.com/avatar.jpg	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
3	Gundala	Indonesia local superhero based on comics.	123	Action	https://example.com/gundala.jpg	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
\.


--
-- TOC entry 5130 (class 0 OID 19697)
-- Dependencies: 236
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_methods (id, name, logo_url, is_active, created_at) FROM stdin;
1	Gopay	https://example.com/gopay.png	t	2026-01-10 20:55:21.211625
2	OVO	https://example.com/ovo.png	t	2026-01-10 20:55:21.211625
3	Virtual Account BCA	https://example.com/bca.png	t	2026-01-10 20:55:21.211625
\.


--
-- TOC entry 5132 (class 0 OID 19708)
-- Dependencies: 238
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payments (id, booking_id, payment_method_id, amount, transaction_time, status) FROM stdin;
\.


--
-- TOC entry 5122 (class 0 OID 19616)
-- Dependencies: 228
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.seats (id, cinema_id, seat_number, created_at) FROM stdin;
1	1	A1	2026-01-12 01:07:41.726215
2	1	B1	2026-01-12 01:07:41.726215
3	1	C1	2026-01-12 01:07:41.726215
4	1	D1	2026-01-12 01:07:41.726215
5	1	E1	2026-01-12 01:07:41.726215
6	1	A2	2026-01-12 01:07:41.726215
7	1	B2	2026-01-12 01:07:41.726215
8	1	C2	2026-01-12 01:07:41.726215
9	1	D2	2026-01-12 01:07:41.726215
10	1	E2	2026-01-12 01:07:41.726215
11	1	A3	2026-01-12 01:07:41.726215
12	1	B3	2026-01-12 01:07:41.726215
13	1	C3	2026-01-12 01:07:41.726215
14	1	D3	2026-01-12 01:07:41.726215
15	1	E3	2026-01-12 01:07:41.726215
16	1	A4	2026-01-12 01:07:41.726215
17	1	B4	2026-01-12 01:07:41.726215
18	1	C4	2026-01-12 01:07:41.726215
19	1	D4	2026-01-12 01:07:41.726215
20	1	E4	2026-01-12 01:07:41.726215
21	1	A5	2026-01-12 01:07:41.726215
22	1	B5	2026-01-12 01:07:41.726215
23	1	C5	2026-01-12 01:07:41.726215
24	1	D5	2026-01-12 01:07:41.726215
25	1	E5	2026-01-12 01:07:41.726215
26	1	A6	2026-01-12 01:07:41.726215
27	1	B6	2026-01-12 01:07:41.726215
28	1	C6	2026-01-12 01:07:41.726215
29	1	D6	2026-01-12 01:07:41.726215
30	1	E6	2026-01-12 01:07:41.726215
31	1	A7	2026-01-12 01:07:41.726215
32	1	B7	2026-01-12 01:07:41.726215
33	1	C7	2026-01-12 01:07:41.726215
34	1	D7	2026-01-12 01:07:41.726215
35	1	E7	2026-01-12 01:07:41.726215
36	1	A8	2026-01-12 01:07:41.726215
37	1	B8	2026-01-12 01:07:41.726215
38	1	C8	2026-01-12 01:07:41.726215
39	1	D8	2026-01-12 01:07:41.726215
40	1	E8	2026-01-12 01:07:41.726215
41	1	A9	2026-01-12 01:07:41.726215
42	1	B9	2026-01-12 01:07:41.726215
43	1	C9	2026-01-12 01:07:41.726215
44	1	D9	2026-01-12 01:07:41.726215
45	1	E9	2026-01-12 01:07:41.726215
46	1	A10	2026-01-12 01:07:41.726215
47	1	B10	2026-01-12 01:07:41.726215
48	1	C10	2026-01-12 01:07:41.726215
49	1	D10	2026-01-12 01:07:41.726215
50	1	E10	2026-01-12 01:07:41.726215
51	2	A1	2026-01-12 01:07:41.726215
52	2	B1	2026-01-12 01:07:41.726215
53	2	C1	2026-01-12 01:07:41.726215
54	2	D1	2026-01-12 01:07:41.726215
55	2	E1	2026-01-12 01:07:41.726215
56	2	A2	2026-01-12 01:07:41.726215
57	2	B2	2026-01-12 01:07:41.726215
58	2	C2	2026-01-12 01:07:41.726215
59	2	D2	2026-01-12 01:07:41.726215
60	2	E2	2026-01-12 01:07:41.726215
61	2	A3	2026-01-12 01:07:41.726215
62	2	B3	2026-01-12 01:07:41.726215
63	2	C3	2026-01-12 01:07:41.726215
64	2	D3	2026-01-12 01:07:41.726215
65	2	E3	2026-01-12 01:07:41.726215
66	2	A4	2026-01-12 01:07:41.726215
67	2	B4	2026-01-12 01:07:41.726215
68	2	C4	2026-01-12 01:07:41.726215
69	2	D4	2026-01-12 01:07:41.726215
70	2	E4	2026-01-12 01:07:41.726215
71	2	A5	2026-01-12 01:07:41.726215
72	2	B5	2026-01-12 01:07:41.726215
73	2	C5	2026-01-12 01:07:41.726215
74	2	D5	2026-01-12 01:07:41.726215
75	2	E5	2026-01-12 01:07:41.726215
76	2	A6	2026-01-12 01:07:41.726215
77	2	B6	2026-01-12 01:07:41.726215
78	2	C6	2026-01-12 01:07:41.726215
79	2	D6	2026-01-12 01:07:41.726215
80	2	E6	2026-01-12 01:07:41.726215
81	2	A7	2026-01-12 01:07:41.726215
82	2	B7	2026-01-12 01:07:41.726215
83	2	C7	2026-01-12 01:07:41.726215
84	2	D7	2026-01-12 01:07:41.726215
85	2	E7	2026-01-12 01:07:41.726215
86	2	A8	2026-01-12 01:07:41.726215
87	2	B8	2026-01-12 01:07:41.726215
88	2	C8	2026-01-12 01:07:41.726215
89	2	D8	2026-01-12 01:07:41.726215
90	2	E8	2026-01-12 01:07:41.726215
91	2	A9	2026-01-12 01:07:41.726215
92	2	B9	2026-01-12 01:07:41.726215
93	2	C9	2026-01-12 01:07:41.726215
94	2	D9	2026-01-12 01:07:41.726215
95	2	E9	2026-01-12 01:07:41.726215
96	2	A10	2026-01-12 01:07:41.726215
97	2	B10	2026-01-12 01:07:41.726215
98	2	C10	2026-01-12 01:07:41.726215
99	2	D10	2026-01-12 01:07:41.726215
100	2	E10	2026-01-12 01:07:41.726215
\.


--
-- TOC entry 5124 (class 0 OID 19631)
-- Dependencies: 230
-- Data for Name: showtimes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.showtimes (id, movie_id, cinema_id, show_date, show_time, price, created_at, updated_at, deleted_at) FROM stdin;
1	1	1	2026-01-14	14:00:00	50000.00	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
2	1	1	2026-01-14	19:00:00	60000.00	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
3	2	2	2026-01-15	15:30:00	45000.00	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
\.


--
-- TOC entry 5114 (class 0 OID 19556)
-- Dependencies: 220
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, email, password, is_verified, created_at, updated_at, deleted_at) FROM stdin;
1	azwinrifai	azwin@example.com	hashed_secret_password	t	2026-01-10 20:55:21.211625	2026-01-10 20:55:21.211625	\N
\.


--
-- TOC entry 5148 (class 0 OID 0)
-- Dependencies: 233
-- Name: booking_seats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.booking_seats_id_seq', 1, false);


--
-- TOC entry 5149 (class 0 OID 0)
-- Dependencies: 231
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_id_seq', 1, false);


--
-- TOC entry 5150 (class 0 OID 0)
-- Dependencies: 225
-- Name: cinemas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cinemas_id_seq', 2, true);


--
-- TOC entry 5151 (class 0 OID 0)
-- Dependencies: 221
-- Name: email_verifications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.email_verifications_id_seq', 1, false);


--
-- TOC entry 5152 (class 0 OID 0)
-- Dependencies: 223
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_id_seq', 3, true);


--
-- TOC entry 5153 (class 0 OID 0)
-- Dependencies: 235
-- Name: payment_methods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_methods_id_seq', 3, true);


--
-- TOC entry 5154 (class 0 OID 0)
-- Dependencies: 237
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payments_id_seq', 1, false);


--
-- TOC entry 5155 (class 0 OID 0)
-- Dependencies: 227
-- Name: seats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seats_id_seq', 100, true);


--
-- TOC entry 5156 (class 0 OID 0)
-- Dependencies: 229
-- Name: showtimes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.showtimes_id_seq', 3, true);


--
-- TOC entry 5157 (class 0 OID 0)
-- Dependencies: 219
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- TOC entry 4952 (class 2606 OID 19685)
-- Name: booking_seats booking_seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seats
    ADD CONSTRAINT booking_seats_pkey PRIMARY KEY (id);


--
-- TOC entry 4947 (class 2606 OID 19667)
-- Name: bookings bookings_booking_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_booking_code_key UNIQUE (booking_code);


--
-- TOC entry 4949 (class 2606 OID 19665)
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- TOC entry 4940 (class 2606 OID 19614)
-- Name: cinemas cinemas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);


--
-- TOC entry 4936 (class 2606 OID 19586)
-- Name: email_verifications email_verifications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.email_verifications
    ADD CONSTRAINT email_verifications_pkey PRIMARY KEY (id);


--
-- TOC entry 4938 (class 2606 OID 19600)
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- TOC entry 4954 (class 2606 OID 19706)
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (id);


--
-- TOC entry 4956 (class 2606 OID 19717)
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- TOC entry 4942 (class 2606 OID 19624)
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (id);


--
-- TOC entry 4945 (class 2606 OID 19642)
-- Name: showtimes showtimes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_pkey PRIMARY KEY (id);


--
-- TOC entry 4930 (class 2606 OID 19574)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4932 (class 2606 OID 19570)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4934 (class 2606 OID 19572)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 4950 (class 1259 OID 19729)
-- Name: idx_bookings_user; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_user ON public.bookings USING btree (user_id);


--
-- TOC entry 4943 (class 1259 OID 19728)
-- Name: idx_showtimes_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_showtimes_date ON public.showtimes USING btree (show_date);


--
-- TOC entry 4962 (class 2606 OID 19686)
-- Name: booking_seats booking_seats_booking_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seats
    ADD CONSTRAINT booking_seats_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(id);


--
-- TOC entry 4963 (class 2606 OID 19691)
-- Name: booking_seats booking_seats_seat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seats
    ADD CONSTRAINT booking_seats_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seats(id);


--
-- TOC entry 4960 (class 2606 OID 19673)
-- Name: bookings bookings_showtime_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_showtime_id_fkey FOREIGN KEY (showtime_id) REFERENCES public.showtimes(id);


--
-- TOC entry 4961 (class 2606 OID 19668)
-- Name: bookings bookings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 4964 (class 2606 OID 19718)
-- Name: payments payments_booking_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(id);


--
-- TOC entry 4965 (class 2606 OID 19723)
-- Name: payments payments_payment_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id);


--
-- TOC entry 4957 (class 2606 OID 19625)
-- Name: seats seats_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id);


--
-- TOC entry 4958 (class 2606 OID 19648)
-- Name: showtimes showtimes_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id);


--
-- TOC entry 4959 (class 2606 OID 19643)
-- Name: showtimes showtimes_movie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id);


-- Completed on 2026-01-13 15:19:56

--
-- PostgreSQL database dump complete
--

\unrestrict ahkLi4uyxkJK2E7LyJZUFBjMOnjpZadknxNge8ksATxpQxDHfNXJvLvsFrrK41B

