-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN phone SET DEFAULT '';
ALTER TABLE users ALTER COLUMN avatar SET DEFAULT 'http://localhost:3002/files/images/avatars/default_avatar.png';
ALTER TABLE users ALTER COLUMN patronymic SET DEFAULT '';
ALTER TABLE users ALTER COLUMN city SET DEFAULT '';
ALTER TABLE users ALTER COLUMN bio SET DEFAULT '';
ALTER TABLE users ALTER COLUMN referal SET DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ALTER COLUMN phone SET DEFAULT NULL;
ALTER TABLE users ALTER COLUMN avatar SET DEFAULT 'http://localhost:3000/files/images/avatars/default_avatar.png';
ALTER TABLE users ALTER COLUMN patronymic SET DEFAULT NULL;
ALTER TABLE users ALTER COLUMN city SET DEFAULT NULL;
ALTER TABLE users ALTER COLUMN bio SET DEFAULT NULL;
ALTER TABLE users ALTER COLUMN referal SET DEFAULT NULL;
-- +goose StatementEnd
