create table if not exists users (
    id UUID primary key,
    login varchar(255) not null,
    password_hash varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
)
