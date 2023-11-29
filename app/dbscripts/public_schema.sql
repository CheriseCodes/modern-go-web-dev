SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET client_min_messages = warning;
SET row_security = off;
CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA pg_catalog;
SET search_path = public, pg_catalog;
SET default_tablespace = '';
-- runners (athletics)
CREATE TABLE runners (
    id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    first_name text NOT NULL,
    last_name text NOT NULL,
    age integer,
    is_active boolean DEFAULT TRUE,
    country text NOT NULL,
    personal_best interval,
    season_best interval,
    COUNSTRAINT runners_pk PRIMARY KEY (id)
)
