create table "user"(
    id bigserial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null
);

create table todo_list(
    id bigserial primary key,
    title varchar(255) not null,
    description text
);

create table todo_item(
    id bigserial primary key,
    title varchar(255) not null,
    description text,
    is_completed bool not null default false
);

create table users_lists(
    id bigserial unique not null,
    user_id bigint references "user"(id) on delete cascade not null,
    list_id bigint references todo_list(id) on delete cascade not null
);

create table items_lists(
    id  bigserial primary key,
    item_id bigint references todo_item(id) on delete cascade not null,
    list_id bigint references todo_list(id) on delete cascade not null
);