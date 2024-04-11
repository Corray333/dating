-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS public.users
(
    user_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 CACHE 1 ),
    username character varying(20) COLLATE pg_catalog."default" NOT NULL UNIQUE,   
    email text COLLATE pg_catalog."default" NOT NULL UNIQUE,
    phone text COLLATE pg_catalog."default" UNIQUE,
    password text COLLATE pg_catalog."default" NOT NULL,
    avatar text COLLATE pg_catalog."default" NOT NULL DEFAULT 'http://localhost:3000/files/images/avatars/default_avatar.png'::text,
    name text COLLATE pg_catalog."default" NOT NULL,
    surname text COLLATE pg_catalog."default" NOT NULL,
    patronymic text COLLATE pg_catalog."default",
    city text COLLATE pg_catalog."default",
    bio text COLLATE pg_catalog."default",
    sex integer NOT NULL DEFAULT 0,
    orientation integer NOT NULL DEFAULT 0,
    birth TIMESTAMP NOT NULL,
    referal text COLLATE pg_catalog."default",
    is_submitted boolean NOT NULL DEFAULT false,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS public.user_search
(
    user_id bigint NOT NULL,
    search integer NOT NULL DEFAULT 0,
    CONSTRAINT pk1 PRIMARY KEY (user_id, search),
    CONSTRAINT fk1 FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS public.interests
(
    interest_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 CACHE 1 ),
    name text COLLATE pg_catalog."default" NOT NULL,
    icon text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT interests_pkey PRIMARY KEY (interest_id)
);

CREATE TABLE IF NOT EXISTS public.user_interest
(
    user_id bigint NOT NULL,
    interest_id integer NOT NULL,
    CONSTRAINT user_interest_pkey PRIMARY KEY (user_id, interest_id),
    CONSTRAINT user_interest_interest_id_fkey FOREIGN KEY (interest_id)
        REFERENCES public.interests (interest_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT user_interest_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS public.chats
(
    chat_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 CACHE 1 ),
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT chats_pkey PRIMARY KEY (chat_id)
);

CREATE TABLE IF NOT EXISTS public.chat_user
(
    chat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    CONSTRAINT chat_user_pkey PRIMARY KEY (chat_id, user_id),
    CONSTRAINT chat_user_chat_id_fkey FOREIGN KEY (chat_id)
        REFERENCES public.chats (chat_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT chat_user_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS public.chat_messages
(
    chat_message_id uuid NOT NULL,
    chat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    text text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT chat_messages_pkey PRIMARY KEY (chat_message_id),
    CONSTRAINT chat_messages_chat_id_fkey FOREIGN KEY (chat_id)
        REFERENCES public.chats (chat_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT chat_messages_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS public.attachments
(
    chat_message_id uuid NOT NULL,
    type text COLLATE pg_catalog."default" NOT NULL,
    url text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT attachments_pkey PRIMARY KEY (chat_message_id, url),
    CONSTRAINT attachments_chat_message_id_fkey FOREIGN KEY (chat_message_id)
        REFERENCES public.chat_messages (chat_message_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS attachments;
DROP TABLE IF EXISTS chat_messages;
DROP TABLE IF EXISTS chat_user;
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS user_id_search;

DROP TABLE IF EXISTS user_interest;
DROP TABLE IF EXISTS interests;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
