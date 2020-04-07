--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2 (Debian 12.2-2.pgdg100+1)
-- Dumped by pg_dump version 12.2

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: boms; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.boms (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    project character varying(255) NOT NULL,
    version character varying(255) NOT NULL
);


ALTER TABLE public.boms OWNER TO postgres;

--
-- Name: component_licenses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.component_licenses (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    license_id character varying(255) NOT NULL,
    component_id uuid NOT NULL
);


ALTER TABLE public.component_licenses OWNER TO postgres;

--
-- Name: components; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.components (
    id uuid NOT NULL,
    type character varying(255) NOT NULL,
    publisher character varying(255),
    "group" character varying(255),
    name character varying(255) NOT NULL,
    version character varying(255) NOT NULL,
    purl character varying(255),
    licenses text[],
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    bom_id uuid NOT NULL
);


ALTER TABLE public.components OWNER TO postgres;

--
-- Name: licenses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.licenses (
    id character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    name character varying(255),
    url character varying(255)
);


ALTER TABLE public.licenses OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: boms boms_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.boms
    ADD CONSTRAINT boms_pkey PRIMARY KEY (id);


--
-- Name: component_licenses component_licenses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.component_licenses
    ADD CONSTRAINT component_licenses_pkey PRIMARY KEY (license_id, component_id);


--
-- Name: components components_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.components
    ADD CONSTRAINT components_pkey PRIMARY KEY (id);


--
-- Name: licenses licenses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.licenses
    ADD CONSTRAINT licenses_pkey PRIMARY KEY (id);


--
-- Name: boms_project_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX boms_project_version_idx ON public.boms USING btree (project, version);


--
-- Name: components_name_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX components_name_version_idx ON public.components USING btree (name, version);


--
-- Name: components_purl_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX components_purl_idx ON public.components USING btree (purl);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: component_licenses component_licenses_component_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.component_licenses
    ADD CONSTRAINT component_licenses_component_id_fkey FOREIGN KEY (component_id) REFERENCES public.components(id) ON DELETE CASCADE;


--
-- Name: component_licenses component_licenses_license_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.component_licenses
    ADD CONSTRAINT component_licenses_license_id_fkey FOREIGN KEY (license_id) REFERENCES public.licenses(id) ON DELETE CASCADE;


--
-- Name: components components_bom_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.components
    ADD CONSTRAINT components_bom_id_fkey FOREIGN KEY (bom_id) REFERENCES public.boms(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

