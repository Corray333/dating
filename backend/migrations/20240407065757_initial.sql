-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS public.users
(
    user_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 CACHE 1 ),
    username character varying(20) COLLATE pg_catalog."default" NOT NULL UNIQUE,   
    email text COLLATE pg_catalog."default" NOT NULL UNIQUE,
    phone text COLLATE pg_catalog."default",
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
    search int NOT NULL DEFAULT 1,
    referal text COLLATE pg_catalog."default",
    by_referal character varying(8) COLLATE pg_catalog."default" DEFAULT '',
    email_verified boolean NOT NULL DEFAULT false,
    phone_verified boolean NOT NULL DEFAULT false,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS public.searching_users
(
    user_id bigint NOT NULL,
    city text COLLATE pg_catalog."default",
    sex integer NOT NULL DEFAULT 0,
    orientation integer NOT NULL DEFAULT 0,
    birth TIMESTAMP NOT NULL,
    search int NOT NULL DEFAULT 1,
    CONSTRAINT searching_users_pkey PRIMARY KEY (user_id)
);



CREATE TABLE IF NOT EXISTS public.user_interest
(
    user_id bigint NOT NULL ,
    interest integer NOT NULL,
    CONSTRAINT user_interest_pkey PRIMARY KEY (user_id, interest),
    CONSTRAINT user_interest_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.chats
(
    chat_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 0 MINVALUE 0 CACHE 1 ),
    user1_id bigint NOT NULL,
    user2_id bigint NOT NULL,
    CONSTRAINT chats_pkey PRIMARY KEY (chat_id),
    CONSTRAINT chats_user1_id_fkey FOREIGN KEY (user1_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT chats_user2_id_fkey FOREIGN KEY (user2_id)
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
    type int NOT NULL DEFAULT 0,
    content text COLLATE pg_catalog."default" NOT NULL,
    time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS user_interest;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
