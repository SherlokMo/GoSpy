BEGIN;

-- Table: public.sites

-- DROP TABLE IF EXISTS public.sites;

CREATE TABLE IF NOT EXISTS public.sites
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    title character varying COLLATE pg_catalog."default" NOT NULL,
    url character varying COLLATE pg_catalog."default" NOT NULL,
    "interval" integer NOT NULL DEFAULT 10,
    CONSTRAINT sites_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.sites
    OWNER to postgres;

-- Table: public.lookup

-- DROP TABLE IF EXISTS public.lookup;

CREATE TABLE IF NOT EXISTS public.lookup
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    site_id bigint NOT NULL,
    dnslookup bigint NOT NULL,
    connection bigint NOT NULL,
    tlshandshake bigint NOT NULL,
    warning character varying COLLATE pg_catalog."default",
    status_code integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT lookup_pkey PRIMARY KEY (id),
    CONSTRAINT lookup_site_id_fkey FOREIGN KEY (site_id)
        REFERENCES public.sites (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.lookup
    OWNER to postgres;

COMMIT;