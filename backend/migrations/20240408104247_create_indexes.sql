-- +goose Up
-- +goose StatementBegin

-- +goose StatementEnd
CREATE INDEX idx_city ON users (city);
CREATE INDEX idx_oriendation_id ON users (orientation_id);
CREATE INDEX idx_sex ON users (sex);
CREATE INDEX idx_orientation_id_sex ON users (orientation_id, sex);
CREATE INDEX idx_city_sex ON users (city, sex);
CREATE INDEX idx_orientation_id_sex_city ON users (orientation_id, sex, city);
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_city;
DROP INDEX IF EXISTS idx_orientation_id;
DROP INDEX IF EXISTS idx_sex;
DROP INDEX IF EXISTS idx_orientation_id_sex;
DROP INDEX IF EXISTS idx_city_sex;
DROP INDEX IF EXISTS idx_orientation_id_sex_city;
-- +goose StatementEnd
