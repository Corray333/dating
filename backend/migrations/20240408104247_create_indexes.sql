-- +goose Up
-- +goose StatementBegin

-- +goose StatementEnd
CREATE INDEX idx_city ON users (city);
CREATE INDEX idx_orientation ON users (orientation);
CREATE INDEX idx_sex ON users (sex);
CREATE INDEX idx_orientation_sex ON users (orientation, sex);
CREATE INDEX idx_city_sex ON users (city, sex);
CREATE INDEX idx_orientation_sex_city ON users (orientation, sex, city);
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_city;
DROP INDEX IF EXISTS idx_orientation;
DROP INDEX IF EXISTS idx_sex;
DROP INDEX IF EXISTS idx_orientation_sex;
DROP INDEX IF EXISTS idx_city_sex;
DROP INDEX IF EXISTS idx_orientation_sex_city;
-- +goose StatementEnd
