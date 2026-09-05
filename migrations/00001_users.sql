-- +goose Up
create table users (
    id bigserial primary key, -- auto incrementing id
    username varchar(255) not null
);

-- +goose Down
drop table users;
