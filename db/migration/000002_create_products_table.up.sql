CREATE TABLE IF NOT EXISTS public.products (
    product_id integer NOT NULL,
    product_name character varying(100),
    price integer NOT NULL
);

ALTER TABLE public.products ALTER COLUMN product_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.products_product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);

