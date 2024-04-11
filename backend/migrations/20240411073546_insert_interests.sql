-- +goose Up
-- +goose StatementBegin
INSERT INTO interests (name, icon) VALUES 
('Sport', 'mdi:football'),
('Music', 'mdi:music'),
('Art', 'mdi:palette'),
('Cinema', 'mdi:movie'),
('Books', 'mdi:book'),
('Travel', 'mdi:airplane'),
('Cooking', 'mdi:food'),
('Photography', 'mdi:camera'),
('Dance', 'mdi:dance-ballroom'),
('Theatre', 'mdi:theatre'),
('Games', 'mdi:gamepad-variant'),
('Nature', 'mdi:tree'),
('Animals', 'mdi:dog'),
('Fashion', 'mdi:shoe-heel'),
('Cars', 'mdi:car'),
('Science', 'mdi:atom'),
('Politics', 'mdi:account-tie'),
('History', 'mdi:history'),
('Psychology', 'mdi:brain'),
('Philosophy', 'mdi:lightbulb'),
('Religion', 'mdi:church'),
('Esoterics', 'mdi:crystal-ball'),
('Astrology', 'mdi:star'),
('Cultures', 'mdi:earth'),
('Languages', 'mdi:language-html5'),
('Education', 'mdi:school');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
