--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

-- Started on 2025-05-14 16:18:01

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

--
-- TOC entry 856 (class 1247 OID 16404)
-- Name: case_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.case_status AS ENUM (
    'abierta',
    'cerrado',
    'en espera'
);


ALTER TYPE public.case_status OWNER TO postgres;

--
-- TOC entry 853 (class 1247 OID 16398)
-- Name: user_role; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.user_role AS ENUM (
    'user',
    'support'
);


ALTER TYPE public.user_role OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 16418)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(25) NOT NULL,
    color character varying(25)
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16417)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.categories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 224 (class 1259 OID 16442)
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.messages (
    id integer NOT NULL,
    content character varying(150) NOT NULL,
    sent_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    support_case_id integer NOT NULL,
    user_id integer NOT NULL,
    sender character varying(25) NOT NULL
);


ALTER TABLE public.messages OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16441)
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.messages ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 222 (class 1259 OID 16424)
-- Name: support_cases; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.support_cases (
    id integer NOT NULL,
    title character varying(50) NOT NULL,
    description character varying(150) NOT NULL,
    status public.case_status DEFAULT 'abierta'::public.case_status NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    user_id integer NOT NULL,
    category_id integer
);


ALTER TABLE public.support_cases OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16423)
-- Name: support_case_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.support_cases ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.support_case_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 218 (class 1259 OID 16412)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(25) NOT NULL,
    password character varying(25) NOT NULL,
    role public.user_role NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16411)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 4930 (class 0 OID 16418)
-- Dependencies: 220
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, color) FROM stdin;
9	General	#3b82f6
10	Otros	#c8a62b
5	Facturación	#8e2c96
1	Hardware	#b06e23
2	Software	#3bc1f6
3	Conectividad	#20938a
4	Cuenta	#8441c2
6	Seguridad	#ccc03c
8	Sitio web	#d78257
7	Aplicación móvil	#bd922e
\.


--
-- TOC entry 4934 (class 0 OID 16442)
-- Dependencies: 224
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.messages (id, content, sent_at, support_case_id, user_id, sender) FROM stdin;
1	No puedo acceder a mi cuenta con la contraseña habitual	2025-04-30 12:15:00+02	1	1	usuario1
2	¿Has intentado restablecer tu contraseña?	2025-04-30 12:20:00+02	1	4	soporte1
3	Sí, pero no recibo el correo de restablecimiento	2025-04-30 12:25:00+02	1	1	usuario1
4	Mi pantalla se queda negra al encender el dispositivo	2025-05-01 16:30:00+02	2	2	usuario2
5	¿Puedes decirnos qué modelo de dispositivo usas?	2025-05-01 16:35:00+02	2	5	soporte2
6	La app se cierra cuando intento pagar	2025-05-02 11:45:00+02	3	3	usuario3
7	¿Qué versión de la aplicación estás usando?	2025-05-02 11:50:00+02	3	4	soporte1
8	La versión 2.3.4	2025-05-02 11:55:00+02	3	3	usuario3
9	No puedo conectarme al wifi corporativo	2025-05-02 18:20:00+02	4	6	usuario4
10	El problema ha sido resuelto, gracias	2025-05-02 19:45:00+02	4	6	usuario4
11	Me han cobrado dos veces el mismo servicio	2025-05-03 13:05:00+02	5	7	usuario5
12	Vamos a verificar la facturación	2025-05-03 13:10:00+02	5	5	soporte2
13	¿Puedes proporcionarnos los números de factura?	2025-05-03 13:15:00+02	5	5	soporte2
14	El software no se actualiza correctamente	2025-05-04 10:30:00+02	6	3	usuario3
15	¿Qué error te muestra al actualizar?	2025-05-04 10:35:00+02	6	9	soporte3
16	Error 404 al intentar descargar la actualización	2025-05-04 10:40:00+02	6	3	usuario3
17	Recibo alertas de seguridad continuamente	2025-05-04 16:45:00+02	7	2	usuario2
18	Vamos a revisar tu configuración de seguridad	2025-05-04 16:50:00+02	7	4	soporte1
19	La página principal no carga correctamente	2025-05-05 12:10:00+02	8	1	usuario1
20	Problema solucionado, era un error temporal	2025-05-05 13:30:00+02	8	1	usuario1
\.


--
-- TOC entry 4932 (class 0 OID 16424)
-- Dependencies: 222
-- Data for Name: support_cases; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.support_cases (id, title, description, status, created_at, user_id, category_id) FROM stdin;
1	No puedo iniciar sesión	Intento iniciar sesión pero el sistema me rechaza la contraseña	abierta	2025-04-30 10:15:00	1	4
2	Pantalla en negro	Mi dispositivo muestra pantalla negra al encenderlo	en espera	2025-05-01 14:30:00	2	1
3	Error en la aplicación	La aplicación se cierra cuando intento abrir la sección de pagos	abierta	2025-05-02 09:45:00	3	7
4	Problemas de conexión	No puedo conectarme al wifi de la empresa	cerrado	2025-05-02 16:20:00	6	3
5	Cobro incorrecto	Me han facturado dos veces el mismo servicio	abierta	2025-05-03 11:05:00	7	5
6	Actualización fallida	El software no se actualiza correctamente	en espera	2025-05-04 08:30:00	3	2
7	Error de seguridad	Recibo alertas de seguridad sin motivo aparente	abierta	2025-05-04 14:45:00	2	6
8	Página no carga	La página principal del sitio web no carga correctamente	cerrado	2025-05-05 10:10:00	1	8
\.


--
-- TOC entry 4928 (class 0 OID 16412)
-- Dependencies: 218
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password, role) FROM stdin;
1	usuario1	1234	user
2	usuario2	1234	user
3	usuario3	1234	user
4	soporte1	1234	support
5	soporte2	1234	support
6	usuario4	1234	user
7	usuario5	1234	user
8	usuario6	1234	user
9	soporte3	1234	support
\.


--
-- TOC entry 4940 (class 0 OID 0)
-- Dependencies: 219
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 10, true);


--
-- TOC entry 4941 (class 0 OID 0)
-- Dependencies: 223
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.messages_id_seq', 20, true);


--
-- TOC entry 4942 (class 0 OID 0)
-- Dependencies: 221
-- Name: support_case_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.support_case_id_seq', 8, true);


--
-- TOC entry 4943 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 9, true);


--
-- TOC entry 4769 (class 2606 OID 16422)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 4777 (class 2606 OID 16447)
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- TOC entry 4773 (class 2606 OID 16430)
-- Name: support_cases support_case_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.support_cases
    ADD CONSTRAINT support_case_pkey PRIMARY KEY (id);


--
-- TOC entry 4767 (class 2606 OID 16416)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4774 (class 1259 OID 16460)
-- Name: idx_messages_support_case; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_messages_support_case ON public.messages USING btree (support_case_id);


--
-- TOC entry 4775 (class 1259 OID 16461)
-- Name: idx_messages_user; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_messages_user ON public.messages USING btree (user_id);


--
-- TOC entry 4770 (class 1259 OID 16459)
-- Name: idx_support_case_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_support_case_status ON public.support_cases USING btree (status);


--
-- TOC entry 4771 (class 1259 OID 16458)
-- Name: idx_support_case_user; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_support_case_user ON public.support_cases USING btree (user_id);


--
-- TOC entry 4780 (class 2606 OID 16448)
-- Name: messages messages_support_case_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_support_case_id_fkey FOREIGN KEY (support_case_id) REFERENCES public.support_cases(id);


--
-- TOC entry 4781 (class 2606 OID 16453)
-- Name: messages messages_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 4778 (class 2606 OID 16436)
-- Name: support_cases support_case_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.support_cases
    ADD CONSTRAINT support_case_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- TOC entry 4779 (class 2606 OID 16431)
-- Name: support_cases support_case_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.support_cases
    ADD CONSTRAINT support_case_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2025-05-14 16:18:01

--
-- PostgreSQL database dump complete
--

