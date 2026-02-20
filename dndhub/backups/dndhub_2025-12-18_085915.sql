--
-- PostgreSQL database dump
--

\restrict mWAVE5XPbjoMzSXtBwDcFXxT8UU39pXtc4zoJZucsJA9JESen2zwJ88lXBgCT4w

-- Dumped from database version 16.10 (Debian 16.10-1.pgdg13+1)
-- Dumped by pg_dump version 16.10 (Debian 16.10-1.pgdg13+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: franchise_history; Type: TABLE; Schema: public; Owner: myappuser
--

CREATE TABLE public.franchise_history (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    franchise_id uuid NOT NULL,
    revenue bigint DEFAULT 0 NOT NULL,
    expenses bigint DEFAULT 0 NOT NULL,
    profit bigint DEFAULT 0 NOT NULL,
    roll integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    activities text
);


ALTER TABLE public.franchise_history OWNER TO myappuser;

--
-- Name: franchises; Type: TABLE; Schema: public; Owner: myappuser
--

CREATE TABLE public.franchises (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    name text NOT NULL,
    funds_cents bigint DEFAULT 0 NOT NULL,
    property_value_cents bigint DEFAULT 0 NOT NULL,
    unskilled_workers integer DEFAULT 0 NOT NULL,
    lowskilled_workers integer DEFAULT 0 NOT NULL,
    highskilled_workers integer DEFAULT 0 NOT NULL,
    cost_unskilled_cents integer DEFAULT 0 NOT NULL,
    cost_lowskilled_cents integer DEFAULT 0 NOT NULL,
    cost_highskilled_cents integer DEFAULT 0 NOT NULL,
    revenue_modifier_bp integer DEFAULT 0 NOT NULL,
    upkeep_modifier_bp integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.franchises OWNER TO myappuser;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: myappuser
--

CREATE TABLE public.sessions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    token text NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    client_info text,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.sessions OWNER TO myappuser;

--
-- Name: unique_workers; Type: TABLE; Schema: public; Owner: myappuser
--

CREATE TABLE public.unique_workers (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    franchise_id uuid NOT NULL,
    name text NOT NULL,
    monthly_cost_cents integer DEFAULT 0 NOT NULL,
    creativity integer DEFAULT 0 NOT NULL,
    discipline integer DEFAULT 0 NOT NULL,
    charisma integer DEFAULT 0 NOT NULL,
    efficiency integer DEFAULT 0 NOT NULL,
    exploration integer DEFAULT 0 NOT NULL,
    notes text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.unique_workers OWNER TO myappuser;

--
-- Name: users; Type: TABLE; Schema: public; Owner: myappuser
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    username text,
    password_hash text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO myappuser;

--
-- Data for Name: franchise_history; Type: TABLE DATA; Schema: public; Owner: myappuser
--

COPY public.franchise_history (id, franchise_id, revenue, expenses, profit, roll, created_at, activities) FROM stdin;
0145d12a-b29e-41d3-ad34-dfef111e31a0	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	19890	1036	18854	23	2025-12-18 07:19:25.117856+00	\N
3489f86c-224d-4805-9008-cb783fd71e0d	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	18750	1036	17714	60	2025-12-18 07:20:17.136047+00	\N
e329ecfa-e09f-498f-95dd-03394dcbc00e	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	2340	1036	1304	22	2025-12-18 07:20:44.616723+00	\N
fc279ca8-73eb-4b71-8939-66c06dbe75ce	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	7263	1036	6227	96	2025-12-18 07:21:46.788421+00	\N
33dfd7dc-119a-49db-b221-beebb022ecab	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	6650	1996	4654	88	2025-12-18 07:22:03.121651+00	\N
32607b76-ae14-4393-85c8-3814db015d6f	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	7510	2313	5197	82	2025-12-18 07:22:36.810683+00	\N
efcb2e80-e0ba-4202-be99-58f748e8973f	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	8352	2313	6039	100	2025-12-18 07:22:49.91284+00	\N
687d647e-9b76-4681-90ec-5b9c9ee9f61f	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4591	2313	2278	59	2025-12-18 07:22:53.075903+00	\N
17224a3d-992b-47e9-8a4a-29ce059f76fb	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	7276	2313	4963	85	2025-12-18 07:23:00.066399+00	\N
19a88602-02f4-4a6f-b514-396b966ce301	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	1803	2313	-510	20	2025-12-18 07:23:01.112564+00	\N
2fd06e13-0e2d-4c46-a4b9-eef0218ddf88	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4473	2313	2160	43	2025-12-18 07:23:08.997927+00	\N
142a9587-ff3b-4e13-a62e-b58bf03c8d43	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4442	2313	2129	58	2025-12-18 07:23:10.021413+00	\N
a60dcd3e-63bf-47f8-aa38-660b17f12eca	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4413	2313	2100	60	2025-12-18 07:23:10.552382+00	\N
684b9466-a3e6-4240-aecd-554a38381ceb	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	3510	2313	1197	32	2025-12-18 07:23:11.092368+00	\N
806979fb-5939-44b9-8fd1-c2725b975d6f	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	6111	2313	3798	75	2025-12-18 07:23:19.644265+00	\N
896a2623-b9b9-46f3-84f3-4b1a0f8ffb1f	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4343	2313	2030	58	2025-12-18 07:23:21.0321+00	\N
bea90647-9254-4792-aa9f-694bde1087f4	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	0	2313	-2313	1	2025-12-18 07:23:21.832187+00	\N
783c0fc2-668b-44b0-9ebf-90d6e4207597	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	6032	2313	3719	80	2025-12-18 07:23:39.335807+00	\N
b7b63236-0858-4a81-8c70-ec85d2996284	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	0	2410	-2410	1	2025-12-18 07:28:30.963448+00	Marketing by Bob Bobson: +133 Revenue Mod
9b2c5b81-efb3-412f-bf7f-1a9065f0ba13	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	2757	2410	347	26	2025-12-18 07:29:07.310699+00	Marketing by Bob Bobson: +133 Revenue Mod
d77db820-9dd8-4952-908d-07c3508d1e39	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	5662	2410	3252	67	2025-12-18 07:29:25.774366+00	Marketing by Bob Bobson: +133 Revenue Mod
150c1d1f-e4f2-420f-9d01-94e144dd7546	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	8691	2410	6281	92	2025-12-18 07:29:32.10534+00	Marketing by Bob Bobson: +133 Revenue Mod
f0adb02d-5c1e-4319-9280-e6dc41befd5d	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	5713	2410	3303	67	2025-12-18 07:33:21.121894+00	none
059df232-729f-4196-b836-fc7f18fd248c	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	5641	2410	3231	61	2025-12-18 07:33:22.717487+00	none
48f178b3-b815-4978-b3b4-fb5c7ff316a9	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4646	2410	2236	48	2025-12-18 07:33:23.20715+00	none
ecc420f0-12c6-4eb0-9eaa-ad160e4e598b	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	6436	2410	4026	77	2025-12-18 07:33:23.633234+00	none
c842b0bf-eb8f-4447-bbc4-8b43cfbfe768	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	8196	2410	5786	92	2025-12-18 07:33:23.914575+00	none
eaade3e7-573a-45f3-badc-a8ae607da2b2	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	7222	2410	4812	89	2025-12-18 07:33:24.125889+00	none
51ac8f50-c363-41cf-8a91-e939bd780382	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	0	2410	-2410	10	2025-12-18 07:33:24.278032+00	none
8977b377-0484-4990-b1db-97a22ead97fb	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	8001	2410	5591	94	2025-12-18 07:33:24.442119+00	none
f4720231-4c3f-4a6b-af7c-dfcfeeea8e5d	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	2649	2410	239	23	2025-12-18 07:33:24.602718+00	none
a29c1b5c-c6bd-4412-8b33-71713e3a98da	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	7902	2410	5492	92	2025-12-18 07:33:24.787657+00	none
e31cbaa7-8e0b-42c1-8221-ee4d344889dd	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4366	2410	1956	43	2025-12-18 07:33:24.953104+00	none
a19ecc94-e7af-4831-8343-46234e83ddb7	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	6083	2410	3673	76	2025-12-18 07:33:25.127539+00	none
94f9653d-cd5e-426d-9f4e-de8a7a0e7bf5	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	3461	2410	1051	36	2025-12-18 07:33:25.279238+00	none
331a36ea-a7d8-4131-ab01-78ce7106fd47	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	2585	2410	175	29	2025-12-18 07:33:25.431325+00	none
173ecfae-faee-46e7-a0e0-4640125878a0	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4293	2410	1883	60	2025-12-18 07:33:25.718998+00	none
3766f820-b5be-4e5f-b820-ea786f10552a	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	4280	2410	1870	45	2025-12-18 07:33:25.870347+00	none
\.


--
-- Data for Name: franchises; Type: TABLE DATA; Schema: public; Owner: myappuser
--

COPY public.franchises (id, user_id, name, funds_cents, property_value_cents, unskilled_workers, lowskilled_workers, highskilled_workers, cost_unskilled_cents, cost_lowskilled_cents, cost_highskilled_cents, revenue_modifier_bp, upkeep_modifier_bp, created_at, updated_at) FROM stdin;
3364a490-82ef-48fd-aeed-d39098a1f709	7c243cd1-82a3-4e37-90e3-1aa79159c0ee	sdf	0	0	0	0	0	0	0	0	0	0	2025-09-30 14:21:30.756187+00	2025-09-30 14:21:30.756187+00
6458d127-9eb8-4f96-819c-0dcaabf35e96	4db72abc-daf6-4b53-afa2-6590bfc445ed	Test	0	2	2	2	2	2	2	2	2	2	2025-09-19 13:08:39.645381+00	2025-09-30 14:24:45.725609+00
d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	397ae6cc-59c6-47f0-a6e2-298bbb2145f5	test	407025	20000	2	6	2	15	30	60	84	0	2025-12-15 07:46:38.864248+00	2025-12-18 07:33:25.869568+00
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: myappuser
--

COPY public.sessions (id, user_id, token, expires_at, client_info, created_at) FROM stdin;
\.


--
-- Data for Name: unique_workers; Type: TABLE DATA; Schema: public; Owner: myappuser
--

COPY public.unique_workers (id, franchise_id, name, monthly_cost_cents, creativity, discipline, charisma, efficiency, exploration, notes, created_at, updated_at) FROM stdin;
fb957451-c7ec-4f29-acec-519b3aef3d64	6458d127-9eb8-4f96-819c-0dcaabf35e96	Greta	2	0	0	0	0	0		2025-09-19 13:47:46.047358+00	2025-09-19 13:56:40.670347+00
b8ff54e6-6cbf-4e7b-aefd-4039cd30532e	3364a490-82ef-48fd-aeed-d39098a1f709	sdf	2	2	2	2	2	2		2025-09-30 14:22:10.655719+00	2025-09-30 14:22:10.655719+00
7919236f-7ab7-4bc6-b3d5-2fa43df0b9c1	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	Bob Bobson	60	3	4	4	2	1		2025-12-15 07:47:28.210692+00	2025-12-15 07:47:28.210692+00
0603294c-0953-425e-9937-e8606ffd9b5d	d1c6b6ad-a1df-4fe8-8875-b58264fb44b0	sdf	20	6	2	4	2	2		2025-12-15 07:54:23.6523+00	2025-12-18 07:20:35.385508+00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: myappuser
--

COPY public.users (id, username, password_hash, created_at) FROM stdin;
4db72abc-daf6-4b53-afa2-6590bfc445ed	raphael	$2a$10$x/YWBFRJ56wFvmbzh8T/KeVd9HulKFU8Jce3Zva2E0uJgW8Edmy1G	2025-09-19 13:01:58.734744+00
7c243cd1-82a3-4e37-90e3-1aa79159c0ee	test	$2a$10$KwomzaULU7BFFOwAbapIUOUyoRoiuadNILbxupmsyD9FtGpLMcOQO	2025-09-30 14:21:21.743428+00
397ae6cc-59c6-47f0-a6e2-298bbb2145f5	hallo	$2a$10$eQhfOKvzHRRgwj7S1msvxeVAJz2QWduqpMMjMcXjy6ZU6nTOhVjI6	2025-12-15 07:46:25.801799+00
5dc91ba8-3c2a-44ff-9a89-363c9cfe44de	hallo	$2a$10$RYOtnTZA1oBCykv.Eu4PBuwlsGyg6UkLwBVQSDHVot.A/Vg9.zRbG	2025-12-18 07:18:58.081093+00
\.


--
-- Name: franchise_history franchise_history_pkey; Type: CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.franchise_history
    ADD CONSTRAINT franchise_history_pkey PRIMARY KEY (id);


--
-- Name: franchises franchises_pkey; Type: CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.franchises
    ADD CONSTRAINT franchises_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_token_key; Type: CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_token_key UNIQUE (token);


--
-- Name: unique_workers unique_workers_pkey; Type: CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.unique_workers
    ADD CONSTRAINT unique_workers_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_sessions_token_expires; Type: INDEX; Schema: public; Owner: myappuser
--

CREATE INDEX idx_sessions_token_expires ON public.sessions USING btree (token, expires_at);


--
-- Name: idx_sessions_user_id; Type: INDEX; Schema: public; Owner: myappuser
--

CREATE INDEX idx_sessions_user_id ON public.sessions USING btree (user_id);


--
-- Name: franchise_history franchise_history_franchise_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.franchise_history
    ADD CONSTRAINT franchise_history_franchise_id_fkey FOREIGN KEY (franchise_id) REFERENCES public.franchises(id) ON DELETE CASCADE;


--
-- Name: franchises franchises_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.franchises
    ADD CONSTRAINT franchises_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: sessions sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: unique_workers unique_workers_franchise_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myappuser
--

ALTER TABLE ONLY public.unique_workers
    ADD CONSTRAINT unique_workers_franchise_id_fkey FOREIGN KEY (franchise_id) REFERENCES public.franchises(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict mWAVE5XPbjoMzSXtBwDcFXxT8UU39pXtc4zoJZucsJA9JESen2zwJ88lXBgCT4w

