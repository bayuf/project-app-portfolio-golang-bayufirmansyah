--
-- PostgreSQL database dump
--

\restrict WmmBesn5qY9PcXFs7liJd591ctHy1gJChAfhPe9GchjVIeK9AyIdECYHfPA8MjS

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

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
-- Name: address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.address (
    id integer NOT NULL,
    user_id integer NOT NULL,
    address text NOT NULL,
    city character varying(100) NOT NULL,
    province character varying(100),
    postal_code character varying(10),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    country character varying(30) DEFAULT 'Indonesia'::character varying NOT NULL
);


ALTER TABLE public.address OWNER TO postgres;

--
-- Name: address_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.address_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.address_id_seq OWNER TO postgres;

--
-- Name: address_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.address_id_seq OWNED BY public.address.id;


--
-- Name: auth_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_users (
    id integer NOT NULL,
    email character varying(150) NOT NULL,
    password_hash text NOT NULL,
    last_login_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone
);


ALTER TABLE public.auth_users OWNER TO postgres;

--
-- Name: auth_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.auth_users_id_seq OWNER TO postgres;

--
-- Name: auth_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_users_id_seq OWNED BY public.auth_users.id;


--
-- Name: feedbacks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.feedbacks (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    message text NOT NULL,
    profile_img text DEFAULT 'public/assets/img/testimonials/user.png'::text,
    created_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.feedbacks OWNER TO postgres;

--
-- Name: feedbacks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.feedbacks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.feedbacks_id_seq OWNER TO postgres;

--
-- Name: feedbacks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.feedbacks_id_seq OWNED BY public.feedbacks.id;


--
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.messages (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(150) NOT NULL,
    subject character varying(150) NOT NULL,
    message text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.messages OWNER TO postgres;

--
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.messages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.messages_id_seq OWNER TO postgres;

--
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- Name: offers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.offers (
    id integer NOT NULL,
    user_id integer NOT NULL,
    title character varying(100) NOT NULL,
    description text NOT NULL,
    icon_url text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.offers OWNER TO postgres;

--
-- Name: offers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.offers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.offers_id_seq OWNER TO postgres;

--
-- Name: offers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.offers_id_seq OWNED BY public.offers.id;


--
-- Name: profiles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profiles (
    id integer NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    headline character varying(150),
    about text,
    email character varying(100),
    phone character varying(20),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    experience integer DEFAULT 0,
    CONSTRAINT profiles_experience_check CHECK ((experience >= 0))
);


ALTER TABLE public.profiles OWNER TO postgres;

--
-- Name: profiles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.profiles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.profiles_id_seq OWNER TO postgres;

--
-- Name: profiles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.profiles_id_seq OWNED BY public.profiles.id;


--
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id integer NOT NULL,
    user_id integer NOT NULL,
    title character varying(100) NOT NULL,
    description text NOT NULL,
    icon_url text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    link_url text
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.projects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.projects_id_seq OWNER TO postgres;

--
-- Name: projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.projects_id_seq OWNED BY public.projects.id;


--
-- Name: skills; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.skills (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying(100) NOT NULL,
    level character varying(20) NOT NULL,
    img_url text NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.skills OWNER TO postgres;

--
-- Name: skills_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.skills_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.skills_id_seq OWNER TO postgres;

--
-- Name: skills_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.skills_id_seq OWNED BY public.skills.id;


--
-- Name: address id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.address ALTER COLUMN id SET DEFAULT nextval('public.address_id_seq'::regclass);


--
-- Name: auth_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_users ALTER COLUMN id SET DEFAULT nextval('public.auth_users_id_seq'::regclass);


--
-- Name: feedbacks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.feedbacks ALTER COLUMN id SET DEFAULT nextval('public.feedbacks_id_seq'::regclass);


--
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- Name: offers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.offers ALTER COLUMN id SET DEFAULT nextval('public.offers_id_seq'::regclass);


--
-- Name: profiles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles ALTER COLUMN id SET DEFAULT nextval('public.profiles_id_seq'::regclass);


--
-- Name: projects id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects ALTER COLUMN id SET DEFAULT nextval('public.projects_id_seq'::regclass);


--
-- Name: skills id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills ALTER COLUMN id SET DEFAULT nextval('public.skills_id_seq'::regclass);


--
-- Data for Name: address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.address (id, user_id, address, city, province, postal_code, created_at, updated_at, deleted_at, country) FROM stdin;
1	1	Wringinanom	Gresik	Jawa Timur	61176	2025-12-23 10:26:16.294638+07	2025-12-23 10:26:16.294638+07	\N	Indonesia
\.


--
-- Data for Name: auth_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_users (id, email, password_hash, last_login_at, created_at, updated_at) FROM stdin;
1	bayu19fr@gmail.com	$2a$10$o70eTFFnGJz6F55RUGlncesJ72BFCBwJLE0YoJC1OKjFsyjlSf2TG	2025-12-23 16:40:02.714701+07	2025-12-23 16:40:02.714701+07	2025-12-23 16:40:02.714701+07
\.


--
-- Data for Name: feedbacks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.feedbacks (id, name, message, profile_img, created_at, deleted_at) FROM stdin;
1	Andi Pratama	Website yang dibuat sangat rapi, cepat, dan mudah digunakan. Struktur backend-nya juga sangat baik.	public/assets/img/testimonials/t1.jpg	2025-12-23 14:10:15.930333+07	\N
2	Budi Santoso	Sistem yang dikembangkan stabil dan mudah dikembangkan kembali. Komunikasi juga sangat baik.	public/assets/img/testimonials/t2.jpg	2025-12-23 14:10:15.930333+07	\N
3	Siti Rahmawati	Pengerjaan project tepat waktu dengan hasil yang sesuai kebutuhan. Sangat direkomendasikan.	public/assets/img/testimonials/user.png	2025-12-23 14:10:36.758489+07	\N
4	Ahmad Fauzi	Website portofolio tampil rapi dan profesional. Performa backend juga sangat baik.	public/assets/img/testimonials/user.png	2025-12-23 14:32:38.029948+07	\N
5	Nadia Putri	Struktur aplikasi mudah dipahami dan dikembangkan. Hasil kerja sangat memuaskan.	public/assets/img/testimonials/user.png	2025-12-23 14:32:38.029948+07	\N
6	Rizky Pratama	Pengerjaan project sesuai kebutuhan dengan komunikasi yang jelas dan responsif.	public/assets/img/testimonials/user.png	2025-12-23 14:32:38.029948+07	\N
\.


--
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.messages (id, name, email, subject, message, created_at, deleted_at) FROM stdin;
1	Bayu Firmansyah	bayu@gmail.com	Say Hallo	Hallo Apakah kita bisa saling mengenal?	2025-12-23 15:25:22.503045+07	\N
2	Andi	andi@student.ac.id	Belajar Golang	Tolong Ajari aku Pemrograman Golang	2025-12-23 15:30:46.404651+07	\N
3	Bayu Firmansyah	bayufirmansyah203@gmail.com	Belajar Golang	Ayo belajar Golang	2025-12-23 22:45:41.633558+07	\N
\.


--
-- Data for Name: offers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.offers (id, user_id, title, description, icon_url, created_at, updated_at, deleted_at) FROM stdin;
1	1	Backend Development	Pengembangan backend yang aman, cepat, dan mudah dikembangkan menggunakan Golang serta PostgreSQL.	public/assets/img/services/s1.png	2025-12-23 13:04:34.441622+07	2025-12-23 13:04:34.441622+07	\N
2	1	Network Engineer	Perancangan, konfigurasi, dan pemeliharaan jaringan komputer yang stabil dan aman.	public/assets/img/services/s2.png	2025-12-23 13:04:34.441622+07	2025-12-23 13:04:34.441622+07	\N
3	1	SEO Optimization	Optimasi website agar lebih mudah ditemukan di mesin pencari dan meningkatkan visibilitas bisnis.	public/assets/img/services/s3.png	2025-12-23 13:04:34.441622+07	2025-12-23 13:04:34.441622+07	\N
4	1	Web Development	Pembuatan website modern, responsif, dan cepat sesuai kebutuhan personal maupun bisnis.	public/assets/img/services/s4.png	2025-12-23 13:04:34.441622+07	2025-12-23 13:04:34.441622+07	\N
\.


--
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profiles (id, first_name, last_name, headline, about, email, phone, created_at, updated_at, deleted_at, experience) FROM stdin;
1	Bayu	Firmansyah	Backend Engineer | Network Engineer	Saya adalah backend engineer yang fokus pada pengembangan REST API menggunakan Golang. Terbiasa membangun aplikasi dengan arsitektur yang rapi, mudah dirawat, dan scalable.	bayu19fr@email.com	+62 89687390746	2025-12-21 14:07:51.058045+07	2025-12-21 14:07:51.058045+07	\N	1
\.


--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, user_id, title, description, icon_url, created_at, updated_at, deleted_at, link_url) FROM stdin;
1	1	Portofolio Website	Website portofolio dinamis dengan backend Golang, terintegrasi dengan database untuk pengelolaan konten.	public/assets/img/portfolio/p1.jpg	2025-12-23 13:39:31.115292+07	2025-12-23 13:39:31.115292+07	\N	https://github.com/bayuf/project-app-portfolio-golang-bayufirmansyah
2	1	Inventory System	Sistem inventaris berbasis web menggunakan Golang dan PostgreSQL untuk pengelolaan data barang secara terstruktur.	public/assets/img/portfolio/p2.jpg	2025-12-23 13:39:31.115292+07	2025-12-23 13:39:31.115292+07	\N	https://github.com/bayuf/project-app-inventaris-cli-bayufirmansyah
3	1	Todo List Golang CLI	Aplikasi Todo List berbasis Command Line Interface menggunakan Golang dan Cobra CLI.	public/assets/img/portfolio/p3.jpg	2025-12-23 13:39:31.115292+07	2025-12-23 13:39:31.115292+07	\N	https://github.com/bayuf/project-app-todo-list-cli-bayufirmansyah
\.


--
-- Data for Name: skills; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.skills (id, user_id, name, level, img_url, created_at, updated_at, deleted_at) FROM stdin;
1	1	HTML	intermediate	public/assets/img/brands/logo1.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
2	1	CSS	intermediate	public/assets/img/brands/logo2.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
3	1	JavaScript	intermediate	public/assets/img/brands/logo3.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
4	1	Golang	intermediate	public/assets/img/brands/logo4.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
5	1	PostgreSQL	intermediate	public/assets/img/brands/logo5.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
6	1	MongoDB	intermediate	public/assets/img/brands/logo6.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
7	1	Redis	intermediate	public/assets/img/brands/logo7.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
8	1	Go-Chi	intermediate	public/assets/img/brands/logo8.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
9	1	Mikrotik	intermediate	public/assets/img/brands/logo9.png	2025-12-23 11:25:32.458489+07	2025-12-23 11:25:32.458489+07	\N
\.


--
-- Name: address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.address_id_seq', 1, true);


--
-- Name: auth_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_users_id_seq', 1, true);


--
-- Name: feedbacks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.feedbacks_id_seq', 6, true);


--
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.messages_id_seq', 3, true);


--
-- Name: offers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.offers_id_seq', 4, true);


--
-- Name: profiles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.profiles_id_seq', 1, true);


--
-- Name: projects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.projects_id_seq', 3, true);


--
-- Name: skills_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.skills_id_seq', 9, true);


--
-- Name: address address_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.address
    ADD CONSTRAINT address_pkey PRIMARY KEY (id);


--
-- Name: auth_users auth_users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_users
    ADD CONSTRAINT auth_users_email_key UNIQUE (email);


--
-- Name: auth_users auth_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_users
    ADD CONSTRAINT auth_users_pkey PRIMARY KEY (id);


--
-- Name: feedbacks feedbacks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.feedbacks
    ADD CONSTRAINT feedbacks_pkey PRIMARY KEY (id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: offers offers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.offers
    ADD CONSTRAINT offers_pkey PRIMARY KEY (id);


--
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: skills skills_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT skills_pkey PRIMARY KEY (id);


--
-- Name: address fk_address_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.address
    ADD CONSTRAINT fk_address_user FOREIGN KEY (user_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: offers fk_offers_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.offers
    ADD CONSTRAINT fk_offers_user FOREIGN KEY (user_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: profiles fk_profiles_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT fk_profiles_user FOREIGN KEY (id) REFERENCES public.auth_users(id) ON DELETE CASCADE;


--
-- Name: projects fk_projects_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT fk_projects_user FOREIGN KEY (user_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: skills fk_skills_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT fk_skills_user FOREIGN KEY (user_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict WmmBesn5qY9PcXFs7liJd591ctHy1gJChAfhPe9GchjVIeK9AyIdECYHfPA8MjS

