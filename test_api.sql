--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.4

-- Started on 2023-11-26 14:29:38 IST

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

ALTER TABLE IF EXISTS ONLY public.user_attribs DROP CONSTRAINT IF EXISTS user_attribs_fk;
DROP TRIGGER IF EXISTS on_user_update ON public.users;
DROP TRIGGER IF EXISTS on_user_details_change ON public.user_attribs;
DROP INDEX IF EXISTS public.users_username_idx;
DROP INDEX IF EXISTS public.user_attribs_user_id_idx;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_pk;
ALTER TABLE IF EXISTS ONLY public.user_attribs DROP CONSTRAINT IF EXISTS user_attribs_pk;
ALTER TABLE IF EXISTS public.users ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.user_attribs ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.users_id_seq;
DROP TABLE IF EXISTS public.users;
DROP SEQUENCE IF EXISTS public.user_details_id_seq;
DROP TABLE IF EXISTS public.user_attribs;
DROP FUNCTION IF EXISTS public.updated_at();
--
-- TOC entry 218 (class 1255 OID 16404)
-- Name: updated_at(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    new.updated_at = current_timestamp;
END;
$$;


ALTER FUNCTION public.updated_at() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 16410)
-- Name: user_attribs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_attribs (
    id integer NOT NULL,
    user_id integer NOT NULL,
    key_name character varying(35) NOT NULL,
    value text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.user_attribs OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16409)
-- Name: user_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_details_id_seq OWNER TO postgres;

--
-- TOC entry 4001 (class 0 OID 0)
-- Dependencies: 216
-- Name: user_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_details_id_seq OWNED BY public.user_attribs.id;


--
-- TOC entry 215 (class 1259 OID 16394)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(15) NOT NULL,
    middle_name character varying(15) DEFAULT ''::character varying NOT NULL,
    last_name character varying(20) NOT NULL,
    username character varying(24) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16393)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4002 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3842 (class 2604 OID 16413)
-- Name: user_attribs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_attribs ALTER COLUMN id SET DEFAULT nextval('public.user_details_id_seq'::regclass);


--
-- TOC entry 3838 (class 2604 OID 16397)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3849 (class 2606 OID 16419)
-- Name: user_attribs user_attribs_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_attribs
    ADD CONSTRAINT user_attribs_pk PRIMARY KEY (id);


--
-- TOC entry 3846 (class 2606 OID 16402)
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- TOC entry 3850 (class 1259 OID 16436)
-- Name: user_attribs_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_attribs_user_id_idx ON public.user_attribs USING btree (user_id, key_name);


--
-- TOC entry 3847 (class 1259 OID 16403)
-- Name: users_username_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);


--
-- TOC entry 3853 (class 2620 OID 16437)
-- Name: user_attribs on_user_details_change; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER on_user_details_change BEFORE UPDATE ON public.user_attribs FOR EACH ROW EXECUTE FUNCTION public.updated_at();


--
-- TOC entry 3852 (class 2620 OID 16406)
-- Name: users on_user_update; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER on_user_update BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.updated_at();


--
-- TOC entry 3851 (class 2606 OID 16425)
-- Name: user_attribs user_attribs_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_attribs
    ADD CONSTRAINT user_attribs_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2023-11-26 14:29:38 IST

--
-- PostgreSQL database dump complete
--

