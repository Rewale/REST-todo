CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    username varchar(255) unique not null,
    password_hash varchar(255) not null
);

create table todo_lists
(
    id SERIAL PRIMARY KEY,
    title varchar(255) not null,
    description varchar(255)
);

create table users_lists
(
    id SERIAL PRIMARY KEY,
    user_id int references users(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);

create table todo_items
(
    id SERIAL primary key,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

create table lists_items
(
    id SERIAL PRIMARY KEY,
    item_id int references todo_items(id) on delete cascade not null ,
    list_id int references todo_lists(id) on delete cascade not null
);

