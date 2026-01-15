CREATE TABLE IF NOT EXISTS public.posts (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    content text NOT NULL,
    user_id integer,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone
);

ALTER TABLE public.posts ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);