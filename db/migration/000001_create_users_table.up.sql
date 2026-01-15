CREATE TABLE IF NOT EXISTS public.users (
    id integer NOT NULL,
    email character varying(100),
    gender character(1) DEFAULT 'L'::bpchar NOT NULL,
    profile_img character varying(255),
    password character varying(255) NOT NULL,
    role character varying(5) DEFAULT 'user'::character varying NOT NULL,
    CONSTRAINT users_gender_check CHECK ((gender = ANY (ARRAY['L'::bpchar, 'P'::bpchar])))
);

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (email);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_profile_img_key UNIQUE (profile_img);


