-- +goose Up
-- +goose StatementBegin
INSERT INTO users VALUES (DEFAULT, 'Markovnik', 'markovnik3457@gmail.com', '89320509129', '$2a$14$kvhZe4Kf2DyQDckxiSH7Tu8dm5qXTJdTvWXhFGov8TVZTc4Zmuz/G', 'http://localhost:3002/files/images/avatars/avatar0.png', 'Mark', 'Anikin', '', 'Krasnodar', 'Suck my dick', 0, 0, '2002-10-12 10:57:00', 1, 'ABCDEFGH', '', DEFAULT, DEFAULT);
INSERT INTO users VALUES (DEFAULT, 'Lububuska', 'maria.masenko2004@gmail.com', '89384000874', '$2a$14$kvhZe4Kf2DyQDckxiSH7Tu8dm5qXTJdTvWXhFGov8TVZTc4Zmuz/G', 'http://localhost:3002/files/images/avatars/avatar1.png', 'Maria', 'Masenko', '', 'Krasnodar', 'chipi chipi chapa chapa', 1, 0, '2004-02-11 10:57:00', 1, 'BCDEFGHI', 'ABCDEFGH', DEFAULT, DEFAULT);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users;
-- +goose StatementEnd
