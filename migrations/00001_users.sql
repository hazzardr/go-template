-- +goose Up
create table users (
    id integer primary key autoincrement,
    username text not null
);

-- +goose Down
drop table users;
